package dataloader

import (
	"context"
	"time"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/util/aerrors"
)

func NewUserLoader(
	ctx context.Context,
	repo user.Repository,
) *generated.UserLoader {
	return generated.NewUserLoader(generated.UserLoaderConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(userUIDs []string) ([]*user.User, []error) {
			users, err := repo.GetByUIDs(ctx, userUIDs)
			if err != nil {
				errs := make([]error, len(userUIDs))
				for i := range make([]int, len(userUIDs)) {
					errs[i] = aerrors.Wrap(err)
				}
				return nil, errs
			}

			userMap := make(map[string]*user.User, len(users))
			for _, u := range users {
				userMap[u.UID] = u
			}

			ret := make([]*user.User, len(userUIDs))
			for i, id := range userUIDs {
				u, ok := userMap[id]
				if ok {
					ret[i] = u
				}
			}

			return ret, nil
		},
	})
}
