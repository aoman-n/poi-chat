package repository

import (
	"context"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infra/redis"
)

type GlobalUserRepo struct {
	redisClient *redis.Client
}

var _ domain.GlobalUserRepo = (*GlobalUserRepo)(nil)

func NewGlobalUserRepo(redis *redis.Client) *GlobalUserRepo {
	return &GlobalUserRepo{redis}
}

func (r *GlobalUserRepo) Insert(ctx context.Context, u *domain.GlobalUser) error {
	return nil
}

func (r *GlobalUserRepo) Delete(ctx context.Context, u *domain.GlobalUser) error {
	return nil
}

func (r *GlobalUserRepo) Get(ctx context.Context, uID string) (*domain.GlobalUser, error) {
	return nil, nil
}
