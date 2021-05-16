package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/laster18/poi/api/src/delivery"
	"github.com/laster18/poi/api/src/domain"
)

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

			userUID, _ := decodeIDStr(UserPrefix, user.ID)

			ctx := context.WithValue(r.Context(), CurrentUserKey, &domain.GlobalUser{
				UID:       userUID,
				Name:      user.Name,
				AvatarURL: user.AvatarURL,
			})

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func handleSessionError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "server error")
}

func GetCurrentUser(ctx context.Context) (*domain.GlobalUser, error) {
	errNoUserInContext := errors.New("no user in context")
	if ctx.Value(CurrentUserKey) == nil {
		return nil, errNoUserInContext
	}

	user, ok := ctx.Value(CurrentUserKey).(*domain.GlobalUser)
	if !ok || user.UID == "" {
		return nil, errNoUserInContext
	}

	return user, nil
}

// --------------------------
// TODO: 下記は消す

func decodeIDStr(prefix Prefix, id string) (string, error) {
	idParts := strings.Split(id, ":")
	if !strings.HasPrefix(id, string(prefix)) || len(idParts) != 2 {
		return "", fmt.Errorf("invalid id %q", id)
	}

	return idParts[1], nil
}

type Prefix string

// TODO: ':'は定数に入れないようにする
var (
	roomPrefix     Prefix = "Room:"
	messagePrefix  Prefix = "Message:"
	UserPrefix     Prefix = "User:"
	roomUserPrefix Prefix = "RoomUser:"
)
