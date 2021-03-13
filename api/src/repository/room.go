package repository

import (
	"context"
	"fmt"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infrastructure/db"
)

type RoomRepo struct {
	db *db.Db
}

var _ domain.IRoomRepo = (*RoomRepo)(nil)

func NewRoomRepo(db *db.Db) *RoomRepo {
	return &RoomRepo{db}
}

func (r *RoomRepo) GetByID(ctx context.Context, id int32) (*domain.Room, error) {
	return nil, nil
}

func (r *RoomRepo) List(ctx context.Context, req *domain.RoomListReq) (*domain.RoomListResp, error) {
	if req.Limit == 0 {
		req.Limit = 10
	}

	fmt.Println("list start")

	var rooms []*domain.Room
	if err := r.db.
		Order("created_at").
		Limit(req.Limit).
		Find(&rooms).Error; err != nil {
		return nil, err
	}

	return &domain.RoomListResp{
		List: rooms,
	}, nil
}

func (r *RoomRepo) Create(ctx context.Context, room *domain.Room) error {
	return r.db.Create(room).Error
}
