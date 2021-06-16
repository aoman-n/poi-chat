package dataloader

import (
	"context"
	"time"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/src/domain"
)

func NewRoomMessageCountLoader(
	ctx context.Context,
	repo domain.IMessageRepo,
) *generated.RoomMessageCountLoader {
	return generated.NewRoomMessageCountLoader(generated.RoomMessageCountLoaderConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(keys []int) ([]int, []error) {
			counts, err := repo.CountByRoomIDs(ctx, keys)

			if err != nil {
				errs := make([]error, len(keys))
				for i := range make([]int, len(keys)) {
					errs[i] = err
				}
				return nil, errs
			}

			return counts, nil
		},
	})
}
