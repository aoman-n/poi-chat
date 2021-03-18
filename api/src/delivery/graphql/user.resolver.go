package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/middleware"
)

func (r *mutationResolver) Move(ctx context.Context, input model.MoveInput) (*model.MovedUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *roomDetailResolver) JoinedUsers(ctx context.Context, obj *model.RoomDetail) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) SubMovedUser(ctx context.Context, roomID string) (<-chan *model.MovedUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) SubExitedUser(ctx context.Context, roomID string) (<-chan *model.ExitedUser, error) {
	panic(fmt.Errorf("not implemented"))
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

	// TODO: check the room exists

	joinedUser := &domain.JoinedUser{
		RoomID:      domainRoomID,
		AvatarURL:   currentUser.AvatarURL,
		DisplayName: currentUser.Name,
		UserID:      currentUser.ID,
		// set default position
		X: 100,
		Y: 100,
	}

	if err := r.joinedUserRepo.Create(ctx, joinedUser); err != nil {
		log.Println("failed to create joinedUser, err:", err)
		return nil, errUnexpected
	}

	go func() {
		<-ctx.Done()
		if err := r.joinedUserRepo.Delete(ctx, joinedUser); err != nil {
			// TODO: retry process
			log.Println("failed to delete joinedUser err:", err)
		}
	}()

	ch := make(chan *model.User)

	go func() {
		time.Sleep(1 * time.Second)

		ch <- &model.User{
			ID:          encodeID(roomPrefix, joinedUser.ID),
			DisplayName: joinedUser.DisplayName,
			AvatarURL:   joinedUser.AvatarURL,
			X:           joinedUser.X,
			Y:           joinedUser.Y,
		}
	}()

	return ch, nil
}
