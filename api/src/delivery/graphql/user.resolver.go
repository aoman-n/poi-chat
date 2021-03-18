package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/laster18/poi/api/graph/model"
)

func (r *roomDetailResolver) JoinedUsers(ctx context.Context, obj *model.RoomDetail) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) SubMovedUser(ctx context.Context, roomID string) (<-chan *model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) SubExitedUser(ctx context.Context, roomID string) (<-chan *model.ExitedUser, error) {
	panic(fmt.Errorf("not implemented"))
}
