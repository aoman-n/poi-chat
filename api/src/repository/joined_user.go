package repository

import (
	"context"
	"errors"

	"github.com/laster18/poi/api/src/domain"
	"gorm.io/gorm"
)

type JoinedUserRepo struct {
	db *gorm.DB
}

var _ domain.IJoinedUserRepo = (*JoinedUserRepo)(nil)

func NewJoinedUserRepo(db *gorm.DB) *JoinedUserRepo {
	return &JoinedUserRepo{db: db}
}

func (r *JoinedUserRepo) Create(ctx context.Context, u *domain.JoinedUser) error {
	return r.db.Create(u).Error
}

func (r *JoinedUserRepo) List(ctx context.Context, roomID int) ([]*domain.JoinedUser, error) {
	var users []*domain.JoinedUser
	if err := r.db.Where("room_id = ?", roomID).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*domain.JoinedUser{}, nil
		}

		return nil, err
	}

	return users, nil
}

func (r *JoinedUserRepo) Delete(ctx context.Context, u *domain.JoinedUser) error {
	if u.ID == 0 {
		return errZeroID
	}

	return r.db.Delete(u).Error
}
