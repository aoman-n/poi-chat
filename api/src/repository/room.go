package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/presentation/graphql/subscriber"
	"github.com/laster18/poi/api/src/util/aerrors"
	"gorm.io/gorm"
)

type Room struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRoom(db *gorm.DB, redisClient *redis.Client) *Room {
	return &Room{db, redisClient}
}

var _ room.Repository = (*Room)(nil)

const RoomUserStatusIndexKey = "roomUserStatusIndex"

// makeRoomUserStatusIndexKey roomUserStatusIndex:<roomId>
func (r *Room) makeRoomUserStatusIndexKey(roomID int) string {
	return fmt.Sprintf("%s:%d", RoomUserStatusIndexKey, roomID)
}

func (r *Room) GetByID(ctx context.Context, id int) (*room.Room, error) {
	var room room.Room
	if err := r.db.First(&room, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := fmt.Sprintf("not found room, id: %d", id)
			return nil, aerrors.New(msg).SetCode(aerrors.CodeNotFound).Message("not found room")
		}

		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return &room, nil
}

func (r *Room) GetByIDs(ctx context.Context, ids []int) ([]*room.Room, error) {
	var rooms []*room.Room
	if err := r.db.Where("id IN ?", ids).Find(&rooms).Error; err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return rooms, nil
}

func (r *Room) GetByName(ctx context.Context, name string) (*room.Room, error) {
	var room room.Room
	if err := r.db.Where("name = ?", name).First(&room).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := fmt.Sprintf("not found room, name: %s", name)
			return nil, aerrors.New(msg).SetCode(aerrors.CodeNotFound).Message("not found room")
		}

		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return &room, nil
}

func (r *Room) GetAll(ctx context.Context) ([]*room.Room, error) {
	panic("not implemented") // TODO: Implement
}

func (r *Room) List(ctx context.Context, req *room.ListReq) (*room.ListResp, error) {
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

	var rooms []*room.Room
	if err := db.Find(&rooms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &room.ListResp{
				List:    []*room.Room{},
				HasNext: false,
			}, nil
		}

		return nil, err
	}

	if len(rooms) >= req.Limit {
		return &room.ListResp{
			List:    rooms[:req.Limit],
			HasNext: true,
		}, nil
	}

	return &room.ListResp{
		List:    rooms,
		HasNext: false,
	}, nil
}

func (r *Room) Create(ctx context.Context, room *room.Room) error {
	if err := r.db.Create(room).Error; err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return nil
}

func (r *Room) Count(ctx context.Context) (int, error) {
	var count int64
	if err := r.db.Model(&room.Room{}).Count(&count).Error; err != nil {
		return 0, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return int(count), nil
}

func (r *Room) CountUserByRoomIDs(ctx context.Context, roomIDs []int) ([]int, error) {
	roomUserStatusIndexKeys := make([]string, len(roomIDs))
	for i, id := range roomIDs {
		roomUserStatusIndexKeys[i] = r.makeRoomUserStatusIndexKey(id)
	}

	// roomUserStatusKeys:
	// 型: ["roomUser:<roomId>:<userId>", ...]
	// 例: ["roomUser:1:1111", ...]
	roomUserStatusKeys, err := r.redis.SUnion(ctx, roomUserStatusIndexKeys...).Result()
	if err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	countMap := map[int]int{}
	for _, roomID := range roomIDs {
		countMap[roomID] = 0
	}

	for _, key := range roomUserStatusKeys {
		roomID, _, _ := subscriber.DestructRoomUserStatusKey(key)
		countMap[roomID]++
	}

	ret := make([]int, len(roomIDs))
	for i, roomID := range roomIDs {
		ret[i] = countMap[roomID]
	}

	return ret, nil
}

type messageCount struct {
	RoomID int
	Count  int
}

func (r *Room) CountMessageByRoomIDs(ctx context.Context, roomIDs []int) ([]int, error) {
	var messageCountRet []messageCount

	if err := r.db.Table("messages").
		Select("room_id, count(room_id) as count").
		Where("room_id IN ?", roomIDs).
		Group("room_id").
		Find(&messageCountRet).
		Error; err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	// 渡されたroomID順に詰めて返す
	counts := make([]int, len(roomIDs))
	for i, roomID := range roomIDs {
		for _, c := range messageCountRet {
			if roomID == c.RoomID {
				counts[i] = c.Count
			}
		}
	}

	return counts, nil
}

func (r *Room) GetUsers(ctx context.Context, roomID int) ([]*user.User, error) {
	indexCh := r.makeRoomUserStatusIndexKey(roomID)
	keys, err := r.redis.SMembers(ctx, indexCh).Result()
	if err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}
	if len(keys) <= 0 {
		return []*user.User{}, nil
	}

	statuses, err := r.redis.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	userUIDs := []string{}
	for i, status := range statuses {
		if status == nil {
			if err := r.redis.SRem(ctx, indexCh, keys[i]).Err(); err != nil {
				log.Printf("failed to delete roomUserStatus index which never existed from Sets data")
			}
			continue
		}

		jsonStr, ok := status.(string)
		if !ok {
			log.Printf(
				"roomUserStus data is invalid type on Redis, type: %T, data: %v\n",
				status,
				status,
			)
			continue
		}

		var s room.UserStatus
		if err := json.Unmarshal([]byte(jsonStr), &s); err != nil {
			log.Printf(
				"roomUserStatus data is invalid format on Redis, data: %+v\n",
				status,
			)
			continue
		}

		userUIDs = append(userUIDs, s.UserUID)
	}

	var users []*user.User
	if err := r.db.Where("uid IN ?", userUIDs).Find(&users).Error; err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return users, nil
}

