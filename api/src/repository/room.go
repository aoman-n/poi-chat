package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/laster18/poi/api/src/domain"
	"gorm.io/gorm"
)

type RoomRepo struct {
	db *gorm.DB
}

var _ domain.IRoomRepo = (*RoomRepo)(nil)

func NewRoomRepo(db *gorm.DB) *RoomRepo {
	return &RoomRepo{db}
}

func (r *RoomRepo) GetByID(ctx context.Context, id int) (*domain.Room, error) {
	var room domain.Room
	if err := r.db.First(&room, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NewNotFoundErr(fmt.Sprintf("not found room, id: %d", id))
		}

		return nil, err
	}

	return &room, nil
}

func (r *RoomRepo) List(ctx context.Context, req *domain.RoomListReq) (*domain.RoomListResp, error) {
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

	var rooms []*domain.Room
	if err := db.Find(&rooms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &domain.RoomListResp{
				List:    []*domain.Room{},
				HasNext: false,
			}, nil
		}

		return nil, err
	}

	if len(rooms) >= req.Limit {
		return &domain.RoomListResp{
			List:    rooms[:req.Limit],
			HasNext: true,
		}, nil
	}

	return &domain.RoomListResp{
		List:    rooms,
		HasNext: false,
	}, nil
}

func (r *RoomRepo) ListAll(ctx context.Context) ([]*domain.Room, error) {
	var rooms []*domain.Room
	if err := r.db.Find(&rooms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*domain.Room{}, nil
		}

		return nil, err
	}

	return rooms, nil
}

func (r *RoomRepo) Count(ctx context.Context) (int, error) {
	var count int64
	if err := r.db.Model(&domain.Room{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *RoomRepo) Create(ctx context.Context, room *domain.Room) error {
	return r.db.Create(room).Error
}

func (r *RoomRepo) Join(ctx context.Context, u *domain.JoinedUser) error {
	return r.db.Create(u).Error
}

func (r *RoomRepo) GetUsers(ctx context.Context, roomID int) ([]*domain.JoinedUser, error) {
	var users []*domain.JoinedUser
	if err := r.db.Where("room_id = ?", roomID).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*domain.JoinedUser{}, nil
		}

		return nil, err
	}

	return users, nil
}

func (r *RoomRepo) Exit(ctx context.Context, u *domain.JoinedUser) error {
	if u.ID == 0 {
		return errZeroID
	}

	return r.db.Delete(u).Error
}

func (r *RoomRepo) UpdateUser(ctx context.Context, u *domain.JoinedUser) error {
	if u.ID == 0 {
		return errZeroID
	}

	return r.db.Save(u).Error
}

func (r *RoomRepo) GetUserByID(ctx context.Context, id int) (*domain.JoinedUser, error) {
	var u domain.JoinedUser
	if err := r.db.First(&u, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NewNotFoundErr(fmt.Sprintf("not found joined user, id: %d", id))
		}

		return nil, err
	}

	return &u, nil
}
