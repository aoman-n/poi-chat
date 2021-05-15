package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/laster18/poi/api/src/delivery"
)

type key string

// CurrentUserKey for middleware
const CurrentUserKey key = "currentUser"

// AuthMiddleware inject user struct when exists user info in cookie
func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sess, err := delivery.GetUserSession(r)
			if err != nil {
				log.Printf("session get error in auth middleware, err: %v", err)
				handleSessionError(w)
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

			fmt.Printf("authed user: %+v\n", user)

			ctx := context.WithValue(r.Context(), CurrentUserKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func handleSessionError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "server error")
}

func GetCurrentUserFromCtx(ctx context.Context) (*delivery.User, error) {
	errNoUserInContext := errors.New("no user in context")
	if ctx.Value(CurrentUserKey) == nil {
		return nil, errNoUserInContext
	}

	user, ok := ctx.Value(CurrentUserKey).(*delivery.User)
	if !ok || user.ID == "" {
		return nil, errNoUserInContext
	}

	return user, nil
}
