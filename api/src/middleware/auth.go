package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/session"
)

// Authorize inject user struct when exists user info in cookie
func Authorize() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sess, err := session.GetUserSession(r)
			if err != nil {
				log.Printf("session get error in auth middleware, err: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "server error")
				return
			}

			if sess.IsNew() {
				next.ServeHTTP(w, r)
				return
			}

			user, err := sess.GetUser()
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			logger := acontext.GetLogger(r.Context())
			logger.Debugf("authenticated user is %+v\n", user)

			newCtx := acontext.SetUser(r.Context(), &domain.GlobalUser{
				UID:       user.ID,
				Name:      user.Name,
				AvatarURL: user.AvatarURL,
			})

			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
