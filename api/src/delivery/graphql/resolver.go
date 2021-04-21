package graphql

import (
	"context"
	"fmt"
	"log"
	"strings"

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
	// subscripters.Add("Room:1")

	return &Resolver{
		roomRepo:     roomRepo,
		messageRepo:  messageRepo,
		pubsubRepo:   pubsubRepo,
		subscripters: subscripters,
		redisClient:  redis,
	}
}

type Resolver struct {
	roomRepo     domain.IRoomRepo
	messageRepo  domain.IMessageRepo
	pubsubRepo   *repository.PubsubRepo
	subscripters *Subscripters
	redisClient  *redis.Client
}

// onlineUser:*
// room:

var (
	userChKeyspaceFormat = "__keyspace@0__:onlineUser:%s"
	userChFormat         = "onlineUser:%s"
)

var (
	setEvent     = "set"
	expireEvent  = "expire"
	expiredEvent = "expired"
	delEvent     = "del"
)

func (r *Resolver) StartSubscribeUserStatus() {
	log.Println("start subscribe user status!!")

	ctx := context.Background()
	subChName := fmt.Sprintf(userChKeyspaceFormat, "*")
	pubsub := r.redisClient.PSubscribe(ctx, subChName)

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			log.Println("failed to receive message from subsciribe redis, err:", err)
		}

		eventType := msg.Payload

		switch eventType {
		case setEvent:
			ch := msg.Channel
			pos := strings.Index(ch, ":")
			key := ch[pos+1:]
			user := r.redisClient.Get(ctx, key)
			payload := user.Val()

			fmt.Printf(
				"getted, eventName: %s, key: %s, payload: %s\n",
				eventType, key, payload,
			)
			// ユーザーにオンラインユーザー情報を通知する
		case expiredEvent:
			// ユーザーにオフラインになったユーザー情報を通知する
		case delEvent:
			// ユーザーにオフラインになったユーザー情報を通知する
		default:
			log.Printf(`"%s" is unknown keyspace event type`, eventType)
		}
	}
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
		case msg := <-chs.UserEventCh:
			if s, ok := r.subscripters.Get(roomIDStr); ok {
				s.PublishUserEvent(msg)
			}
		case <-ctx.Done():
			return
		}
	}
}