func (r *Room) SaveUserStatus(ctx context.Context, status *room.UserStatus) error {
	statusBytes, err := json.Marshal(status)
	if err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeInternal)
	}

	key := subscriber.MakeRoomUserStatusKey(status.RoomID, status.UserUID)

	if err := r.redis.Set(ctx, key, statusBytes, expireTimeSecond).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}
	if err := r.redis.SAdd(ctx, r.makeRoomUserStatusIndexKey(status.RoomID), key).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	return nil
}

func (r *Room) DeleteUserStatus(ctx context.Context, status *room.UserStatus) error {
	key := subscriber.MakeRoomUserStatusKey(status.RoomID, status.UserUID)

	if err := r.redis.Del(ctx, key).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}
	if err := r.redis.SRem(ctx, r.makeRoomUserStatusIndexKey(status.RoomID), key).Err(); err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	return nil
}

func (r *Room) GetUserStatus(ctx context.Context, roomID int, userUID string) (*room.UserStatus, error) {
	key := subscriber.MakeRoomUserStatusKey(roomID, userUID)

	userStatusStr, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, aerrors.Wrap(err).SetCode(aerrors.CodeNotFound).Message("not found user status")
		}
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	var us room.UserStatus
	if err := json.Unmarshal([]byte(userStatusStr), &us); err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeInternal)
	}

	return &us, nil
}

func (r *Room) GetUserStatuses(ctx context.Context, roomID int) ([]*room.UserStatus, error) {
	indexKey := r.makeRoomUserStatusIndexKey(roomID)
	keys, err := r.redis.SMembers(ctx, indexKey).Result()
	if err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}
	if len(keys) <= 0 {
		return []*room.UserStatus{}, nil
	}

	userStatusJSONs, err := r.redis.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
	}

	userStatuses := make([]*room.UserStatus, len(userStatusJSONs))
	for i, userStatusJSON := range userStatusJSONs {
		if userStatusJSON == nil {
			continue
		}

		jsonStr, ok := userStatusJSON.(string)
		if !ok {
			continue
		}

		var status room.UserStatus
		if err := json.Unmarshal([]byte(jsonStr), &status); err != nil {
			continue
		}

		userStatuses[i] = &status
	}

	return userStatuses, nil
}

// func (r *Room) GetUserStatuses(ctx context.Context, roomID int, userUIDs []string) ([]*room.UserStatus, error) {
// 	keys := make([]string, len(userUIDs))
// 	for i, uid := range userUIDs {
// 		keys[i] = subscriber.MakeRoomUserStatusKey(roomID, uid)
// 	}

// 	userStatusJSONs, err := r.redis.MGet(ctx, keys...).Result()
// 	if err != nil {
// 		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeRedis)
// 	}

// 	userStatuses := make([]*room.UserStatus, len(userStatusJSONs))
// 	for i, userStatusJSON := range userStatusJSONs {
// 		if userStatusJSON == nil {
// 			continue
// 		}

// 		jsonStr, ok := userStatusJSON.(string)
// 		if !ok {
// 			continue
// 		}

// 		var status room.UserStatus
// 		if err := json.Unmarshal([]byte(jsonStr), &status); err != nil {
// 			continue
// 		}

// 		userStatuses[i] = &status
// 	}

// 	return userStatuses, nil
// }
