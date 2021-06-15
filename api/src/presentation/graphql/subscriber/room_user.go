package subscriber

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/presentation/graphql/presenter"
	"github.com/laster18/poi/api/src/util/acontext"
)

type RoomUserSubscriber struct {
	client *redis.Client
	mutex  sync.Mutex
	// channels map[roomId]map[userId]chan ...
	chans map[int]map[string]chan model.RoomUserEvent
}

func NewRoomUserSubscriber(ctx context.Context, client *redis.Client) *RoomUserSubscriber {
	subscriber := &RoomUserSubscriber{
		client: client,
		chans:  make(map[int]map[string]chan model.RoomUserEvent),
	}
	go subscriber.start(ctx)
	return subscriber
}

func (s *RoomUserSubscriber) start(ctx context.Context) {
	logger := acontext.GetLogger(ctx)

	// subscribeCh "roomUser:*"
	subscribeCh := fmt.Sprintf("%s:%s:%s", redis.KeySpace, RoomUserChannel, "*")
	pubsub := s.client.PSubscribe(ctx, subscribeCh)
	defer pubsub.Close()

	for {
		logger.Debugf("on subscribe room user event, channel: %s \n", subscribeCh)

		msg := <-pubsub.Channel()

		logger.Debugf("subscribe roomUser, channel: %s, payload: %s\n\n", msg.Channel, msg.Payload)

		ch := removeKeyspacePrefix(msg.Channel)
		roomID, userUID, err := DestructRoomUserKey(ch)
		if err != nil {
			log.Println("getted invalid channel key from redis, err:", err)
			continue
		}

		switch msg.Payload {
		case redis.EventSet:
			// TODO: repositoryへ移譲
			roomUserJSON, err := s.client.Get(ctx, ch).Result()
			if err != nil {
				logger.Infof("failed to get from redis, err: %v", err)
				continue
			}

			var roomUser domain.RoomUser
			if err := json.Unmarshal([]byte(roomUserJSON), &roomUser); err != nil {
				logger.Info("received unexpected json data struct from redis")
				continue
			}

			d, err := s.makePublishDataFromSetEvent(&roomUser)
			if err != nil {
				log.Println(err)
				continue
			}

			s.deliver(roomID, d)
		case redis.EventDel:
			fallthrough
		case redis.EventExpired:
			data := &model.ExitedPayload{
				UserID: userUID,
			}
			s.deliver(roomID, data)
		default:
			fmt.Println("received unknown event:", msg.Payload)
		}
	}
}

func (s *RoomUserSubscriber) deliver(roomID int, data model.RoomUserEvent) {
	// TODO: ここLock必要？
	s.mutex.Lock()
	defer s.mutex.Unlock()

	chs, ok := s.chans[roomID]

	if !ok {
		return
	}

	for _, ch := range chs {
		ch <- data
	}
}

func (s *RoomUserSubscriber) AddCh(ch chan model.RoomUserEvent, roomID int, userUID string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	userChannels, ok := s.chans[roomID]
	if !ok {
		userChannels = make(map[string]chan model.RoomUserEvent)
		s.chans[roomID] = userChannels
	}
	userChannels[userUID] = ch
}

func (s *RoomUserSubscriber) RemoveCh(roomID int, userUID string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	userChannels, ok := s.chans[roomID]
	if !ok {
		return
	}
	delete(userChannels, userUID)
}

func (s *RoomUserSubscriber) makePublishDataFromSetEvent(ru *domain.RoomUser) (model.RoomUserEvent, error) {
	switch ru.LastEvent {
	case domain.JoinEvent:
		return presenter.ToJoinedPayload(ru), nil
	case domain.MoveEvent:
		return presenter.ToMovedPayload(ru), nil
	case domain.AddMessageEvent:
		return presenter.ToSentMassagePayload(ru), nil
	case domain.RemoveLastMessageEvent:
		return presenter.ToRemovedLastMessagePayload(ru), nil
	case domain.ChangeBalloonPositionEvent:
		return presenter.ToChangedBalloonPositionPayload(ru), nil
	default:
		return nil, errors.New("getted unknown roomUser event")
	}
}
