package middleware

import (
	"net/http"

	"github.com/laster18/poi/api/src/dataloader"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/util/acontext"
)

func RoomUserCountLoader(repo domain.IRoomUserRepo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l := dataloader.NewRoomUserCountLoader(r.Context(), repo)
			newCtx := acontext.SetRoomUserCountLoader(r.Context(), l)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
