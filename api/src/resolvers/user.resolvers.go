package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/laster18/poi/api/graph/model"
)

func (r *mutationResolver) AddUser(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context) (*model.User, error) {
	return &model.User{
		ID: "112233",
		Name: "sam",
	}, nil
}
