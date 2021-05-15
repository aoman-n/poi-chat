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
)

type RoomUserSubscriber struct {
	client *redis.Client
	Mutex  sync.Mutex
	// channels map[roomId]map[userId]chan ...
	chs map[int]map[string]chan model.RoomUserEvent
}

func NewRoomUserSubscriber(ctx context.Context, client *redis.Client) *RoomUserSubscriber {
	subscriber := &RoomUserSubscriber{
		client: client,
		Mutex:  sync.Mutex{},
		chs:    make(map[int]map[string]chan model.RoomUserEvent),
	}
	go subscriber.start(ctx)
	return subscriber
}

func (s *RoomUserSubscriber) start(ctx context.Context) {
	// subscribeCh "roomUser:*"
	subscribeCh := fmt.Sprintf("%s:%s:%s", redis.KeySpace, RoomUserChannel, "*")
	pubsub := s.client.PSubscribe(ctx, subscribeCh)
	defer pubsub.Close()

	for {
		msg := <-pubsub.Channel()

		// debug log
		log.Printf("subscribe roomUser, channel: %s, payload: %s\n\n", msg.Channel, msg.Payload)

		ch := removeKeyspacePrefix(msg.Channel)
		roomID, userID, err := destructRoomUserKey(ch)
		if err != nil {
			log.Println("getted invalid channel key from redis")
			continue
		}

		switch msg.Payload {
		case redis.EventSet:
			roomUserJSON, err := s.client.Get(ctx, ch).Result()
			if err != nil {
				log.Println("failed to get from redis, err:", err)
				continue
			}

			var roomUser domain.RoomUser
			if err := json.Unmarshal([]byte(roomUserJSON), &roomUser); err != nil {
				log.Println("getted unexpected json data struct from redis")
				continue
			}

			d, err := s.makeDataFromSet(&roomUser, roomID, userID)
			if err != nil {
				log.Println(err)
			}
			s.deliver(roomID, d)
		case redis.EventDel:
			fallthrough
		case redis.EventExpired:
			data := &model.Exited{
				UserID: makeRoomUserID(userID),
			}
			s.deliver(roomID, data)
		default:
			fmt.Println("received unknown event:", msg.Payload)
		}
	}
}

func (s *RoomUserSubscriber) deliver(roomID int, data model.RoomUserEvent) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	chs, ok := s.chs[roomID]
	if !ok {
		return
	}

	for _, ch := range chs {
		ch <- data
	}
}

func (s *RoomUserSubscriber) Subscribe(ctx context.Context, roomID int, userID string) <-chan model.RoomUserEvent {
	createdCh := make(chan model.RoomUserEvent)

	s.Mutex.Lock()
	userChannels, ok := s.chs[roomID]
	if !ok {
		userChannels = make(map[string]chan model.RoomUserEvent)
		s.chs[roomID] = userChannels
	}
	userChannels[userID] = createdCh
	s.Mutex.Unlock()

	go func() {
		<-ctx.Done()
	}()

	return createdCh
}

func (s *RoomUserSubscriber) makeDataFromSet(ru *domain.RoomUser, roomID, userID int) (model.RoomUserEvent, error) {
	switch ru.LastEvent {
	case domain.JoinEvent:
		return &model.Joined{
			UserID:    makeRoomUserID(userID),
			Name:      ru.Name,
			AvatarURL: ru.AvatarURL,
			X:         ru.X,
			Y:         ru.Y,
		}, nil
	case domain.MoveEvent:
		return &model.Moved{
			UserID: makeRoomUserID(userID),
			X:      ru.X,
			Y:      ru.Y,
		}, nil
	case domain.MessageEvent:
		return &model.SendedMassage{
			UserID:  makeRoomUserID(userID),
			Message: ru.LastMessage,
		}, nil
	default:
		return nil, errors.New("getted unknown roomUser event")
	}
}

func makeRoomUserID(roomUserID int) string {
	return fmt.Sprintf("RoomUser:%d", roomUserID)
}
