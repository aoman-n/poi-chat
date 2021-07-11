package message

import "context"

type Repository interface {
	List(ctx context.Context, req *ListReq) (*ListResp, error)
	Create(ctx context.Context, message *Message) error
	Count(ctx context.Context, roomID int) (int, error)
	// CountByRoomIDs(ctx context.Context, roomIDs []int) ([]int, error)
}

type ListReq struct {
	RoomID        int
	Limit         int
	LastKnownID   int
	LastKnownUnix int
}

type ListResp struct {
	List            []*Message
	HasPreviousPage bool
}
