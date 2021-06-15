package domain

import (
	"context"
)

// for Redis
type RoomUser struct {
	UID             string   `json:"id"`
	RoomID          int      `json:"roomId"`
	Name            string   `json:"name"`
	AvatarURL       string   `json:"avatarUrl"`
	X               int      `json:"x"`
	Y               int      `json:"y"`
	LastMessage     *Message `json:"lastMessage"`
	LastEvent       RoomUserEvent
	BalloonPosition BalloonPosition
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
		// default balloon position is TopRight
		BalloonPosition: TopRight,
	}
}

func (r *RoomUser) SetPosition(x, y int) {
	r.LastEvent = MoveEvent
	r.X = x
	r.Y = y
}

func (r *RoomUser) SetMessage(msg *Message) {
	r.LastMessage = msg
	r.LastEvent = AddMessageEvent
}

const (
	DefaultX = 100
	DefaultY = 100
)

type RoomUserEvent int

const (
	JoinEvent RoomUserEvent = iota + 1
	MoveEvent
	AddMessageEvent
	RemoveLastMessageEvent
	ChangeBalloonPositionEvent
)

func (r RoomUserEvent) String() string {
	switch r {
	case JoinEvent:
		return "join_event"
	case MoveEvent:
		return "move_event"
	case AddMessageEvent:
		return "add_message_event"
	case RemoveLastMessageEvent:
		return "remove_last_message_event"
	case ChangeBalloonPositionEvent:
		return "change_balloon_position_event"
	default:
		return "unknown_event"
	}
}

type BalloonPosition int

const (
	TopLeft BalloonPosition = iota + 1
	TopRight
	BottomLeft
	BottomRight
)

func (r BalloonPosition) String() string {
	switch r {
	case TopLeft:
		return "top_left"
	case TopRight:
		return "top_right"
	case BottomLeft:
		return "bottom_left"
	case BottomRight:
		return "bottoom_right"
	default:
		return "unknown_balloon_position"
	}
}

type IRoomUserRepo interface {
	Save(ctx context.Context, u *RoomUser) error
	Delete(ctx context.Context, u *RoomUser) error
	Get(ctx context.Context, roomID int, uID string) (*RoomUser, error)
	GetByRoomID(ctx context.Context, roomID int) ([]*RoomUser, error)
	Counts(ctx context.Context, roomIDs []int) ([]int, error)
}
