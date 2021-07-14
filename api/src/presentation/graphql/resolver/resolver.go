package resolver

import (
	"github.com/laster18/poi/api/src/presentation/graphql/subscriber"
	"github.com/laster18/poi/api/src/registry"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func New(
	repo registry.Repository,
	service registry.Service,
	roomUserSubscriber *subscriber.RoomUserSubscriber,
	globalUserSubscriber *subscriber.GlobalUserSubscriber,
) *Resolver {
	return &Resolver{
		repo:                 repo,
		service:              service,
		roomUserSubscriber:   roomUserSubscriber,
		globalUserSubscriber: globalUserSubscriber,
	}
}

type Resolver struct {
	repo                 registry.Repository
	service              registry.Service
	roomUserSubscriber   *subscriber.RoomUserSubscriber
	globalUserSubscriber *subscriber.GlobalUserSubscriber
}
