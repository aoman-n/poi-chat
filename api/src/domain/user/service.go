package user

import (
	"context"
	"errors"

	"github.com/laster18/poi/api/src/util/aerrors"
)

type Service interface {
	FindOrCreate(context.Context, *User) (*User, error)
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
