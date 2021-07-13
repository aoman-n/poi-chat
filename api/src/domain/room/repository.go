package room

import (
	"context"

	"github.com/laster18/poi/api/src/domain/user"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (*Room, error)
	GetByName(ctx context.Context, name string) (*Room, error)
	GetAll(ctx context.Context) ([]*Room, error)
	List(ctx context.Context, req *ListReq) (*ListResp, error)
	Create(ctx context.Context, room *Room) error
	Count(ctx context.Context) (int, error)
	CountUserByRoomIDs(ctx context.Context, roomIDs []int) ([]int, error)
	CountMessageByRoomIDs(ctx context.Context, roomIDs []int) ([]int, error)
	GetUsers(ctx context.Context, roomID int) ([]*user.User, error)
	SaveUserStatus(ctx context.Context, status *UserStatus) error
	DeleteUserStatus(ctx context.Context, status *UserStatus) error
	GetUserStatus(ctx context.Context, roomID int, userUID string) (*UserStatus, error)
	GetUserStatuses(ctx context.Context, roomID int, userUIDs []string) ([]*UserStatus, error)
}

type ListReq struct {
	Limit         int
	LastKnownID   int
	LastKnownUnix int
}

type ListResp struct {
	List    []*Room
	HasNext bool
}
