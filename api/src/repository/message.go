package repository

import (
	"context"
	"errors"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/util/aerrors"
	"gorm.io/gorm"
)

type MessageRepo struct {
	db *gorm.DB
}

var _ domain.IMessageRepo = (*MessageRepo)(nil)

func NewMessageRepo(db *gorm.DB) *MessageRepo {
	return &MessageRepo{db}
}

func (r *MessageRepo) List(ctx context.Context, req *domain.MessageListReq) (*domain.MessageListResp, error) {
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

	db = db.Where("room_id = ?", req.RoomID).
		Order("created_at desc, id desc").
		Limit(req.Limit + 1)

	var messages []*domain.Message
	if err := db.Find(&messages).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &domain.MessageListResp{
				List:            []*domain.Message{},
				HasPreviousPage: false,
			}, nil
		}

		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	for i := 0; i < len(messages)/2; i++ {
		messages[i], messages[len(messages)-i-1] = messages[len(messages)-i-1], messages[i]
	}

	if len(messages) > req.Limit {
		return &domain.MessageListResp{
			List:            messages[1:],
			HasPreviousPage: true,
		}, nil
	}

	return &domain.MessageListResp{
		List:            messages,
		HasPreviousPage: false,
	}, nil
}

func (r *MessageRepo) Create(ctx context.Context, message *domain.Message) error {
	if err := r.db.Create(message).Error; err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return nil
}

func (r *MessageRepo) Count(ctx context.Context, roomID int) (int, error) {
	var count int64
	if err := r.db.Model(&domain.Message{}).
		Where("room_id = ?", roomID).
		Count(&count).Error; err != nil {
		return 0, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return int(count), nil
}

type messageCount struct {
	RoomID int
	Count  int
}

func (r *MessageRepo) CountByRoomIDs(ctx context.Context, roomIDs []int) ([]int, error) {
	var messageCounts []messageCount

	if err := r.db.Table("messages").
		Select("room_id, count(room_id) as count").
		Where("room_id IN ?", roomIDs).
		Group("room_id").
		Find(&messageCounts).
		Error; err != nil {
		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	// 渡されたroomID順に詰めて返す
	var counts []int
	for _, roomID := range roomIDs {
		for _, c := range messageCounts {
			if roomID == c.RoomID {
				counts = append(counts, c.Count)
			}
		}
	}

	return counts, nil
}
