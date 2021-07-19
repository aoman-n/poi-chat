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
		Fetch: func(uids []string) ([]*user.Status, []error) {
			statuses, err := repo.GetStatuses(ctx, uids)
			if err != nil {
				errs := make([]error, len(uids))
				for i := range make([]int, len(uids)) {
					errs[i] = aerrors.Wrap(err)
				}
				return nil, errs
			}

			statusMap := make(map[string]*user.Status, len(statuses))
			for _, s := range statuses {
				statusMap[s.UserUID] = s
			}

			ret := make([]*user.Status, len(uids))
			for i, uid := range uids {
				s, ok := statusMap[uid]
				if ok {
					ret[i] = s
				}
			}

			return ret, nil
		},
	})
}
