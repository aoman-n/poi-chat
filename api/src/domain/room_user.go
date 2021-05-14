package domain

import "context"

// for Redis
type RoomUser struct {
	ID          int    `json:"id"`
	RoomID      int    `json:"roomId"`
	Name        string `json:"name"`
	AvatarURL   string `json:"avatarUrl"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	LastMessage string `json:"lastMessage"`
	LastEvent   RoomUserEvent
}

type RoomUserEvent int

const (
	JoinEvent RoomUserEvent = iota + 1
	MoveEvent
	MessageEvent
)

type RoomUserRepo interface {
	Create(ctx context.Context, u *RoomUser) error
	Update(ctx context.Context, u *RoomUser) error
	Delete(ctx context.Context, u *RoomUser) error
	Get(ctx context.Context, uID int) (*RoomUser, error)
}
