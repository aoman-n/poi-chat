package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/middleware"
)

func (r *messageResolver) ID(ctx context.Context, obj *model.Message) (string, error) {
	return fmt.Sprintf(messageIDFormat, obj.ID), nil
}

func (r *roomDetailResolver) Messages(ctx context.Context, obj *model.RoomDetail, last *int, before *string) (*model.MessageConnection, error) {
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
		nodes[i] = &model.Message{
			ID:            strconv.Itoa(message.ID),
			UserName:      message.UserName,
			UserAvatarURL: message.UserAvatarURL,
			Body:          message.Body,
			CreatedAt:     message.CreatedAt,
		}
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
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}

	subscripter, ok := r.subscripters.Get(roomID)
	if !ok {
		return nil, errNotFoundRoom
	}

	ch := make(chan *model.Message, 1)
	subscripter.AddMessageChan(currentUser.ID, ch)

	go func() {
		<-ctx.Done()
		subscripter.DeleteMessageChan(currentUser.ID)
	}()

	return ch, nil
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

type messageResolver struct{ *Resolver }
