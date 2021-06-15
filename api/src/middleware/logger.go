package middleware

import (
	"net/http"

	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/alog"
)

func Logger() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := acontext.GetRequestID(r.Context())
			user := acontext.GetUser(r.Context())
			ip := r.RemoteAddr

			var u *alog.User
			if user != nil {
				u = &alog.User{
					ID:   user.UID,
					Name: user.Name,
				}
			}

			logger := alog.New(
				alog.WithRequetID(reqID),
				alog.WithIP(ip),
				u,
			)

			newCtx := acontext.SetLogger(r.Context(), logger)

			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
