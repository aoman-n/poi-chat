package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"strconv"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/aerrors"
)

func (r *exitedPayloadResolver) UserID(ctx context.Context, obj *model.ExitedPayload) (string, error) {
	return encodeIDStr(roomUserPrefix, obj.UserID), nil
}

func (r *globalUserResolver) ID(ctx context.Context, obj *model.GlobalUser) (string, error) {
	return encodeIDStr(globalUserPrefix, obj.ID), nil
}

func (r *meResolver) ID(ctx context.Context, obj *model.Me) (string, error) {
	return encodeIDStr(globalUserPrefix, obj.ID), nil
}

func (r *mutationResolver) Move(ctx context.Context, input model.MoveInput) (*model.MovePayload, error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		handleErr(ctx, errUnauthorized)
		return nil, nil
	}

	domainRoomID, err := decodeID(roomPrefix, input.RoomID)
	if err != nil {
		handleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	roomUser, err := r.roomUserRepo.Get(ctx, domainRoomID, currentUser.UID)
	if err != nil {
		log.Println("failed to get roomUser err:", err)
	}
	if roomUser == nil {
		roomUser = domain.NewDefaultRoomUser(domainRoomID, currentUser)
	}
	roomUser.SetPosition(input.X, input.Y)

	if err := r.roomUserRepo.Insert(ctx, roomUser); err != nil {
		handleErr(ctx, aerrors.Wrap(err, "failed to roomUserRepo.Insert"))
		return nil, nil
	}

	return toMovePayload(roomUser), nil
}

func (r *offlinedPayloadResolver) UserID(ctx context.Context, obj *model.OfflinedPayload) (string, error) {
	return encodeIDStr(globalUserPrefix, obj.UserID), nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.Me, error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		return nil, errUnauthorized
	}

	me := &model.Me{
		ID:        currentUser.UID,
		Name:      currentUser.Name,
		AvatarURL: currentUser.AvatarURL,
	}

	return me, nil
}

func (r *queryResolver) GlobalUsers(ctx context.Context) ([]*model.GlobalUser, error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		handleErr(ctx, errUnauthorized)
		return nil, nil
	}

	users, err := r.globalUserRepo.GetAll(ctx)
	if err != nil {
		handleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	return toGlobalUsers(users), nil
}

func (r *roomResolver) Users(ctx context.Context, obj *model.Room) ([]*model.RoomUser, error) {
	roomID, _ := strconv.Atoi(obj.ID)

	users, err := r.roomUserRepo.GetByRoomID(ctx, roomID)
	if err != nil {
		handleErr(ctx, aerrors.Wrap(err, "failed to roomUserRepo.GetByRoomID"))
		return nil, nil
	}

	return toRoomUsers(users), nil
}

func (r *roomUserResolver) ID(ctx context.Context, obj *model.RoomUser) (string, error) {
	return encodeIDStr(roomUserPrefix, obj.ID), nil
}

func (r *subscriptionResolver) ActedGlobalUserEvent(ctx context.Context) (<-chan model.GlobalUserEvent, error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		handleErr(ctx, errUnauthorized)
		return nil, nil
	}

	ch := make(chan model.GlobalUserEvent)
	r.globalUserSubscriber.AddCh(ch, currentUser.UID)

	newGlobalUser := &domain.GlobalUser{
		UID:       currentUser.UID,
		Name:      currentUser.Name,
		AvatarURL: currentUser.AvatarURL,
	}
	if err := r.globalUserRepo.Insert(ctx, newGlobalUser); err != nil {
		handleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	go func() {
		<-ctx.Done()
		r.globalUserSubscriber.RemoveCh(currentUser.UID)
		if err := r.globalUserRepo.Delete(context.Background(), currentUser.UID); err != nil {
			log.Println("failed to delete globalUser, err:", err)
		}
	}()

	return ch, nil
}

func (r *subscriptionResolver) ActedRoomUserEvent(
	ctx context.Context,
	roomID string,
) (<-chan model.RoomUserEvent, error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		handleErr(ctx, errUnauthorized)
		return nil, nil
	}

	domainRoomID, err := decodeID(roomPrefix, roomID)
	if err != nil {
		handleErr(ctx, aerrors.Wrap(err, "roomId is invalid format"))
		return nil, nil
	}

	// TODO: roomの存在チェック

	ch := make(chan model.RoomUserEvent)
	r.roomUserSubscriber.AddCh(ch, domainRoomID, currentUser.UID)

	newRoomUser := domain.NewDefaultRoomUser(domainRoomID, currentUser)
	if err := r.roomUserRepo.Insert(ctx, newRoomUser); err != nil {
		handleErr(ctx, aerrors.Wrap(err, "failed to roomUserRepo.Insert"))
		return nil, nil
	}

	go func() {
		<-ctx.Done()
		r.roomUserSubscriber.RemoveCh(domainRoomID, currentUser.UID)
		if err := r.roomUserRepo.Delete(context.Background(), newRoomUser); err != nil {
			log.Println("failed to delete roomUser, err:", err)
		}
	}()

	return ch, nil
}

// ExitedPayload returns generated.ExitedPayloadResolver implementation.
func (r *Resolver) ExitedPayload() generated.ExitedPayloadResolver { return &exitedPayloadResolver{r} }

// GlobalUser returns generated.GlobalUserResolver implementation.
func (r *Resolver) GlobalUser() generated.GlobalUserResolver { return &globalUserResolver{r} }

// Me returns generated.MeResolver implementation.
func (r *Resolver) Me() generated.MeResolver { return &meResolver{r} }

// OfflinedPayload returns generated.OfflinedPayloadResolver implementation.
func (r *Resolver) OfflinedPayload() generated.OfflinedPayloadResolver {
	return &offlinedPayloadResolver{r}
}

// RoomUser returns generated.RoomUserResolver implementation.
func (r *Resolver) RoomUser() generated.RoomUserResolver { return &roomUserResolver{r} }

type exitedPayloadResolver struct{ *Resolver }
type globalUserResolver struct{ *Resolver }
type meResolver struct{ *Resolver }
type offlinedPayloadResolver struct{ *Resolver }
type roomUserResolver struct{ *Resolver }
