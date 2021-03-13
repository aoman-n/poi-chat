package graphql

import (
	"fmt"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infrastructure/db"
	"github.com/laster18/poi/api/src/repository"
)

var (
	roomIDPrefix    = "Room:"
	messageIDPrefix = "Messsage:"
)

var (
	roomIDFormat    = fmt.Sprintf("%s%%s", roomIDPrefix)
	messageIDFormat = fmt.Sprintf("%s%%s", messageIDPrefix)
)

var (
	invalidIDMsg = "invalid id format: %s"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(db *db.Db) *Resolver {
	fmt.Println("hoge")

	roomRepo := repository.NewRoomRepo(db)

	return &Resolver{
		roomRepo: roomRepo,
	}
}

type Resolver struct {
	// db        *infrastructure.Db
	// roomRepoF func(db *infrastructure.Db) *repository.RoomRepo
	roomRepo domain.IRoomRepo
}
