package registry

import (
	"github.com/laster18/poi/api/src/domain/message"
	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/repository"
	"gorm.io/gorm"
)

type Repository interface {
	NewUser() user.Repository
	NewRoom() room.Repository
	NewMessage() message.Repository
}

type repositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) Repository {
	return &repositoryImpl{db, redis}
}

func (r *repositoryImpl) NewUser() user.Repository {
	return repository.NewUser(r.db, r.redis)
}

func (r *repositoryImpl) NewRoom() room.Repository {
	return repository.NewRoom(r.db, r.redis)
}

func (r *repositoryImpl) NewMessage() message.Repository {
	return repository.NewMessage(r.db)
}
