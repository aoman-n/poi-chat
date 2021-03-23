package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/laster18/poi/api/graph/model"
)

var (
	// "room:{roomID}:{dataType}"
	subChFormat = "room:%s:%s"
)

var (
	typeMessagse = "message"
	typeMove     = "move"
)

type PubsubRepo struct {
	client *redis.Client
}

func NewPubsubRepo(client *redis.Client) *PubsubRepo {
	return &PubsubRepo{client}
}

type SubscribeChs struct {
	MessageCh chan *model.Message
	MoveCh    chan *model.MovedUser
}

func NewSubscribeChs() *SubscribeChs {
	return &SubscribeChs{
		MessageCh: make(chan *model.Message),
		MoveCh:    make(chan *model.MovedUser),
	}
}

func (r *PubsubRepo) PSub(ctx context.Context, roomID int, subChs *SubscribeChs) error {
	pubsub := r.client.PSubscribe(ctx, fmt.Sprintf(subChFormat, roomID, "*"))

	for {
		select {
		case <-ctx.Done():
			log.Printf("stop subscribe, roomId: %d", roomID)
			break
		default:
		}

		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			log.Println("failed to receive message from subsciribe redis, err:", err)
			continue
		}

		chName := msg.Channel
		dataType := strings.Split(chName, ":")[2]
		payload := msg.Payload

		switch dataType {
		case typeMessagse:
			var msg model.Message
			if err := json.Unmarshal([]byte(payload), &msg); err != nil {
				log.Println(`failed to convert data "message" from redis`)
				continue
			}

			subChs.MessageCh <- &msg
		case typeMove:
			var movedUser model.MovedUser
			if err := json.Unmarshal([]byte(payload), &movedUser); err != nil {
				log.Println(`failed to convert data "movedUser" from redis`)
				continue
			}

			subChs.MoveCh <- &movedUser
		// typeDeleteのときにreturnする
		default:
			log.Printf(
				"receive unknown data type message from subscribe redis, channel: %s, data: %s",
				chName,
				payload,
			)
		}
	}
}
