package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/middleware"
)

func (r *mutationResolver) Move(ctx context.Context, input model.MoveInput) (*model.MovePayload, error) {
	// currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	// if err != nil {
	// 	return nil, errUnauthenticated
	// }

	// domainRoomID, err := decodeID(roomPrefix, input.RoomID)
	// if err != nil {
	// 	return nil, err
	// }

	// movedUser := &model.MovePayload{
	// 	UserID: currentUser.ID,
	// 	X:      input.X,
	// 	Y:      input.Y,
	// }

	// if err := r.pubsubRepo.PubMovedUser(ctx, movedUser, domainRoomID); err != nil {
	// 	return nil, err
	// }

	// // userの位置の更新は最悪失敗しても良いので非同期で行う
	// go func() {
	// 	ctx2 := context.Background()
	// 	domainUserID, err := decodeID(userPrefix, currentUser.ID)
	// 	if err != nil {
	// 		log.Println("failed to decode user id, err:", err)
	// 		return
	// 	}

	// 	u, err := r.roomRepo.GetUserByID(ctx2, domainUserID)
	// 	if err != nil {
	// 		log.Println("failed to get user, err:", err)
	// 		return
	// 	}

	// 	u.X = input.X
	// 	u.Y = input.Y

	// 	if err := r.roomRepo.UpdateUser(ctx, u); err != nil {
	// 		log.Println("failed to update user position, err:", err)
	// 		return
	// 	}
	// }()

	// return movedUser, nil

	return nil, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.Me, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}

	me := &model.Me{
		ID:        currentUser.ID,
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
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) ActedRoomUserEvent(ctx context.Context, roomID string) (<-chan model.RoomUserEvent, error) {
	panic(fmt.Errorf("not implemented"))
}
