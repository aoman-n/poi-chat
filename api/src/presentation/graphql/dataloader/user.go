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
		Fetch: func(userIDs []int) ([]*user.User, []error) {
			users, err := repo.GetByIDs(ctx, userIDs)
			if err != nil {
				errs := make([]error, len(userIDs))
				for i := range make([]int, len(userIDs)) {
					errs[i] = aerrors.Wrap(err)
				}
				return nil, errs
			}

			userMap := make(map[int]*user.User, len(users))
			for _, u := range users {
				userMap[u.ID] = u
			}

			ret := make([]*user.User, len(userIDs))
			for i, id := range userIDs {
				u, ok := userMap[id]
				if ok {
					ret[i] = u
				}
			}

			return ret, nil
		},
	})
}
