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
	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/aerrors"
	"github.com/laster18/poi/api/src/util/clock"
)

func (r *messageResolver) ID(ctx context.Context, obj *model.Message) (string, error) {
	return fmt.Sprintf(messageIDFormat, obj.ID), nil
}

func (r *messageResolver) UserID(ctx context.Context, obj *model.Message) (string, error) {
	return encodeIDStr(userPrefix, obj.UserID), nil
}

func (r *mutationResolver) SendMessage(ctx context.Context, input *model.SendMessageInput) (*model.SendMassagePaylaod, error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		return nil, errUnauthorized
	}

	domainRoomID, err := decodeID(roomPrefix, input.RoomID)
	if err != nil {
		return nil, aerrors.Wrap(err)
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
		return nil, aerrors.Wrap(err, "failed to messageRepo.Create")
	}

	roomUser, err := r.roomUserRepo.Get(ctx, domainRoomID, currentUser.UID)
	if err != nil {
		return nil, aerrors.Wrap(err, "failed to roomUserRepo.Get")
	}
	if roomUser == nil {
		roomUser = domain.NewDefaultRoomUser(domainRoomID, currentUser)
	}
	roomUser.SetMessage(msg)

	if err := r.roomUserRepo.Insert(ctx, roomUser); err != nil {
		return nil, aerrors.Wrap(err, "failed to roomUserRepo.Insert")
	}

	return toSendMessagePayload(msg), nil
}

func (r *roomResolver) Messages(ctx context.Context, obj *model.Room, last *int, before *string) (*model.MessageConnection, error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		return nil, errUnauthorized
	}

	roomID, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, aerrors.Wrap(err)
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
			return nil, aerrors.Wrap(err)
		}

		messageListReq.LastKnownID = id
		messageListReq.LastKnownUnix = unix
	}

	messageListResp, err := r.messageRepo.List(ctx, messageListReq)
	if err != nil {
		return nil, aerrors.Wrap(err, "failed to messageRepo.List")
	}

	totalCount, err := r.messageRepo.Count(ctx, roomID)
	if err != nil {
		return nil, aerrors.Wrap(err, "failed to messageRepo.Count")
	}

	return toMessageConnection(before, messageListResp, totalCount), nil
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

type messageResolver struct{ *Resolver }
