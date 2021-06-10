package subscriber

import (
	"context"
	"fmt"
	"sync"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/util/acontext"
)

// GlobalUserSubscriber: ユーザーのオンライン/オフライン状態の更新イベントを管理
type GlobalUserSubscriber struct {
	globalUserRepo domain.GlobalUserRepo
	client         *redis.Client
	mutex          sync.Mutex
	// chans map[userUID]chan ...
	chans map[string]chan<- model.GlobalUserEvent
}

func NewGlobalUserSubscriber(
	ctx context.Context,
	client *redis.Client,
	globalUserRepo domain.GlobalUserRepo,
) *GlobalUserSubscriber {
	subscriber := &GlobalUserSubscriber{
		globalUserRepo: globalUserRepo,
		client:         client,
		chans:          make(map[string]chan<- model.GlobalUserEvent),
	}
	go subscriber.start(ctx)
	return subscriber
}

func (s *GlobalUserSubscriber) start(ctx context.Context) {
	logger := acontext.GetLogger(ctx)

	subscribeCh := fmt.Sprintf("%s:%s:%s", redis.KeySpace, GlobalUserChannel, "*")
	pubsub := s.client.PSubscribe(ctx, subscribeCh)
	defer pubsub.Close()

	for {
		msg := <-pubsub.Channel()

		// debug log
		logger.Debugf("subscribe globalUser, channel: %s, payload: %s", msg.Channel, msg.Payload)

		ch := removeKeyspacePrefix(msg.Channel)
		userUID, err := destructGlobalUserKey(ch)
		if err != nil {
			logger.Infof("received invalid channel key from redis, err: %v", err)
			continue
		}

		switch msg.Payload {
		case redis.EventSet:
			// onlineになった
			globalUser, err := s.globalUserRepo.Get(ctx, userUID)
			if err != nil {
				logger.Infof("failed to get global user, err: %v", err)
				continue
			}
			if globalUser == nil {
				logger.Infof("global user uid=%q is not found", userUID)
				continue
			}

			s.deliver(&model.OnlinedPayload{
				GlobalUser: &model.GlobalUser{
					ID:        globalUser.UID,
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
				logger.Infof("failed to delete global user, err: %v", err)
			}
			s.deliver(&model.OfflinedPayload{
				UserID: userUID,
			})
		default:
			logger.Infof("received unknown event: %s", msg.Payload)
		}
	}
}

func (s *GlobalUserSubscriber) deliver(data model.GlobalUserEvent) {
	for _, ch := range s.chans {
		ch <- data
	}
}

func (s *GlobalUserSubscriber) AddCh(ch chan<- model.GlobalUserEvent, userUID string) {
	s.mutex.Lock()
	s.chans[userUID] = ch
	s.mutex.Unlock()
}

func (s *GlobalUserSubscriber) RemoveCh(userUID string) {
	s.mutex.Lock()
	delete(s.chans, userUID)
	s.mutex.Unlock()
}
