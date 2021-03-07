package repository

import (
	"context"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infrastructure"
)

type RoomRepository struct {
	db *infrastructure.Db
}

var _ domain.IRoomRepository = (*RoomRepository)(nil)

func NewRoomRepository(db *infrastructure.Db) *RoomRepository {
	return &RoomRepository{db}
}

func (r *RoomRepository) FindByID(ctx context.Context, id int32) (*domain.Room, error) {
	return nil, nil
}
