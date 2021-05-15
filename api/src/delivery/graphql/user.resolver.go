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
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
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
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
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

func (r *queryResolver) OnlineUsers(ctx context.Context) ([]*model.OnlineUserStatus, error) {
	// ret := r.redisClient.Keys(ctx, fmt.Sprintf(userChFormat, "*"))
	// userKeys, err := ret.Result()
	// if err != nil {
	// 	return nil, errUnexpected
	// }

	// ret2 := r.redisClient.MGet(ctx, userKeys...)
	// retOnlineUsers := ret2.Val()

	// onlineUsers := []*model.OnlineUserStatus{}
	// for _, user := range retOnlineUsers {
	// 	userStr := user.(string)
	// 	var u model.OnlineUserStatus
	// 	if err := json.Unmarshal([]byte(userStr), &u); err != nil {
	// 		log.Println("failed to unmarshal payload from redis data, data is", userStr)
	// 		continue
	// 	}
	// 	onlineUsers = append(onlineUsers, &u)
	// }

	// return onlineUsers, nil

	return nil, nil
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
			ID:        encodeIDStr(userPrefix, ju.UserID),
			Name:      ju.DisplayName,
			AvatarURL: ju.AvatarURL,
			X:         ju.X,
			Y:         ju.Y,
		}
	}

	return users, nil
}

func (r *subscriptionResolver) ActedGlobalUserEvent(ctx context.Context) (<-chan model.GlobalUserEvent, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
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
		if err := r.globalUserRepo.Delete(context.Background(), newGlobalUser); err != nil {
			log.Println("failed to delete globalUser, err:", err)
		}
	}()

	return ch, nil
}

func (r *subscriptionResolver) ActedRoomUserEvent(ctx context.Context, roomID string) (<-chan model.RoomUserEvent, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
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
