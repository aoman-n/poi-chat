package middleware

import (
	"net/http"

	"github.com/laster18/poi/api/src/config"
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

			logger := alog.New(&alog.Conf{
				RequestID: reqID,
				User:      u,
				IP:        ip,
				IsJSON:    config.IsProd(),
				IsCaller:  config.IsDev(),
				Level:     config.Conf.LogLevel,
			})

			newCtx := acontext.SetLogger(r.Context(), logger)

			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
