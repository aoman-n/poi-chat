package registry

import (
	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/domain/user"
)

type Service interface {
	NewUser() user.Service
	NewRoom() room.Service
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

func (s *serviceImpl) NewRoom() room.Service {
	return room.NewService(s.repo.NewRoom())
}
