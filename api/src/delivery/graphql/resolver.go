package graphql

import (
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/repository"
	"github.com/laster18/poi/api/src/subscriber"
	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(
	db *gorm.DB,
	redis *redis.Client,
	roomUserSubscriber *subscriber.RoomUserSubscriber,
	globalUserSubscriber *subscriber.GlobalUserSubscriber,
) *Resolver {
	roomRepo := repository.NewRoomRepo(db)
	messageRepo := repository.NewMessageRepo(db)
	roomUserRepo := repository.NewRoomUserRepo(db, redis)
	globalUserRepo := repository.NewGlobalUserRepo(redis)
	pubsubRepo := repository.NewPubsubRepo(redis)
	subscripters := NewSubscripters()

	return &Resolver{
		roomRepo:             roomRepo,
		messageRepo:          messageRepo,
		roomUserRepo:         roomUserRepo,
		globalUserRepo:       globalUserRepo,
		pubsubRepo:           pubsubRepo,
		subscripters:         subscripters,
		redisClient:          redis,
		roomUserSubscriber:   roomUserSubscriber,
		globalUserSubscriber: globalUserSubscriber,
	}
}

type Resolver struct {
	roomRepo             domain.IRoomRepo
	messageRepo          domain.IMessageRepo
	roomUserRepo         domain.IRoomUserRepo
	globalUserRepo       domain.GlobalUserRepo
	pubsubRepo           *repository.PubsubRepo
	subscripters         *Subscripters
	redisClient          *redis.Client
	roomUserSubscriber   *subscriber.RoomUserSubscriber
	globalUserSubscriber *subscriber.GlobalUserSubscriber
}
