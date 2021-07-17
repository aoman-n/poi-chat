package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain/message"
	"github.com/laster18/poi/api/src/presentation/graphql"
	"github.com/laster18/poi/api/src/presentation/graphql/presenter"
	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/aerrors"
	"github.com/laster18/poi/api/src/util/clock"
)

func (r *messageResolver) ID(ctx context.Context, obj *model.Message) (string, error) {
	return graphql.MessageIDStr(obj.ID), nil
}

func (r *messageResolver) User(ctx context.Context, obj *model.Message) (*model.User, error) {
	uID, _ := strconv.Atoi(obj.User.ID)
	user, err := acontext.GetUserLoader(ctx).Load(uID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to userLoader.Load"))
		return nil, nil
	}

	return presenter.ToUser(user), nil
}

func (r *mutationResolver) SendMessage(
	ctx context.Context,
	input *model.SendMessageInput,
) (*model.SendMassagePaylaod, error) {
	currentUser := acontext.GetUser(ctx)
	domainRoomID, err := graphql.DecodeRoomID(input.RoomID)
	if err != nil {
		return nil, aerrors.Wrap(err)
	}

	msg := &message.Message{
		UserID:    currentUser.ID,
		Body:      input.Body,
		RoomID:    domainRoomID,
		CreatedAt: clock.Now(),
	}

	messageRepo := r.repo.NewMessage()
	if err := messageRepo.Create(ctx, msg); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to messageRepo.Create"))
		return nil, nil
	}

	roomRepo := r.repo.NewRoom()
	roomSvc := r.service.NewRoom()

	roomUserStatus, err := roomSvc.FindOrNewUserStatus(ctx, currentUser, domainRoomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomSvc.FindOrNewUserStatus"))
		return nil, nil
	}

	roomUserStatus.SetMessage(msg)

	if err := roomRepo.SaveUserStatus(ctx, roomUserStatus); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomUserRepo.Save"))
		return nil, nil
	}

	return presenter.ToSendMessagePayload(msg), nil
}

func (r *roomResolver) Messages(
	ctx context.Context,
	obj *model.Room,
	last *int,
	before *string,
) (*model.MessageConnection, error) {
	roomID, err := strconv.Atoi(obj.ID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	messageListReq := &message.ListReq{
		RoomID: roomID,
	}

	if last != nil {
		messageListReq.Limit = *last
	}
	if before != nil {
		id, unix, err := graphql.DecodeMessageCursor(before)
		if err != nil {
			graphql.HandleErr(ctx, aerrors.Wrap(err))
			return nil, nil
		}

		messageListReq.LastKnownID = id
		messageListReq.LastKnownUnix = unix
	}

	messageRepo := r.repo.NewMessage()

	messageListResp, err := messageRepo.List(ctx, messageListReq)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to messageRepo.List"))
		return nil, nil
	}

	totalCount, err := messageRepo.Count(ctx, roomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to messageRepo.Count"))
		return nil, nil
	}

	return presenter.ToMessageConnection(before, messageListResp, totalCount), nil
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

type messageResolver struct{ *Resolver }
