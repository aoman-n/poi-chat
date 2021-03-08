package repository

import (
	"context"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infrastructure"
)

type RoomRepo struct {
	db *infrastructure.Db
}

var _ domain.IRoomRepo = (*RoomRepo)(nil)

func NewRoomRepo(db *infrastructure.Db) *RoomRepo {
	return &RoomRepo{db}
}

func (r *RoomRepo) FindByID(ctx context.Context, id int32) (*domain.Room, error) {
	return nil, nil
}
