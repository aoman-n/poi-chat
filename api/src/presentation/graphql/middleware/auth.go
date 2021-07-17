package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/aerrors"
	"github.com/laster18/poi/api/src/util/session"
)

// Authorize inject user struct when exists user info in cookie
func Authorize(userRepo user.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := acontext.GetLogger(r.Context())

			sess, err := session.GetUserSession(r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "server error")
				logger.WarnWithErr(err, "failed to session.GetUserSession")
				return
			}

			if sess.IsNew() {
				next.ServeHTTP(w, r)
				return
			}

			uid, err := sess.GetUserUID()
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			u, err := userRepo.GetByUID(r.Context(), uid)
			if err != nil {
				var errApp *aerrors.ErrApp
				if errors.As(err, &errApp) {
					if errApp.Code() == aerrors.CodeNotFound {
						next.ServeHTTP(w, r)
						return
					}
				}

				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "server error")
				logger.WarnWithErr(err, "failed to userRepo.GetByUID")
				return
			}

			newCtx := acontext.SetUser(r.Context(), u)

			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
