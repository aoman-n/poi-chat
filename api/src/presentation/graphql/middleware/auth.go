package middleware

import (
	"fmt"
	"net/http"

	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/session"
)

// Authorize inject user struct when exists user info in cookie
func Authorize() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := acontext.GetLogger(r.Context())

			sess, err := session.GetUserSession(r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "server error")
				return
			}

			if sess.IsNew() {
				next.ServeHTTP(w, r)
				return
			}

			u, err := sess.GetUser()
			logger.Debugf("u: %v, err: %v \n", u, err)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			logger.Debugf("authenticated user is %+v\n", u)

			newCtx := acontext.SetUser(r.Context(), u)

			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
