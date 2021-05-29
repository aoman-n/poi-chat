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

func (r *exitedResolver) UserID(ctx context.Context, obj *model.Exited) (string, error) {
	return encodeIDStr(userPrefix, obj.UserID), nil
}

func (r *meResolver) ID(ctx context.Context, obj *model.Me) (string, error) {
	return encodeIDStr(userPrefix, obj.ID), nil
}

func (r *movePayloadResolver) UserID(ctx context.Context, obj *model.MovePayload) (string, error) {
	return encodeIDStr(userPrefix, obj.UserID), nil
}

func (r *mutationResolver) Move(ctx context.Context, input model.MoveInput) (*model.MovePayload, error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		return nil, errUnauthorized
	}

	domainRoomID, err := decodeID(roomPrefix, input.RoomID)
	if err != nil {
		return nil, aerrors.Wrap(err)
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
		return nil, aerrors.Wrap(err, "failed to roomUserRepo.Insert")
	}

	return toMovePayload(roomUser), nil
}

func (r *offlinedResolver) UserID(ctx context.Context, obj *model.Offlined) (string, error) {
	return encodeIDStr(userPrefix, obj.UserID), nil
}

func (r *onlineUserResolver) ID(ctx context.Context, obj *model.OnlineUser) (string, error) {
	return encodeIDStr(userPrefix, obj.ID), nil
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

func (r *queryResolver) OnlineUsers(ctx context.Context) ([]*model.OnlineUser, error) {
	globalUsers, err := r.globalUserRepo.GetAll(ctx)
	if err != nil {
		return nil, aerrors.Wrap(err, "failed to globalUserRepo.GetAll")
	}

	return toOnlineUsers(globalUsers), nil
}

func (r *roomResolver) Users(ctx context.Context, obj *model.Room) ([]*model.RoomUser, error) {
	roomID, _ := strconv.Atoi(obj.ID)

	users, err := r.roomUserRepo.GetByRoomID(ctx, roomID)
	if err != nil {
		return nil, aerrors.Wrap(err, "failed to roomUserRepo.GetByRoomID")
	}

	return toRoomUsers(users), nil
}

func (r *roomUserResolver) ID(ctx context.Context, obj *model.RoomUser) (string, error) {
	return encodeIDStr(roomUserPrefix, obj.ID), nil
}

func (r *subscriptionResolver) ActedGlobalUserEvent(ctx context.Context) (<-chan model.GlobalUserEvent, error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		return nil, errUnauthorized
	}

	newGlobalUser := &domain.GlobalUser{
		UID:       currentUser.UID,
		Name:      currentUser.Name,
		AvatarURL: currentUser.AvatarURL,
	}
	if err := r.globalUserRepo.Insert(ctx, newGlobalUser); err != nil {
		return nil, err
	}

	ch := make(chan model.GlobalUserEvent)
	r.globalUserSubscriber.AddCh(ch, currentUser.UID)

	go func() {
		<-ctx.Done()
		r.globalUserSubscriber.RemoveCh(currentUser.UID)
		if err := r.globalUserRepo.Delete(context.Background(), currentUser.UID); err != nil {
			log.Println("failed to delete globalUser, err:", err)
		}
	}()

	return ch, nil
}

func (r *subscriptionResolver) ActedRoomUserEvent(ctx context.Context, roomID string) (<-chan model.RoomUserEvent, error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		return nil, errUnauthorized
	}

	domainRoomID, err := decodeID(roomPrefix, roomID)
	if err != nil {
		return nil, aerrors.Wrap(err, "roomId is invalid format")
	}

	// TODO: roomの存在チェック

	newRoomUser := domain.NewDefaultRoomUser(domainRoomID, currentUser)
	if err := r.roomUserRepo.Insert(ctx, newRoomUser); err != nil {
		return nil, aerrors.Wrap(err, "failed to roomUserRepo.Insert")
	}

	ch := make(chan model.RoomUserEvent)
	r.roomUserSubscriber.AddCh(ch, domainRoomID, currentUser.UID)

	go func() {
		<-ctx.Done()
		r.roomUserSubscriber.RemoveCh(domainRoomID, currentUser.UID)
		if err := r.roomUserRepo.Delete(context.Background(), newRoomUser); err != nil {
			log.Println("failed to delete roomUser, err:", err)
		}
	}()

	return ch, nil
}

// Exited returns generated.ExitedResolver implementation.
func (r *Resolver) Exited() generated.ExitedResolver { return &exitedResolver{r} }

// Me returns generated.MeResolver implementation.
func (r *Resolver) Me() generated.MeResolver { return &meResolver{r} }

// MovePayload returns generated.MovePayloadResolver implementation.
func (r *Resolver) MovePayload() generated.MovePayloadResolver { return &movePayloadResolver{r} }

// Offlined returns generated.OfflinedResolver implementation.
func (r *Resolver) Offlined() generated.OfflinedResolver { return &offlinedResolver{r} }

// OnlineUser returns generated.OnlineUserResolver implementation.
func (r *Resolver) OnlineUser() generated.OnlineUserResolver { return &onlineUserResolver{r} }

// RoomUser returns generated.RoomUserResolver implementation.
func (r *Resolver) RoomUser() generated.RoomUserResolver { return &roomUserResolver{r} }

type exitedResolver struct{ *Resolver }
type meResolver struct{ *Resolver }
type movePayloadResolver struct{ *Resolver }
type offlinedResolver struct{ *Resolver }
type onlineUserResolver struct{ *Resolver }
type roomUserResolver struct{ *Resolver }
