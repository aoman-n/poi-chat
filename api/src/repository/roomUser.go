package repository

import (
	"context"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infra/redis"
	"gorm.io/gorm"
)

type RoomUserRepo struct {
	tx          *gorm.DB
	redisClient *redis.Client
}

func NewRoomUserRepo(tx *gorm.DB, redisClient *redis.Client) *RoomUserRepo {
	return &RoomUserRepo{tx, redisClient}
}

func (r *RoomUserRepo) Create(ctx context.Context, u *domain.RoomUser) error {
	return nil
}

func (r *RoomUserRepo) Update(ctx context.Context, u *domain.RoomUser) error {
	return nil
}

func (r *RoomUserRepo) Delete(ctx context.Context, u *domain.RoomUser) error {
	return nil
}

func (r *RoomUserRepo) Get(ctx context.Context, uID int) (*domain.RoomUser, error) {
	return nil, nil
}
