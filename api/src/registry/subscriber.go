package registry

import (
	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/presentation/graphql/subscriber"
)

type Subscriber interface {
	NewUser() *subscriber.GlobalUserSubscriber
	NewRoomUser() *subscriber.RoomUserSubscriber
}

type subscriberImpl struct {
	client   *redis.Client
	userRepo user.Repository
}

func NewSubscriber(c *redis.Client, userRepo user.Repository) Subscriber {
	return &subscriberImpl{c, userRepo}
}

func (s *subscriberImpl) NewUser() *subscriber.GlobalUserSubscriber {
	return subscriber.NewGlobalUserSubscriber(s.client, s.userRepo)
}

func (s *subscriberImpl) NewRoomUser() *subscriber.RoomUserSubscriber {
	return subscriber.NewRoomUserSubscriber(s.client)
}
