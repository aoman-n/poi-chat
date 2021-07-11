package repository

import (
	"context"
	"errors"

	"github.com/laster18/poi/api/src/domain/message"
	"github.com/laster18/poi/api/src/util/aerrors"
	"gorm.io/gorm"
)

type Message struct {
	db *gorm.DB
}

func NewMessage(db *gorm.DB) *Message {
	return &Message{db}
}

var _ message.Repository = (*Message)(nil)

func (r *Message) List(ctx context.Context, req *message.ListReq) (*message.ListResp, error) {
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

	var messages []*message.Message
	if err := db.Find(&messages).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &message.ListResp{
				List:            []*message.Message{},
				HasPreviousPage: false,
			}, nil
		}

		return nil, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	for i := 0; i < len(messages)/2; i++ {
		messages[i], messages[len(messages)-i-1] = messages[len(messages)-i-1], messages[i]
	}

	if len(messages) > req.Limit {
		return &message.ListResp{
			List:            messages[1:],
			HasPreviousPage: true,
		}, nil
	}

	return &message.ListResp{
		List:            messages,
		HasPreviousPage: false,
	}, nil
}

func (r *Message) Create(ctx context.Context, message *message.Message) error {
	if err := r.db.Create(message).Error; err != nil {
		return aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return nil
}

func (r *Message) Count(ctx context.Context, roomID int) (int, error) {
	var count int64
	if err := r.db.Model(&message.Message{}).
		Where("room_id = ?", roomID).
		Count(&count).Error; err != nil {
		return 0, aerrors.Wrap(err).SetCode(aerrors.CodeDatabase)
	}

	return int(count), nil
}
