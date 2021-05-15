package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/subscriber"
	"gorm.io/gorm"
)

const (
	ExpireTimeSecond = 60 * time.Second
)

type RoomUserRepo struct {
	tx          *gorm.DB
	redisClient *redis.Client
}

var _ domain.IRoomUserRepo = (*RoomUserRepo)(nil)

func NewRoomUserRepo(tx *gorm.DB, redisClient *redis.Client) *RoomUserRepo {
	return &RoomUserRepo{tx, redisClient}
}

func (r *RoomUserRepo) Insert(ctx context.Context, ru *domain.RoomUser) error {
	ruJSON, err := json.Marshal(ru)
	if err != nil {
		return err
	}

	roomUserKey := subscriber.MakeRoomUserKey(ru.RoomID, ru.UID)

	fmt.Println("[set] roomUserKey:", roomUserKey)

	if err := r.redisClient.Set(
		ctx,
		roomUserKey,
		ruJSON,
		ExpireTimeSecond,
	).Err(); err != nil {
		return err
	}

	// ルーム内ユーザーのkey一覧を一括で取得するためindexとしてsaddで保存する
	// @see: https://stackoverflow.com/questions/32474699/redis-find-keys-matching-a-pattern
	if err := r.redisClient.SAdd(
		ctx,
		r.makeRoomUserIndexKey(ru.RoomID),
		roomUserKey,
	).Err(); err != nil {
		return err
	}

	return nil
}

func (r *RoomUserRepo) Delete(ctx context.Context, ru *domain.RoomUser) error {
	roomUserKey := subscriber.MakeRoomUserKey(ru.RoomID, ru.UID)

	if err := r.redisClient.Del(ctx, roomUserKey).Err(); err != nil {
		return err
	}
	if err := r.redisClient.SRem(
		ctx,
		roomUserKey,
		r.makeRoomUserIndexKey(ru.RoomID),
		roomUserKey,
	).Err(); err != nil {
		return err
	}

	return nil
}

func (r *RoomUserRepo) Get(ctx context.Context, roomID int, uID string) (*domain.RoomUser, error) {
	c := r.redisClient.Get(ctx, subscriber.MakeRoomUserKey(roomID, uID))
	ruJSON, err := c.Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, err
	}

	var roomUser domain.RoomUser
	if err := json.Unmarshal([]byte(ruJSON), &roomUser); err != nil {
		return nil, err
	}

	return &roomUser, nil
}

const RoomUserIndexKey = "roomUserIndex"

func (r *RoomUserRepo) makeRoomUserIndexKey(roomID int) string {
	return fmt.Sprintf("%s:%d", RoomUserIndexKey, roomID)
}
