package domain

import "context"

// for Redis
type GlobalUser struct {
	UID       string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
}

type GlobalUserRepo interface {
	Insert(ctx context.Context, u *GlobalUser) error
	Delete(ctx context.Context, uID string) error
	Get(ctx context.Context, uID string) (*GlobalUser, error)
	GetAll(ctx context.Context) ([]*GlobalUser, error)
}
