package acontext

import (
	"context"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/src/domain"
)

type key string

const (
	userKey                key = "user"
	roomUserCountLoaderKey key = "roomUserCountLoader"
	requestIDKey           key = "requestId"
)

func SetUser(c context.Context, u *domain.GlobalUser) context.Context {
	return context.WithValue(c, userKey, u)
}

func GetUser(c context.Context) *domain.GlobalUser {
	if c.Value(userKey) == nil {
		return nil
	}

	user, ok := c.Value(userKey).(*domain.GlobalUser)
	if !ok || user.UID == "" {
		return nil
	}

	return user
}

func SetRoomUserCountLoader(c context.Context, l *generated.RoomUserCountLoader) context.Context {
	return context.WithValue(c, roomUserCountLoaderKey, l)
}

func GetRoomUserCountLoader(c context.Context) *generated.RoomUserCountLoader {
	l := c.Value(roomUserCountLoaderKey)
	if l == nil {
		// TODO: return error
		panic("must inject roomUserCountLoader")
	}

	loader, ok := l.(*generated.RoomUserCountLoader)
	if !ok {
		// TODO: return error
		panic("roomUserCountLoeader is different type on context")
	}

	return loader
}

func SetRequestID(c context.Context, i string) context.Context {
	return context.WithValue(c, requestIDKey, i)
}

func GetRequestID(c context.Context) string {
	i := c.Value(requestIDKey)
	if i == nil {
		return ""
	}

	id, ok := i.(string)
	if !ok {
		return ""
	}
	return id
}
