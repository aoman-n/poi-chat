package message

import (
	"time"

	"github.com/laster18/poi/api/src/domain"
)

type Message struct {
	ID        int
	UserID    int
	RoomID    int
	Body      string
	CreatedAt time.Time
}

var _ domain.INode = (*Message)(nil)

func (r *Message) GetID() int {
	return r.ID
}

func (r *Message) GetCreatedAtUnix() int {
	return int(r.CreatedAt.Unix())
}
