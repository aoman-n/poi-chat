package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/infra/redis"
)

var (
	// "room:{roomID}:{dataType}"
	subChFormat = "room:%s:%s"
)

var (
	typeMessagse = "message"
	typeMove     = "move"
	typeJoin     = "join"
	typeExit     = "exit"
)

type PubsubRepo struct {
	client *redis.Client
}

func NewPubsubRepo(client *redis.Client) *PubsubRepo {
	return &PubsubRepo{client}
}

type SubscribeChs struct {
	MessageCh chan *model.Message
}

func NewSubscribeChs() *SubscribeChs {
	return &SubscribeChs{
		MessageCh: make(chan *model.Message),
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

			fmt.Println(strings.Repeat("*", 100))
			fmt.Println(msg)

			subChs.MessageCh <- &msg
		// case typeMove:
		// 	var movedUser model.MovedUser
		// 	if err := json.Unmarshal([]byte(payload), &movedUser); err != nil {
		// 		log.Println(`failed to convert data "movedUser" from redis`)
		// 		continue
		// 	}

		// 	subChs.UserEventCh <- &movedUser
		// case typeJoin:
		// 	var joinedUser model.JoinedUser
		// 	if err := json.Unmarshal([]byte(payload), &joinedUser); err != nil {
		// 		log.Println(`failed to convert data "joinedUser" from redis`)
		// 		continue
		// 	}

		// 	subChs.UserEventCh <- &joinedUser
		// case typeExit:
		// 	var exitedUser model.ExitedUser
		// 	if err := json.Unmarshal([]byte(payload), &exitedUser); err != nil {
		// 		log.Println(`failed to convert data "exitedUser" from redis`)
		// 		continue
		// 	}

		// 	subChs.UserEventCh <- &exitedUser
		// Todo: typeDeleteのときにreturnする
		default:
			log.Printf(
				"receive unknown data type message from subscribe redis, channel: %s, data: %s",
				chName,
				payload,
			)
		}
	}
}

func (r *PubsubRepo) PubMessage(ctx context.Context, m *model.Message, roomID int) error {
	channelName := fmt.Sprintf(subChFormat, roomID, typeMessagse)

	payload, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return r.client.Publish(ctx, channelName, payload).Err()
}

// func (r *PubsubRepo) PubJoinedUser(ctx context.Context, u *model.JoinedUser, roomID int) error {
// 	channelName := fmt.Sprintf(subChFormat, roomID, typeJoin)

// 	payload, err := json.Marshal(u)
// 	if err != nil {
// 		return err
// 	}

// 	return r.client.Publish(ctx, channelName, payload).Err()
// }

// func (r *PubsubRepo) PubExitedUser(ctx context.Context, u *model.ExitedUser, roomID int) error {
// 	channelName := fmt.Sprintf(subChFormat, roomID, typeExit)

// 	payload, err := json.Marshal(u)
// 	if err != nil {
// 		return err
// 	}

// 	return r.client.Publish(ctx, channelName, payload).Err()
// }

// func (r *PubsubRepo) PubMovedUser(ctx context.Context, u *model.MovedUser, roomID int) error {
// 	channelName := fmt.Sprintf(subChFormat, roomID, typeMove)

// 	payload, err := json.Marshal(u)
// 	if err != nil {
// 		return err
// 	}

// 	return r.client.Publish(ctx, channelName, payload).Err()
// }
