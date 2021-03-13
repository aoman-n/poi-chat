package repository

import (
	"context"

	"github.com/laster18/poi/api/src/domain"
	"gorm.io/gorm"
)

type RoomRepo struct {
	db *gorm.DB
}

var _ domain.IRoomRepo = (*RoomRepo)(nil)

func NewRoomRepo(db *gorm.DB) *RoomRepo {
	return &RoomRepo{db}
}

func (r *RoomRepo) GetByID(ctx context.Context, id int) (*domain.Room, error) {
	return nil, nil
}

func (r *RoomRepo) List(ctx context.Context, req *domain.RoomListReq) (*domain.RoomListResp, error) {
	if req.Limit == 0 {
		req.Limit = 10
	}

	db := r.db
	if req.LastKnownID != 0 && req.LastKnownUnix != 0 {
		db = db.Where(
			"(UNIX_TIMESTAMP(created_at) < ?) OR (UNIX_TIMESTAMP(created_at) = ? AND id < ?)",
			req.LastKnownUnix,
			req.LastKnownUnix,
			req.LastKnownID,
		)
	}

	db = db.Order("created_at desc, id desc").Limit(req.Limit + 1)

	var rooms []*domain.Room
	if err := db.Find(&rooms).Error; err != nil {
		return nil, err
	}

	if len(rooms) >= req.Limit {
		return &domain.RoomListResp{
			List:    rooms[:req.Limit],
			HasNext: true,
		}, nil
	}

	return &domain.RoomListResp{
		List:    rooms,
		HasNext: false,
	}, nil
}

func (r *RoomRepo) Count(ctx context.Context) (int, error) {
	var count int64
	if err := r.db.Model(&domain.Room{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *RoomRepo) Create(ctx context.Context, room *domain.Room) error {
	return r.db.Create(room).Error
}
