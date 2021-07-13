package user

import (
	"context"
	"errors"

	"github.com/laster18/poi/api/src/util/aerrors"
)

type Service interface {
	SaveIfNotExists(context.Context, *User) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) SaveIfNotExists(ctx context.Context, u *User) error {
	_, err := s.repo.GetByUID(ctx, u.UID)
	if err == nil {
		return nil
	}

	var aerr *aerrors.ErrApp
	if !errors.As(err, &aerr) {
		return aerrors.Wrap(err)
	}

	if aerr.Code() == aerrors.CodeNotFound {
		if err := s.repo.Save(ctx, u); err != nil {
			return aerrors.Wrap(err)
		}
	}

	return nil
}
