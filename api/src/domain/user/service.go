package user

import (
	"context"
	"errors"

	"github.com/laster18/poi/api/src/util/aerrors"
)

type Service interface {
	FindOrCreate(ctx context.Context, u *User) (*User, error)
	ExistsStatus(ctx context.Context, uid string) (bool, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) FindOrCreate(ctx context.Context, u *User) (*User, error) {
	user, err := s.repo.GetByUID(ctx, u.UID)
	if err != nil {
		var aerr *aerrors.ErrApp
		if errors.As(err, &aerr) {
			if aerr.Code() == aerrors.CodeNotFound {
				if err := s.repo.Save(ctx, u); err != nil {
					return nil, aerrors.Wrap(err)
				}

				return u, nil
			}
		}

		return nil, aerrors.Wrap(err)
	}

	return user, nil
}

func (s *service) ExistsStatus(ctx context.Context, uid string) (bool, error) {
	_, err := s.repo.GetStatus(ctx, uid)
	if err != nil {
		var errApp *aerrors.ErrApp
		if errors.As(err, &errApp) {
			if errApp.Code() == aerrors.CodeNotFound {
				return false, nil
			}
		}

		return false, aerrors.Wrap(err)
	}

	return true, nil
}
