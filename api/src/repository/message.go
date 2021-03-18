package repository

import (
	"context"
	"errors"

	"github.com/laster18/poi/api/src/domain"
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

		return nil, err
	}

	if len(messages) >= req.Limit {
		return &domain.MessageListResp{
			List:            messages[:req.Limit],
			HasPreviousPage: true,
		}, nil
	}

	return &domain.MessageListResp{
		List:            messages,
		HasPreviousPage: false,
	}, nil
}

func (r *MessageRepo) Create(ctx context.Context, message *domain.Message) error {
	return r.db.Create(message).Error
}

func (r *MessageRepo) Count(ctx context.Context, roomID int) (int, error) {
	var count int64
	if err := r.db.Model(&domain.Message{}).
		Where("room_id = ?", roomID).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}