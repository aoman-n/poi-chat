package dataloader

import (
	"context"
	"fmt"
	"time"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/util/aerrors"
)

func NewRoomLoader(
	ctx context.Context,
	repo room.Repository,
) *generated.RoomLoader {
	return generated.NewRoomLoader(generated.RoomLoaderConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(ids []int) ([]*room.Room, []error) {
			fmt.Printf("ids: %v \n", ids)
			rooms, err := repo.GetByIDs(ctx, ids)
			if err != nil {
				errs := make([]error, len(ids))
				for i := range make([]int, len(ids)) {
					errs[i] = aerrors.Wrap(err)
				}
				return nil, errs
			}

			roomMap := make(map[int]*room.Room, len(rooms))
			for _, r := range rooms {
				roomMap[r.ID] = r
			}

			ret := make([]*room.Room, len(ids))
			for i, id := range ids {
				r, ok := roomMap[id]
				if ok {
					ret[i] = r
				}
			}

			return ret, nil
		},
	})
}
