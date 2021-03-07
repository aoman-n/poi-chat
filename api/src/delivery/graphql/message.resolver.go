package graphql

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

func (r *roomResolver) Messages(ctx context.Context, obj *model.Room, first *int) ([]*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

type messageResolver struct{ *Resolver }