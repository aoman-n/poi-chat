package repository

import (
	"context"
	"errors"

	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/presentation/graphql/subscriber"
	"github.com/laster18/poi/api/src/util/aerrors"
	"gorm.io/gorm"
)

const OnlineUserIndexKey = "onlineUserIndex"

type User struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewUser(db *gorm.DB, redisClient *redis.Client) *User {
	return &User{db, redisClient}
}

var _ user.Repository = (*User)(nil)

func (r *User) Save(ctx context.Context, u *user.User) error {
	if err := r.db.Create(u).Error; err != nil {
		return aerrors.Wrap(err)
	}

	return nil
}

func (r *User) Get(ctx context.Context, id int) (*user.User, error) {
	var u user.User
	if err := r.db.First(&u, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, aerrors.Wrap(err).SetCode(aerrors.CodeNotFound).Message("not found user")
		}

		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return &u, nil
}

func (r *User) GetByUID(ctx context.Context, uid string) (*user.User, error) {
	var u user.User
	if err := r.db.Where("uid = ?", uid).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, aerrors.Wrap(err).SetCode(aerrors.CodeNotFound).Message("not found user")
		}

		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return &u, nil
}

func (r *User) Online(ctx context.Context, u *user.User) error {
	key := subscriber.MakeOnlineUserKey(u.UID)

	if err := r.redisClient.Set(ctx, key, nil, expireTimeSecond).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}
	if err := r.redisClient.SAdd(ctx, OnlineUserIndexKey, key).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	return nil
}

func (r *User) Offline(ctx context.Context, u *user.User) error {
	key := subscriber.MakeOnlineUserKey(u.UID)

	if err := r.redisClient.Del(ctx, key).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}
	if err := r.redisClient.SRem(ctx, OnlineUserIndexKey, key).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	return nil
}

func (r *User) GetOnlineUsers(ctx context.Context) ([]*user.User, error) {
	keys, err := r.redisClient.SMembers(ctx, OnlineUserIndexKey).Result()
	if err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	if len(keys) <= 0 {
		return []*user.User{}, nil
	}

	userUIDs := make([]string, len(keys))
	for i, key := range keys {
		uID, err := subscriber.DestructOnlineUserKey(key)
		if err != nil {
			return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
		}
		userUIDs[i] = uID
	}

	var users []*user.User
	if err := r.db.Where("uid IN ?", userUIDs).Find(&users).Error; err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return users, nil
}
