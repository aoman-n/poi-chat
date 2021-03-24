package domain

import (
	"context"
	"time"
)

type Message struct {
	ID            int
	UserUID       string
	Body          string
	UserName      string
	UserAvatarURL string
	RoomID        int
	CreatedAt     time.Time
}

var _ INode = (*Message)(nil)

func (r *Message) GetID() int {
	return r.ID
}

func (r *Message) GetCreatedAtUnix() int {
	return int(r.CreatedAt.Unix())
}

type MessageListReq struct {
	RoomID        int
	Limit         int
	LastKnownID   int
	LastKnownUnix int
}

type MessageListResp struct {
	List            []*Message
	HasPreviousPage bool
}

type IMessageRepo interface {
	List(ctx context.Context, req *MessageListReq) (*MessageListResp, error)
	Create(ctx context.Context, message *Message) error
	Count(ctx context.Context, roomID int) (int, error)
}
