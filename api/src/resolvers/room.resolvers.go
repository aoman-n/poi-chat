package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/laster18/poi/api/graph/model"
)

func (r *mutationResolver) CreateRoom(ctx context.Context, input *model.CreateRoomInput) (*model.RoomDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*model.RoomSummry, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Room(ctx context.Context, id string) ([]*model.RoomDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) SubMessage(ctx context.Context, roomID string) (<-chan *model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}
