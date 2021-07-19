package subscriber

import (
	"context"
	"fmt"
	"sync"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/presentation/graphql/presenter"
	"github.com/laster18/poi/api/src/util/acontext"
)

// GlobalUserSubscriber: ユーザーのオンライン/オフライン状態の更新イベントを管理
type GlobalUserSubscriber struct {
	userRepo user.Repository
	client   *redis.Client
	mutex    sync.Mutex
	// chans map[userUID]chan ...
	chans map[string]chan<- model.UserEvent
}

func NewGlobalUserSubscriber(client *redis.Client, userRepo user.Repository) *GlobalUserSubscriber {
	subscriber := &GlobalUserSubscriber{
		userRepo: userRepo,
		client:   client,
		chans:    make(map[string]chan<- model.UserEvent),
	}
	return subscriber
}

func (s *GlobalUserSubscriber) Start(ctx context.Context) {
	logger := acontext.GetLogger(ctx)

	subscribeCh := fmt.Sprintf("%s:%s:%s", redis.KeySpace, OnlineUserChannel, "*")
	pubsub := s.client.PSubscribe(ctx, subscribeCh)
	defer pubsub.Close()

	for {
		select {
		case <-ctx.Done():
			logger.Info("stop user subscriber")
			return
		case msg := <-pubsub.Channel():
			logger.Debugf("subscribe user onlined or offlined event, channel: %s, payload: %s", msg.Channel, msg.Payload)

			ch := removeKeyspacePrefix(msg.Channel)
			userID, err := DestructOnlineUserKey(ch)
			if err != nil {
				logger.Warnf("received invalid channel key from redis, err: %v", err)
				continue
			}

			switch msg.Payload {
			// onlineになった
			case redis.EventSet:
				user, err := s.userRepo.Get(ctx, userID)
				if err != nil {
					logger.Warnf("failed to get user on subscriber, err: %v", err)
					continue
				}
				s.deliver(presenter.ToOnlinedPayload(user))
			// offlineになった
			case redis.EventDel:
				fallthrough
			// offlineになった
			case redis.EventExpired:
				user, err := s.userRepo.Get(ctx, userID)
				if err != nil {
					logger.Warnf("failed to get user on subscriber, err: %v", err)
					continue
				}
				s.deliver(presenter.ToOfflinedPayload(user))
			default:
				logger.Infof("received unknown event: %s", msg.Payload)
			}
		}
	}
}

func (s *GlobalUserSubscriber) deliver(data model.UserEvent) {
	for _, ch := range s.chans {
		ch <- data
	}
}

func (s *GlobalUserSubscriber) AddCh(ch chan<- model.UserEvent, userUID string) {
	s.mutex.Lock()
	s.chans[userUID] = ch
	s.mutex.Unlock()
}

func (s *GlobalUserSubscriber) RemoveCh(userUID string) {
	s.mutex.Lock()
	delete(s.chans, userUID)
	s.mutex.Unlock()
}
