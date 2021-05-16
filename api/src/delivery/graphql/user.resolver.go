package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/middleware"
)

func (r *mutationResolver) Move(ctx context.Context, input model.MoveInput) (*model.MovePayload, error) {
	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}

	domainRoomID, err := decodeID(roomPrefix, input.RoomID)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return toMovePayload(roomUser), nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.Me, error) {
	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}

	me := &model.Me{
		ID:        encodeIDStr(userPrefix, currentUser.UID),
		Name:      currentUser.Name,
		AvatarURL: currentUser.AvatarURL,
	}

	return me, nil
}

func (r *queryResolver) OnlineUsers(ctx context.Context) ([]*model.OnlineUser, error) {
	globalUsers, err := r.globalUserRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return toOnlineUsers(globalUsers), nil
}

func (r *roomResolver) Users(ctx context.Context, obj *model.Room) ([]*model.RoomUser, error) {
	roomID, _ := strconv.Atoi(obj.ID)

	users, err := r.roomUserRepo.GetByRoomID(ctx, roomID)
	if err != nil {
		return nil, err
	}

	return toRoomUsers(users), nil
}

func (r *subscriptionResolver) ActedGlobalUserEvent(ctx context.Context) (<-chan model.GlobalUserEvent, error) {
	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		return nil, errUnauthenticated
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
	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}

	domainRoomID, err := decodeID(roomPrefix, roomID)
	if err != nil {
		return nil, errors.New("roomId is invalid format")
	}

	// TODO: roomの存在チェック

	newRoomUser := domain.NewDefaultRoomUser(domainRoomID, currentUser)
	if err := r.roomUserRepo.Insert(ctx, newRoomUser); err != nil {
		return nil, err
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
