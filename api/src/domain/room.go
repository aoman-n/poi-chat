package domain

import "context"

type Room struct {
	ID   int32
	Name string
}

type IRoomRepository interface {
	FindByID(ctx context.Context, id int32) (*Room, error)
}
