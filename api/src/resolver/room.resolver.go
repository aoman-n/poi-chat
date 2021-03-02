package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/graph/model"
)

func (r *messageResolver) ID(ctx context.Context, obj *model.Message) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateRoom(ctx context.Context, input *model.CreateRoomInput) (*model.RoomDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*model.RoomSummary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Room(ctx context.Context, id string) ([]*model.RoomDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *roomDetailResolver) ID(ctx context.Context, obj *model.RoomDetail) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *roomSummaryResolver) ID(ctx context.Context, obj *model.RoomSummary) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) SubMessage(ctx context.Context, roomID string) (<-chan *model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

// RoomDetail returns generated.RoomDetailResolver implementation.
func (r *Resolver) RoomDetail() generated.RoomDetailResolver { return &roomDetailResolver{r} }

// RoomSummary returns generated.RoomSummaryResolver implementation.
func (r *Resolver) RoomSummary() generated.RoomSummaryResolver { return &roomSummaryResolver{r} }

type messageResolver struct{ *Resolver }
type roomDetailResolver struct{ *Resolver }
type roomSummaryResolver struct{ *Resolver }
