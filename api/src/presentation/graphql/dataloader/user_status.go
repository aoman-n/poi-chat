package dataloader

import (
	"context"
	"time"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/util/aerrors"
)

func NewUserStatusLoader(
	ctx context.Context,
	repo user.Repository,
) *generated.UserStatusLoader {
	return generated.NewUserStatusLoader(generated.UserStatusLoaderConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(ids []int) ([]*user.Status, []error) {
			statuses, err := repo.GetStatuses(ctx, ids)
			if err != nil {
				errs := make([]error, len(ids))
				for i := range make([]int, len(ids)) {
					errs[i] = aerrors.Wrap(err)
				}
				return nil, errs
			}

			statusMap := make(map[int]*user.Status, len(statuses))
			for _, s := range statuses {
				statusMap[s.UserID] = s
			}

			ret := make([]*user.Status, len(ids))
			for i, id := range ids {
				s, ok := statusMap[id]
				if ok {
					ret[i] = s
				}
			}

			return ret, nil
		},
	})
}
