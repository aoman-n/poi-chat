package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/presentation/graphql/subscriber"
	"github.com/laster18/poi/api/src/util/aerrors"
	"gorm.io/gorm"
)

const GlobalUserIndexKey = "onlineUserIndex"

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

func (r *User) GetByIDs(ctx context.Context, ids []int) ([]*user.User, error) {
	var u []*user.User
	if err := r.db.Where("id in ?", ids).Find(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, aerrors.Wrap(err).SetCode(aerrors.CodeNotFound).Message("not found user")
		}

		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return u, nil
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

func (r *User) GetByUIDs(ctx context.Context, uids []string) ([]*user.User, error) {
	var u []*user.User
	if err := r.db.Where("uid in ?", uids).Find(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, aerrors.Wrap(err).SetCode(aerrors.CodeNotFound).Message("not found user")
		}

		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return u, nil
}

func (r *User) SaveStatus(ctx context.Context, id int, status *user.Status) error {
	statusBytes, err := json.Marshal(status)
	if err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeInternal)
	}

	key := subscriber.MakeOnlineUserKey(id)

	if err := r.redisClient.Set(ctx, key, statusBytes, expireTimeSecond).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}
	if err := r.redisClient.SAdd(ctx, GlobalUserIndexKey, key).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	return nil
}

func (r *User) DeleteStatus(ctx context.Context, id int) error {
	key := subscriber.MakeOnlineUserKey(id)

	if err := r.redisClient.Del(ctx, key).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}
	if err := r.redisClient.SRem(ctx, GlobalUserIndexKey, key).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	return nil
}

func (r *User) GetStatus(ctx context.Context, id int) (*user.Status, error) {
	key := subscriber.MakeOnlineUserKey(id)

	j, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, aerrors.Wrap(err).SetCode(aerrors.CodeNotFound).Message("not found user status")
		}

		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	var status user.Status
	if err := json.Unmarshal([]byte(j), &status); err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeInternal)
	}

	return &status, nil
}

func (r *User) GetStatuses(ctx context.Context, ids []int) ([]*user.Status, error) {
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = subscriber.MakeOnlineUserKey(id)
	}

	statusJSONs, err := r.redisClient.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	statuses := []*user.Status{}
	for _, statusJSON := range statusJSONs {
		if statusJSON == nil {
			continue
		}

		str, ok := statusJSON.(string)
		if !ok {
			continue
		}

		var s user.Status
		if err := json.Unmarshal([]byte(str), &s); err != nil {
			continue
		}

		statuses = append(statuses, &s)
	}

	return statuses, nil
}

func (r *User) GetOnlineUsers(ctx context.Context) ([]*user.User, error) {
	keys, err := r.redisClient.SMembers(ctx, GlobalUserIndexKey).Result()
	if err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	if len(keys) <= 0 {
		return []*user.User{}, nil
	}

	userIDs := make([]int, len(keys))
	for i, key := range keys {
		uID, err := subscriber.DestructOnlineUserKey(key)
		if err != nil {
			return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
		}
		userIDs[i] = uID
	}

	var users []*user.User
	if err := r.db.Where("id IN ?", userIDs).Find(&users).Error; err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return users, nil
}
