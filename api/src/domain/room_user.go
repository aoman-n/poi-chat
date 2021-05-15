package domain

import (
	"context"

	"github.com/laster18/poi/api/src/delivery"
)

// for Redis
type RoomUser struct {
	UID         string `json:"id"`
	RoomID      int    `json:"roomId"`
	Name        string `json:"name"`
	AvatarURL   string `json:"avatarUrl"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	LastMessage string `json:"lastMessage"`
	LastEvent   RoomUserEvent
}

func NewDefaultRoomUser(roomID int, u *delivery.User) *RoomUser {
	return &RoomUser{
		UID:         u.ID,
		RoomID:      roomID,
		Name:        u.Name,
		AvatarURL:   u.AvatarURL,
		X:           DefaultX,
		Y:           DefaultY,
		LastMessage: "",
		LastEvent:   JoinEvent,
	}
}

const (
	DefaultX = 100
	DefaultY = 100
)

type RoomUserEvent int

const (
	JoinEvent RoomUserEvent = iota + 1
	MoveEvent
	MessageEvent
)

type IRoomUserRepo interface {
	Insert(ctx context.Context, u *RoomUser) error
	Delete(ctx context.Context, u *RoomUser) error
	Get(ctx context.Context, roomID int, uID string) (*RoomUser, error)
}