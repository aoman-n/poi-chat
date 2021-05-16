package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/subscriber"
	"gorm.io/gorm"
)

const RoomUserIndexKey = "roomUserIndex"

func (r *RoomUserRepo) makeRoomUserIndexKey(roomID int) string {
	return fmt.Sprintf("%s:%d", RoomUserIndexKey, roomID)
}

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
		expireTimeSecond,
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

func (r *RoomUserRepo) GetByRoomID(ctx context.Context, id int) ([]*domain.RoomUser, error) {
	indexCh := r.makeRoomUserIndexKey(id)
	keys, err := r.redisClient.SMembers(ctx, indexCh).Result()

	if err != nil {
		return nil, err
	}
	if len(keys) <= 0 {
		return []*domain.RoomUser{}, nil
	}

	userJSONs, err := r.redisClient.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	roomUsers := []*domain.RoomUser{}
	for i, u := range userJSONs {
		if u == nil {
			if err := r.redisClient.SRem(ctx, indexCh, keys[i]).Err(); err != nil {
				log.Printf("failed to delete roomUser index which never existed from Sets data")
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

		var ru domain.RoomUser
		if err := json.Unmarshal([]byte(jsonStr), &ru); err != nil {
			log.Printf(
				"roomUser data is invalid format on Redis, data: %+v\n",
				u,
			)
			continue
		}

		roomUsers = append(roomUsers, &ru)
	}

	return roomUsers, nil
}

func (r *RoomUserRepo) Counts(ctx context.Context, roomIDs []int) ([]int, error) {
	roomUserIndexKeys := make([]string, len(roomIDs))
	for i, id := range roomIDs {
		roomUserIndexKeys[i] = r.makeRoomUserIndexKey(id)
	}

	// roomUserKeys:
	// ["roomUser:<roomId>:<userId>", ...]
	// ["roomUser:1:1111", ...]
	roomUserKeys, err := r.redisClient.SUnion(ctx, roomUserIndexKeys...).Result()
	if err != nil {
		return nil, err
	}

	countMap := map[int]int{}
	for _, roomID := range roomIDs {
		countMap[roomID] = 0
	}

	for _, key := range roomUserKeys {
		// Splitで分解するようにしたほうがよい？
		roomID, _, _ := subscriber.DestructRoomUserKey(key)
		countMap[roomID]++
	}

	// TODO: もっと良い書き方をしたい
	counts := make([]int, len(roomIDs))
	i := 0
	for _, count := range countMap {
		counts[i] = count
		i++
	}

	return counts, nil
}
