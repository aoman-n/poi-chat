package middleware

import (
	"context"
	"net/http"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/src/dataloader"
	"github.com/laster18/poi/api/src/domain"
)

const (
	RoomUserCountLoaderKey key = "roomUserCountLoader"
)

func InjectRoomUserCountLoader(repo domain.IRoomUserRepo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			loader := dataloader.NewRoomUserCountLoader(r.Context(), repo)
			ctx := context.WithValue(r.Context(), RoomUserCountLoaderKey, loader)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetRoomUserCountLoader(ctx context.Context) *generated.RoomUserCountLoader {
	l := ctx.Value(RoomUserCountLoaderKey)
	if l == nil {
		// TODO: return error
		panic("must inject roomUserCountLoader")
	}

	loader, ok := l.(*generated.RoomUserCountLoader)
	if !ok {
		// TODO: return error
		panic("roomUserCountLoeader is different type on ctx")
	}

	return loader
}
