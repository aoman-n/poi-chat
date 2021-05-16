package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/middleware"
	"github.com/laster18/poi/api/src/util/clock"
)

func (r *messageResolver) ID(ctx context.Context, obj *model.Message) (string, error) {
	return fmt.Sprintf(messageIDFormat, obj.ID), nil
}

func (r *mutationResolver) SendMessage(ctx context.Context, input *model.SendMessageInput) (*model.Message, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}

	domainRoomID, err := decodeID(roomPrefix, input.RoomID)
	if err != nil {
		return nil, err
	}

	msg := &domain.Message{
		UserUID:       currentUser.UID,
		Body:          input.Body,
		UserName:      currentUser.Name,
		UserAvatarURL: currentUser.AvatarURL,
		RoomID:        domainRoomID,
		CreatedAt:     clock.Now(),
	}

	if err := r.messageRepo.Create(ctx, msg); err != nil {
		log.Println("create message error:", err)
		return nil, errUnexpected
	}

	roomUser, err := r.roomUserRepo.Get(ctx, domainRoomID, currentUser.UID)
	if err != nil {
		log.Println("failed to get roomUser, err:", err)
		return nil, err
	}
	if roomUser == nil {
		roomUser = domain.NewDefaultRoomUser(domainRoomID, currentUser)
	}
	roomUser.SetMessage(msg)

	if err := r.roomUserRepo.Insert(ctx, roomUser); err != nil {
		log.Println("failed to insert roomUser, err:", err)
		return nil, err
	}

	return toMessage(msg), nil
}

func (r *roomDetailResolver) Messages(ctx context.Context, obj *model.RoomDetail, last *int, before *string) (*model.MessageConnection, error) {
	_, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}

	roomID, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, err
	}

	messageListReq := &domain.MessageListReq{
		RoomID: roomID,
	}

	if last != nil {
		messageListReq.Limit = *last
	}
	if before != nil {
		id, unix, err := decodeCursor(messagePrefix, before)
		if err != nil {
			return nil, err
		}

		messageListReq.LastKnownID = id
		messageListReq.LastKnownUnix = unix
	}

	messageListResp, err := r.messageRepo.List(ctx, messageListReq)
	if err != nil {
		return nil, err
	}
	messageCount, err := r.messageRepo.Count(ctx, roomID)
	if err != nil {
		return nil, err
	}

	// create pageInfo
	var pageInfo model.PageInfo
	pageInfo.HasPreviousPage = messageListResp.HasPreviousPage
	if before != nil {
		pageInfo.HasNextPage = true
	} else {
		pageInfo.HasNextPage = false
	}
	startCursor, endCursor := getMessageCursors(messageListResp.List)
	pageInfo.StartCursor = startCursor
	pageInfo.EndCursor = endCursor

	// create nodes, serializing message model
	nodes := make([]*model.Message, len(messageListResp.List))
	for i, message := range messageListResp.List {
		nodes[i] = toMessage(message)
	}

	// create edges
	edges := make([]*model.MessageEdge, len(messageListResp.List))
	for i, message := range messageListResp.List {
		edges[i] = &model.MessageEdge{
			Cursor: *encodeCursor(messagePrefix, message.GetID(), message.GetCreatedAtUnix()),
			Node:   nodes[i],
		}
	}

	return &model.MessageConnection{
		PageInfo:     &pageInfo,
		Nodes:        nodes,
		Edges:        edges,
		MessageCount: messageCount,
	}, nil
}

func (r *subscriptionResolver) SubMessage(ctx context.Context, roomID string) (<-chan *model.Message, error) {
	fmt.Println("start subscribe message")

	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}

	subscripter, ok := r.subscripters.Get(roomID)
	if !ok {
		return nil, errRoomNotFound
	}

	ch := make(chan *model.Message, 1)
	subscripter.AddMessageChan(currentUser.UID, ch)

	go func() {
		<-ctx.Done()
		fmt.Println("stop subscribe message")
		subscripter.DeleteMessageChan(currentUser.UID)
	}()

	return ch, nil
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

type messageResolver struct{ *Resolver }
