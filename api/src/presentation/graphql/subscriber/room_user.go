package subscriber

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/presentation/graphql/presenter"
	"github.com/laster18/poi/api/src/util/acontext"
)

type RoomUserSubscriber struct {
	client *redis.Client
	mutex  sync.Mutex
	// channels map[roomId]map[userId]chan ...
	chans map[int]map[int]chan model.RoomUserEvent
}

func NewRoomUserSubscriber(client *redis.Client) *RoomUserSubscriber {
	subscriber := &RoomUserSubscriber{
		client: client,
		chans:  make(map[int]map[int]chan model.RoomUserEvent),
	}
	return subscriber
}

func (s *RoomUserSubscriber) Start(ctx context.Context) {
	logger := acontext.GetLogger(ctx)

	subscribeCh := fmt.Sprintf("%s:%s:%s", redis.KeySpace, RoomUserStatusChannel, "*")
	pubsub := s.client.PSubscribe(ctx, subscribeCh)
	defer pubsub.Close()

	for {
		select {
		case <-ctx.Done():
			logger.Info("stop room_user subscriber")
			return
		case msg := <-pubsub.Channel():
			logger.Debugf("subscribe roomUser, channel: %s, payload: %s\n\n", msg.Channel, msg.Payload)

			ch := removeKeyspacePrefix(msg.Channel)
			roomID, userID, err := DestructRoomUserStatusKey(ch)
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

				var roomUserStatus room.UserStatus
				if err := json.Unmarshal([]byte(roomUserJSON), &roomUserStatus); err != nil {
					logger.Info("received unexpected json data struct from redis")
					continue
				}

				d, err := s.makePublishDataFromSetEvent(&roomUserStatus)
				if err != nil {
					log.Println(err)
					continue
				}

				s.deliver(roomID, d)
			case redis.EventDel:
				fallthrough
			case redis.EventExpired:
				data := &model.ExitedPayload{
					UserID: strconv.Itoa(userID),
				}
				s.deliver(roomID, data)
			default:
				fmt.Println("received unknown event:", msg.Payload)
			}
		}
	}
}

func (s *RoomUserSubscriber) deliver(roomID int, data model.RoomUserEvent) {
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

func (s *RoomUserSubscriber) AddCh(ch chan model.RoomUserEvent, roomID int, userID int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	userChannels, ok := s.chans[roomID]
	if !ok {
		userChannels = make(map[int]chan model.RoomUserEvent)
		s.chans[roomID] = userChannels
	}
	userChannels[userID] = ch
}

func (s *RoomUserSubscriber) RemoveCh(roomID int, userID int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	userChannels, ok := s.chans[roomID]
	if !ok {
		return
	}
	delete(userChannels, userID)
}

func (s *RoomUserSubscriber) makePublishDataFromSetEvent(ru *room.UserStatus) (model.RoomUserEvent, error) {
	switch ru.LastEvent {
	case room.EnterEvent:
		return presenter.ToEnteredPayload(ru), nil
	case room.MoveEvent:
		return presenter.ToMovedPayload(ru), nil
	case room.AddMessageEvent:
		return presenter.ToSentMassagePayload(ru), nil
	case room.RemoveLastMessageEvent:
		return presenter.ToRemovedLastMessagePayload(ru), nil
	case room.ChangeBalloonPositionEvent:
		return presenter.ToChangedBalloonPositionPayload(ru), nil
	default:
		return nil, errors.New("getted unknown roomUser event")
	}
}
