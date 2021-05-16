package subscriber

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infra/redis"
)

// GlobalUserSubscriber: ユーザーのオンライン/オフライン状態の更新イベントを管理
type GlobalUserSubscriber struct {
	globalUserRepo domain.GlobalUserRepo
	client         *redis.Client
	mutex          sync.Mutex
	// chs map[userUID]chan ...
	chs map[string]chan<- model.GlobalUserEvent
}

func NewGlobalUserSubscriber(
	ctx context.Context,
	client *redis.Client,
	globalUserRepo domain.GlobalUserRepo,
) *GlobalUserSubscriber {
	subscriber := &GlobalUserSubscriber{
		globalUserRepo: globalUserRepo,
		client:         client,
		mutex:          sync.Mutex{},
		chs:            make(map[string]chan<- model.GlobalUserEvent),
	}
	go subscriber.start(ctx)
	return subscriber
}

func (s *GlobalUserSubscriber) start(ctx context.Context) {
	subscribeCh := fmt.Sprintf("%s:%s:%s", redis.KeySpace, GlobalUserChannel, "*")
	pubsub := s.client.PSubscribe(ctx, subscribeCh)
	defer pubsub.Close()

	for {
		msg := <-pubsub.Channel()

		// debug log
		log.Printf("subscribe globalUser, channel: %s, payload: %s\n\n", msg.Channel, msg.Payload)

		ch := removeKeyspacePrefix(msg.Channel)
		userUID, err := destructGlobalUserKey(ch)
		if err != nil {
			log.Println("received invalid channel key from redis, err:", err)
			continue
		}

		switch msg.Payload {
		case redis.EventSet:
			// onlineになった
			globalUser, err := s.globalUserRepo.Get(ctx, userUID)
			if err != nil {
				log.Println("failed to get globa user, err:", err)
				continue
			}
			if globalUser == nil {
				log.Printf("global user uid=%q is not found", userUID)
				continue
			}

			s.deliver(&model.Onlined{
				OnlineUser: &model.OnlineUser{
					ID:        makeUserID(globalUser.UID),
					Name:      globalUser.Name,
					AvatarURL: globalUser.AvatarURL,
				},
			})
		case redis.EventDel:
			// offlineになった
			fallthrough
		case redis.EventExpired:
			// offlineになった
			err := s.globalUserRepo.Delete(ctx, userUID)
			if err != nil {
				log.Printf("failed to delete global user, err: %v", err)
			}
			s.deliver(&model.Offlined{
				UserID: makeUserID(userUID),
			})
		default:
			fmt.Println("received unknown event:", msg.Payload)
		}
	}
}

func (s *GlobalUserSubscriber) deliver(data model.GlobalUserEvent) {
	for _, ch := range s.chs {
		ch <- data
	}
}

func (s *GlobalUserSubscriber) AddCh(ch chan<- model.GlobalUserEvent, userUID string) {
	s.mutex.Lock()
	s.chs[userUID] = ch
	s.mutex.Unlock()
}

func (s *GlobalUserSubscriber) RemoveCh(userUID string) {
	s.mutex.Lock()
	delete(s.chs, userUID)
	s.mutex.Unlock()
}
