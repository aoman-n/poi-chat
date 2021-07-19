package middleware

import (
	"net/http"

	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/presentation/graphql/dataloader"
	"github.com/laster18/poi/api/src/util/acontext"
)

func RoomUserCountLoader(repo room.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dataloader := dataloader.NewRoomUserCountLoader(r.Context(), repo)
			newCtx := acontext.SetRoomUserCountLoader(r.Context(), dataloader)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}

func RoomMessageCountLoader(repo room.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dataloader := dataloader.NewRoomMessageCountLoader(r.Context(), repo)
			newCtx := acontext.SetRoomMessageCountLoader(r.Context(), dataloader)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}

func UserLoader(repo user.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dataloader := dataloader.NewUserLoader(r.Context(), repo)
			newCtx := acontext.SetUserLoader(r.Context(), dataloader)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}

func UserStatusLoader(repo user.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dataloader := dataloader.NewUserStatusLoader(r.Context(), repo)
			newCtx := acontext.SetUserStatusLoader(r.Context(), dataloader)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}

func RoomLoader(repo room.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dataloader := dataloader.NewRoomLoader(r.Context(), repo)
			newCtx := acontext.SetRoomLoader(r.Context(), dataloader)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
