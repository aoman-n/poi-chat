package domain

import "context"

type JoinedUser struct {
	ID          int
	RoomID      int
	AvatarURL   string
	DisplayName string
	UserID      string
	X           int
	Y           int
}

type IJoinedUserRepo interface {
	Create(ctx context.Context, joinedUser *JoinedUser) error
	List(ctx context.Context, roomID int) ([]*JoinedUser, error)
	Delete(ctx context.Context, joinedUser *JoinedUser) error
}
