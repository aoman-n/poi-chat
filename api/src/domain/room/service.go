package room

import (
	"context"
	"errors"

	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/util/aerrors"
)

type Service interface {
	ExistsRoom(ctx context.Context, roomName string) (bool, error)
	FindOrNewUserStatus(ctx context.Context, u *user.User, roomID int) (*UserStatus, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ExistsRoom(ctx context.Context, roomName string) (bool, error) {
	_, err := s.repo.GetByName(ctx, roomName)
	if err != nil {
		var aerr *aerrors.ErrApp
		if errors.As(err, &aerr) {
			if aerr.Code() == aerrors.CodeNotFound {
				return false, nil
			}

			return false, aerrors.Wrap(err)
		}
	}

	return true, nil
}

func (s *service) FindOrNewUserStatus(ctx context.Context, u *user.User, roomID int) (*UserStatus, error) {
	us, err := s.repo.GetUserStatus(ctx, roomID, u.ID)
	if err != nil {
		var errApp *aerrors.ErrApp
		if errors.As(err, &errApp) {
			if errApp.Code() == aerrors.CodeNotFound {
				return NewUserStatus(u, roomID), nil
			}
		}

		return nil, aerrors.Wrap(err)
	}

	return us, nil
}
