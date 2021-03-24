package graphql

import (
	"context"
	"log"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/repository"
	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(db *gorm.DB, redis *redis.Client) *Resolver {
	roomRepo := repository.NewRoomRepo(db)
	messageRepo := repository.NewMessageRepo(db)
	pubsubRepo := repository.NewPubsubRepo(redis)
	subscripters := NewSubscripters()

	// add mock room
	subscripters.Add("Room:1")

	return &Resolver{
		roomRepo:     roomRepo,
		messageRepo:  messageRepo,
		pubsubRepo:   pubsubRepo,
		subscripters: subscripters,
	}
}

type Resolver struct {
	roomRepo     domain.IRoomRepo
	messageRepo  domain.IMessageRepo
	pubsubRepo   *repository.PubsubRepo
	subscripters *Subscripters
}

func (r *Resolver) SetupRoom(roomID int) {
	roomIDStr := encodeID(roomPrefix, roomID)

	ctx, cancel := context.WithCancel(context.Background())

	// subscripter作成・登録
	r.subscripters.Add(roomIDStr)

	// サブスクライブ開始
	chs := repository.NewSubscribeChs()
	go func() {
		err := r.pubsubRepo.PSub(ctx, roomID, chs)
		if err != nil {
			log.Println("failed to psubscribe err:", err)
		}
		cancel()
	}()

	for {
		select {
		case msg := <-chs.MessageCh:
			if s, ok := r.subscripters.Get(roomIDStr); ok {
				s.PublishMessage(msg)
			}
		case msg := <-chs.MoveCh:
			if s, ok := r.subscripters.Get(roomIDStr); ok {
				s.PublishUserEvent(msg)
			}
		case <-ctx.Done():
			return
		}
	}
}
