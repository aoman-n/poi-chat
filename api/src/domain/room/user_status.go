package room

import (
	"github.com/laster18/poi/api/src/domain/message"
	"github.com/laster18/poi/api/src/domain/user"
)

type UserStatus struct {
	RoomID          int
	UserID          int
	UserUID         string
	X               int
	Y               int
	LastMessage     *message.Message
	LastEvent       Event
	BalloonPosition BalloonPosition
}

const (
	DefaultX = 100
	DefaultY = 100
)

func NewUserStatus(u *user.User, roomID int) *UserStatus {
	return &UserStatus{
		RoomID:      roomID,
		UserID:      u.ID,
		UserUID:     u.UID,
		X:           DefaultX,
		Y:           DefaultY,
		LastMessage: nil,
		LastEvent:   JoinEvent,
		// default balloon position is TopRight
		BalloonPosition: BalloonPositionTopRight,
	}
}

func (u *UserStatus) SetPosition(x, y int) {
	u.X = x
	u.Y = y
	u.LastEvent = MoveEvent
}

func (u *UserStatus) RemoveMessgae() {
	u.LastMessage = nil
	u.LastEvent = RemoveLastMessageEvent
}

func (u *UserStatus) SetMessage(msg *message.Message) {
	u.LastMessage = msg
	u.LastEvent = AddMessageEvent
}

func (u *UserStatus) ChangeBalloonPosition(p BalloonPosition) {
	u.BalloonPosition = p
	u.LastEvent = ChangeBalloonPositionEvent
}

type Event int

const (
	JoinEvent Event = iota + 1
	MoveEvent
	AddMessageEvent
	RemoveLastMessageEvent
	ChangeBalloonPositionEvent
)

func (r Event) String() string {
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
	BalloonPositionTopLeft BalloonPosition = iota + 1
	BalloonPositionTopRight
	BalloonPositionBottomLeft
	BalloonPositionBottomRight
)

func (r BalloonPosition) String() string {
	switch r {
	case BalloonPositionTopLeft:
		return "top_left"
	case BalloonPositionTopRight:
		return "top_right"
	case BalloonPositionBottomLeft:
		return "bottom_left"
	case BalloonPositionBottomRight:
		return "bottoom_right"
	default:
		return "unknown_balloon_position"
	}
}
