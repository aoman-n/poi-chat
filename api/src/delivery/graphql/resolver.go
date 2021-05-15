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

func NewResolver(db *gorm.DB, redis *redis.Client, roomUserSubscriber *subscriber.RoomUserSubscriber) *Resolver {
	roomRepo := repository.NewRoomRepo(db)
	messageRepo := repository.NewMessageRepo(db)
	roomUserRepo := repository.NewRoomUserRepo(db, redis)
	// pubsubRepo := repository.NewPubsubRepo(redis)
	// subscripters := NewSubscripters()
	// subscripterForAll := NewSubscripterForAll()

	return &Resolver{
		roomRepo:     roomRepo,
		messageRepo:  messageRepo,
		roomUserRepo: roomUserRepo,
		// pubsubRepo:  pubsubRepo,
		// subscripters:       subscripters,
		// subscripterForAll:  subscripterForAll,
		redisClient:        redis,
		roomUserSubscriber: roomUserSubscriber,
	}
}

type Resolver struct {
	roomRepo     domain.IRoomRepo
	messageRepo  domain.IMessageRepo
	roomUserRepo domain.IRoomUserRepo
	// pubsubRepo  *repository.PubsubRepo
	// subscripters       *Subscripters
	// subscripterForAll  *SubscripterForAll
	redisClient        *redis.Client
	roomUserSubscriber *subscriber.RoomUserSubscriber
}

// onlineUser:*
// room:

// var (
// 	userChKeyspaceFormat = "__keyspace@0__:onlineUser:%s"
// 	userChFormat         = "onlineUser:%s"
// )

// var (
// 	setEvent     = "set"
// 	expireEvent  = "expire"
// 	expiredEvent = "expired"
// 	delEvent     = "del"
// )

// var (
// 	userIDRegex = regexp.MustCompile("onlineUser:(.*)")
// )

// func (r *Resolver) StartSubscribeUserStatus() {
// 	log.Println("start subscribe user status!!")

// 	ctx := context.Background()
// 	subChName := fmt.Sprintf(userChKeyspaceFormat, "*")
// 	pubsub := r.redisClient.PSubscribe(ctx, subChName)

// 	for {
// 		msg, err := pubsub.ReceiveMessage(ctx)
// 		if err != nil {
// 			log.Println("failed to receive message from subsciribe redis, err:", err)
// 			return
// 		}

// 		eventType := msg.Payload

// 		switch eventType {
// 		case setEvent:
// 			pos := strings.Index(msg.Channel, ":")
// 			key := msg.Channel[pos+1:]
// 			user := r.redisClient.Get(ctx, key)
// 			payload := user.Val()

// 			fmt.Printf("setted payload: %+v\n\n", payload)

// 			var onlineUserStatus model.OnlineUserStatus
// 			if err := json.Unmarshal([]byte(payload), &onlineUserStatus); err != nil {
// 				log.Println("failed to unmarshal subscribe setted data")
// 				return
// 			}

// 			fmt.Printf("setted onlineUser: %+v\n\n", onlineUserStatus)

// 			// ユーザーにオンラインユーザー情報を通知する
// 			r.subscripterForAll.PublishUserStatus(onlineUserStatus)
// 			break
// 		case expiredEvent:
// 			fallthrough
// 		case delEvent:
// 			matched := userIDRegex.FindStringSubmatch(msg.Channel)
// 			userID := matched[1]
// 			offlineUserStatus := model.OfflineUserStatus{
// 				ID: userID,
// 			}

// 			fmt.Printf("deleted offlineUser: %+v\n\n", offlineUserStatus)

// 			// ユーザーにオフラインになったユーザー情報を通知する
// 			r.subscripterForAll.PublishUserStatus(offlineUserStatus)
// 		default:
// 			log.Printf(`"%s" is unknown keyspace event type`, eventType)
// 		}
// 	}
// }

// func (r *Resolver) SetupRoom(roomID int) {
// 	roomIDStr := encodeID(roomPrefix, roomID)

// 	ctx, cancel := context.WithCancel(context.Background())

// 	// subscripter作成・登録
// 	r.subscripters.Add(roomIDStr)

// 	// サブスクライブ開始
// 	chs := repository.NewSubscribeChs()
// 	go func() {
// 		err := r.pubsubRepo.PSub(ctx, roomID, chs)
// 		if err != nil {
// 			log.Println("failed to psubscribe err:", err)
// 		}
// 		cancel()
// 	}()

// 	for {
// 		select {
// 		case msg := <-chs.MessageCh:
// 			if s, ok := r.subscripters.Get(roomIDStr); ok {
// 				s.PublishMessage(msg)
// 			}
// 		case msg := <-chs.UserEventCh:
// 			if s, ok := r.subscripters.Get(roomIDStr); ok {
// 				s.PublishUserEvent(msg)
// 			}
// 		case <-ctx.Done():
// 			return
// 		}
// 	}
// }
