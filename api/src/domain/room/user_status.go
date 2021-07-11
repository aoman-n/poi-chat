package room

import "github.com/laster18/poi/api/src/domain"

type UserStatus struct {
	UserUID         string
	X               int
	Y               int
	LastMessage     *domain.Message
	LastEvent       Event
	BalloonPosition BalloonPosition
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
