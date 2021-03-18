package graphql

import (
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/repository"
	"gorm.io/gorm"
)

var (
	invalidIDMsg = "invalid id format: %s"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(db *gorm.DB) *Resolver {
	roomRepo := repository.NewRoomRepo(db)
	messageRepo := repository.NewMessageRepo(db)
	subscripters := NewSubscripters()

	// add mock room
	subscripters.Add("Room:1")

	return &Resolver{
		roomRepo:     roomRepo,
		messageRepo:  messageRepo,
		subscripters: subscripters,
	}
}

type Resolver struct {
	// db        *infrastructure.Db
	// roomRepoF func(db *infrastructure.Db) *repository.RoomRepo
	roomRepo     domain.IRoomRepo
	messageRepo  domain.IMessageRepo
	subscripters *Subscripters
}
