package middleware

import (
	"net/http"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/presentation/graphql/dataloader"
	"github.com/laster18/poi/api/src/util/acontext"
)

func RoomUserCountLoader(repo domain.IRoomUserRepo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dataloader := dataloader.NewRoomUserCountLoader(r.Context(), repo)
			newCtx := acontext.SetRoomUserCountLoader(r.Context(), dataloader)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}

func RoomMessageCountLoader(repo domain.IMessageRepo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dataloader := dataloader.NewRoomMessageCountLoader(r.Context(), repo)
			newCtx := acontext.SetRoomMessageCountLoader(r.Context(), dataloader)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
