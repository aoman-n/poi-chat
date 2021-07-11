package user

import (
	"context"
)

type Repository interface {
	// user
	Save(ctx context.Context, u *User) error
	Get(ctx context.Context, id int) (*User, error)
	GetByUID(ctx context.Context, uid string) (*User, error)
	Online(ctx context.Context, u *User) error
	Offline(ctx context.Context, u *User) error
	GetOnlineUsers(ctx context.Context) ([]*User, error)

	// status in room
	// GetStatusesInRoom(ctx context.Context, roomID int) ([]*StatusInRoom, error)
	// SaveStatusInRoom(ctx context.Context, s StatusInRoom) error
	// DeleteStatusInRoom(ctx context.Context, s StatusInRoom) error

	// old roomUserRepo
	// Save(ctx context.Context, u *RoomUser) error
	// Delete(ctx context.Context, u *RoomUser) error
	// Get(ctx context.Context, roomID int, uID string) (*RoomUser, error)
	// GetByRoomID(ctx context.Context, roomID int) ([]*RoomUser, error)
	// Counts(ctx context.Context, roomIDs []int) ([]int, error)

	// old globalUserRepo
	// Save(ctx context.Context, u *GlobalUser) error
	// Delete(ctx context.Context, uID string) error
	// Get(ctx context.Context, uID string) (*GlobalUser, error)
	// GetAll(ctx context.Context) ([]*GlobalUser, error)
}

// type RoomRepository interface {
// 	UserCount()
// 	MessageCount()
// 	GetUsers()
// }
