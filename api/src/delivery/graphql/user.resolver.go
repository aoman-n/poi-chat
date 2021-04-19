package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/middleware"
)

func (r *mutationResolver) Move(ctx context.Context, input model.MoveInput) (*model.MovedUser, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}

	domainRoomID, err := decodeID(roomPrefix, input.RoomID)
	if err != nil {
		return nil, err
	}

	movedUser := &model.MovedUser{
		ID: currentUser.ID,
		X:  input.X,
		Y:  input.Y,
	}

	if err := r.pubsubRepo.PubMovedUser(ctx, movedUser, domainRoomID); err != nil {
		return nil, err
	}

	// userの位置の更新は最悪失敗しても良いので非同期で行う
	go func() {
		ctx2 := context.Background()
		domainUserID, err := decodeID(userPrefix, currentUser.ID)
		if err != nil {
			log.Println("failed to decode user id, err:", err)
			return
		}

		u, err := r.roomRepo.GetUserByID(ctx2, domainUserID)
		if err != nil {
			log.Println("failed to get user, err:", err)
			return
		}

		u.X = input.X
		u.Y = input.Y

		if err := r.roomRepo.UpdateUser(ctx, u); err != nil {
			log.Println("failed to update user position, err:", err)
			return
		}
	}()

	return movedUser, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.Me, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}

	me := &model.Me{
		ID:          currentUser.ID,
		DisplayName: currentUser.Name,
		AvatarURL:   currentUser.AvatarURL,
	}

	return me, nil
}

func (r *roomDetailResolver) Users(ctx context.Context, obj *model.RoomDetail) ([]*model.User, error) {
	id, _ := strconv.Atoi(obj.ID)

	joinedUsers, err := r.roomRepo.GetUsers(ctx, id)
	if err != nil {
		log.Println("failed to list joinedUser err:", err)
		return nil, errUnexpected
	}

	// serialize
	users := make([]*model.User, len(joinedUsers))
	for i, ju := range joinedUsers {
		users[i] = &model.User{
			ID:          encodeIDStr(userPrefix, ju.UserID),
			DisplayName: ju.DisplayName,
			AvatarURL:   ju.AvatarURL,
			X:           ju.X,
			Y:           ju.Y,
		}
	}

	return users, nil
}

func (r *subscriptionResolver) SubUserEvent(ctx context.Context, roomID string) (<-chan model.UserEvent, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}

	subscripter, ok := r.subscripters.Get(roomID)
	if !ok {
		return nil, errRoomNotFound
	}

	ch := make(chan model.UserEvent, 1)
	subscripter.AddUserEventChan(currentUser.ID, ch)

	go func() {
		<-ctx.Done()
		log.Println("stop subscribe user")
		subscripter.DeleteUserEventChan(currentUser.ID)
	}()

	return ch, nil
}

func (r *subscriptionResolver) JoinRoom(ctx context.Context, roomID string) (<-chan *model.User, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}

	domainRoomID, err := decodeID(roomPrefix, roomID)
	if err != nil {
		return nil, err
	}

	domainUserID, err := decodeID(userPrefix, currentUser.ID)
	if err != nil {
		return nil, err
	}

	_, ok := r.subscripters.Get(roomID)
	if !ok {
		return nil, errRoomNotFound
	}

	domainJoinedUser := &domain.JoinedUser{
		RoomID:      domainRoomID,
		AvatarURL:   currentUser.AvatarURL,
		DisplayName: currentUser.Name,
		UserID:      strconv.Itoa(domainUserID),
		// set default position
		X: 100,
		Y: 100,
	}

	if err := r.roomRepo.Join(ctx, domainJoinedUser); err != nil {
		log.Println("failed to create joinedUser, err:", err)
		return nil, errUnexpected
	}

	joinedUser := &model.JoinedUser{
		ID:          currentUser.ID,
		DisplayName: domainJoinedUser.DisplayName,
		AvatarURL:   domainJoinedUser.AvatarURL,
		X:           domainJoinedUser.X,
		Y:           domainJoinedUser.Y,
	}

	if err := r.pubsubRepo.PubJoinedUser(ctx, joinedUser, domainRoomID); err != nil {
		log.Println("failed to publish joined user err:", err)
		return nil, err
	}

	go func() {
		<-ctx.Done()
		childCtx := context.Background()
		if err := r.roomRepo.Exit(childCtx, domainJoinedUser); err != nil {
			// TODO: retry process
			log.Println("failed to delete joinedUser err:", err)
		}

		exitedUser := &model.ExitedUser{
			ID: joinedUser.ID,
		}

		fmt.Println(strings.Repeat("*", 100))
		fmt.Println(exitedUser)

		if err := r.pubsubRepo.PubExitedUser(childCtx, exitedUser, domainRoomID); err != nil {
			log.Println("failed to publish exited user err:", err)
		}

	}()

	ch := make(chan *model.User)

	return ch, nil
}
