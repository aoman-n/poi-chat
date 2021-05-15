package repository

import (
	"context"
	"encoding/json"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/subscriber"
)

type GlobalUserRepo struct {
	redisClient *redis.Client
}

var _ domain.GlobalUserRepo = (*GlobalUserRepo)(nil)

func NewGlobalUserRepo(redis *redis.Client) *GlobalUserRepo {
	return &GlobalUserRepo{redis}
}

func (r *GlobalUserRepo) Insert(ctx context.Context, u *domain.GlobalUser) error {
	uJSON, err := json.Marshal(u)
	if err != nil {
		return err
	}

	key := subscriber.MakeGlobalUserKey(u.UID)

	if err := r.redisClient.Set(ctx, key, uJSON, expireTimeSecond).Err(); err != nil {
		return err
	}
	if err := r.redisClient.SAdd(ctx, globalUserIndexKey, key).Err(); err != nil {
		return err
	}

	return nil
}

func (r *GlobalUserRepo) Delete(ctx context.Context, u *domain.GlobalUser) error {
	key := subscriber.MakeGlobalUserKey(u.UID)

	if err := r.redisClient.Del(ctx, key).Err(); err != nil {
		return err
	}
	if err := r.redisClient.SRem(ctx, globalUserIndexKey, key).Err(); err != nil {
		return err
	}

	return nil
}

func (r *GlobalUserRepo) Get(ctx context.Context, uID string) (*domain.GlobalUser, error) {
	key := subscriber.MakeGlobalUserKey(uID)
	uJSON, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var globalUser domain.GlobalUser
	if err := json.Unmarshal([]byte(uJSON), &globalUser); err != nil {
		return nil, err
	}

	return &globalUser, nil
}

const globalUserIndexKey = "globalUserIndex"
