package user

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, u *User) error
	Get(ctx context.Context, id int) (*User, error)
	GetByIDs(ctx context.Context, ids []int) ([]*User, error)
	GetByUID(ctx context.Context, uid string) (*User, error)
	GetByUIDs(ctx context.Context, uids []string) ([]*User, error)
	SaveStatus(ctx context.Context, id int, status *Status) error
	DeleteStatus(ctx context.Context, id int) error
	GetStatus(ctx context.Context, id int) (*Status, error)
	GetStatuses(ctx context.Context, ids []int) ([]*Status, error)
	GetOnlineUsers(ctx context.Context) ([]*User, error)
}
