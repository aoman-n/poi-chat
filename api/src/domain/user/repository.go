package user

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, u *User) error
	Get(ctx context.Context, id int) (*User, error)
	GetByUID(ctx context.Context, uid string) (*User, error)
	Online(ctx context.Context, u *User) error
	Offline(ctx context.Context, u *User) error
	GetOnlineUsers(ctx context.Context) ([]*User, error)
}
