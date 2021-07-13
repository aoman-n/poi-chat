package registry

import (
	"github.com/laster18/poi/api/src/domain/user"
)

type Service interface {
	NewUser() user.Service
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo}
}

func (s *serviceImpl) NewUser() user.Service {
	return user.NewService(s.repo.NewUser())
}
