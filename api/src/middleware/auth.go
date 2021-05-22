package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/session"
)

// AuthMiddleware inject user struct when exists user info in cookie
func AuthMiddleware() func(http.Handler) http.Handler {
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

			userUID, _ := decodeIDStr(UserPrefix, user.ID)

			newCtx := acontext.SetUser(r.Context(), &domain.GlobalUser{
				UID:       userUID,
				Name:      user.Name,
				AvatarURL: user.AvatarURL,
			})

			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
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
