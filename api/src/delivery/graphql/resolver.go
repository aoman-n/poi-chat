package graphql

import (
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/repository"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(db *gorm.DB) *Resolver {
	roomRepo := repository.NewRoomRepo(db)
	messageRepo := repository.NewMessageRepo(db)
	joinedUserRepo := repository.NewJoinedUserRepo(db)
	subscripters := NewSubscripters()

	// add mock room
	subscripters.Add("Room:1")

	return &Resolver{
		roomRepo:       roomRepo,
		messageRepo:    messageRepo,
		joinedUserRepo: joinedUserRepo,
		subscripters:   subscripters,
	}
}

type Resolver struct {
	// db        *infrastructure.Db
	// roomRepoF func(db *infrastructure.Db) *repository.RoomRepo
	roomRepo       domain.IRoomRepo
	messageRepo    domain.IMessageRepo
	joinedUserRepo domain.IJoinedUserRepo
	subscripters   *Subscripters
}
