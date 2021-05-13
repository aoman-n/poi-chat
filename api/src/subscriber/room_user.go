package subscriber

import (
	"context"
	"encoding/json"
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
	chs map[int]map[string]chan *model.RoomUserEvent
}

func NewRoomUserSubscriber(ctx context.Context, client *redis.Client) *RoomUserSubscriber {
	subscriber := &RoomUserSubscriber{
		client: client,
		Mutex:  sync.Mutex{},
		chs:    make(map[int]map[string]chan *model.RoomUserEvent),
	}
	go subscriber.start(ctx)
	return subscriber
}

func (s *RoomUserSubscriber) start(ctx context.Context) {
	subscribeCh := fmt.Sprintf("%s:%s:%s", redis.KeySpace, RoomUserChannel, "*")
	pubsub := s.client.PSubscribe(ctx, subscribeCh)
	defer pubsub.Close()

	for {
		msg := <-pubsub.Channel()

		subscribedChannel := removeKeyspacePrefix(msg.Channel)
		roomID, userID, err := destructRoomUserKey(subscribedChannel)
		if err != nil {
			log.Println("getted invalid channel key from redis")
			continue
		}

		switch msg.Payload {
		case redis.EventSet:
			// dataをGet
			roomUserJSON, err := s.client.Get(ctx, subscribedChannel).Result()
			if err != nil {
				log.Println("failed to get from redis, err:", err)
				continue
			}
			// jsonUnmarshal
			var roomUser domain.RoomUser
			if err := json.Unmarshal([]byte(roomUserJSON), &roomUser); err != nil {
				log.Println("getted unexpected json data struct from redis")
				continue
			}
			// eventによる分岐
			switch roomUser.LastEvent {
			case domain.JoinEvent:
				data := &model.Joined{
					UserID:    makeRoomUserID(userID),
					Name:      roomUser.Name,
					AvatarURL: roomUser.AvatarURL,
					X:         roomUser.X,
					Y:         roomUser.Y,
				}
				s.deliver(roomID, data)
			case domain.MoveEvent:
				data := &model.Moved{
					UserID: makeRoomUserID(userID),
					X:      roomUser.X,
					Y:      roomUser.Y,
				}
				s.deliver(roomID, data)
			case domain.MessageEvent:
				data := &model.SendedMassage{
					UserID:  makeRoomUserID(userID),
					Message: roomUser.LastMessage,
				}
				s.deliver(roomID, data)
			default:
				log.Println("getted unknown roomUser event")
			}
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
		ch <- &data
	}
}

func makeRoomUserID(roomUserID int) string {
	return fmt.Sprintf("RoomUser:%d", roomUserID)
}
