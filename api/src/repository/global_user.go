package repository

import (
	"context"
	"encoding/json"
	"log"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/presentation/graphql/subscriber"
	"github.com/laster18/poi/api/src/util/aerrors"
)

const globalUserIndexKey = "globalUserIndex"

type GlobalUserRepo struct {
	redisClient *redis.Client
}

var _ domain.GlobalUserRepo = (*GlobalUserRepo)(nil)

func NewGlobalUserRepo(redis *redis.Client) *GlobalUserRepo {
	return &GlobalUserRepo{redis}
}

func (r *GlobalUserRepo) Save(ctx context.Context, u *domain.GlobalUser) error {
	uJSON, err := json.Marshal(u)
	if err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeInternal)
	}

	key := subscriber.MakeGlobalUserKey(u.UID)

	if err := r.redisClient.Set(ctx, key, uJSON, expireTimeSecond).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}
	if err := r.redisClient.SAdd(ctx, globalUserIndexKey, key).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	return nil
}

func (r *GlobalUserRepo) Delete(ctx context.Context, uID string) error {
	key := subscriber.MakeGlobalUserKey(uID)

	if err := r.redisClient.Del(ctx, key).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}
	if err := r.redisClient.SRem(ctx, globalUserIndexKey, key).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	return nil
}

func (r *GlobalUserRepo) Get(ctx context.Context, uID string) (*domain.GlobalUser, error) {
	key := subscriber.MakeGlobalUserKey(uID)
	uJSON, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	var globalUser domain.GlobalUser
	if err := json.Unmarshal([]byte(uJSON), &globalUser); err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeInternal)
	}

	return &globalUser, nil
}

func (r *GlobalUserRepo) GetAll(ctx context.Context) ([]*domain.GlobalUser, error) {
	keys, err := r.redisClient.SMembers(ctx, globalUserIndexKey).Result()
	if err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	if len(keys) <= 0 {
		return []*domain.GlobalUser{}, nil
	}

	userJSONs, err := r.redisClient.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	// TODO: ここの処理並列でやりたい
	onlineUsers := []*domain.GlobalUser{}
	for i, u := range userJSONs {
		if u == nil {
			// dataが存在しないindexは削除しておく
			if err := r.redisClient.SRem(ctx, globalUserIndexKey, keys[i]).Err(); err != nil {
				log.Printf("failed to delete globalUser index which never existed from Sets data")
			}
			continue
		}

		jsonStr, ok := u.(string)
		if !ok {
			log.Printf(
				"globalUser data is invalid type on Redis, type: %T, data: %v\n",
				u,
				u,
			)
			continue
		}

		var ou domain.GlobalUser
		if err := json.Unmarshal([]byte(jsonStr), &ou); err != nil {
			log.Printf(
				"globalUser data is invalid format on Redis, data: %+v\n",
				u,
			)
			continue
		}

		onlineUsers = append(onlineUsers, &ou)
	}

	return onlineUsers, nil
}
