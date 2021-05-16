package domain

import (
	"context"
)

// for Redis
type RoomUser struct {
	UID         string   `json:"id"`
	RoomID      int      `json:"roomId"`
	Name        string   `json:"name"`
	AvatarURL   string   `json:"avatarUrl"`
	X           int      `json:"x"`
	Y           int      `json:"y"`
	LastMessage *Message `json:"lastMessage"`
	LastEvent   RoomUserEvent
}

func NewDefaultRoomUser(roomID int, u *GlobalUser) *RoomUser {
	return &RoomUser{
		UID:         u.UID,
		RoomID:      roomID,
		Name:        u.Name,
		AvatarURL:   u.AvatarURL,
		X:           DefaultX,
		Y:           DefaultY,
		LastMessage: nil,
		LastEvent:   JoinEvent,
	}
}

func (r *RoomUser) SetPosition(x, y int) {
	r.LastEvent = MoveEvent
	r.X = x
	r.Y = y
}

func (r *RoomUser) SetMessage(msg *Message) {
	r.LastMessage = msg
	r.LastEvent = MessageEvent
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
	GetByRoomID(ctx context.Context, roomID int) ([]*RoomUser, error)
	Counts(ctx context.Context, roomIDs []int) ([]int, error)
}
