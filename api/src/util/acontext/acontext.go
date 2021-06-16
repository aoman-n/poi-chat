package acontext

import (
	"context"

	"github.com/go-chi/chi/middleware"
	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/util/alog"
)

type key string

const (
	userKey                   key = "user"
	roomUserCountLoaderKey    key = "roomUserCountLoader"
	roomMessageCountLoaderKey key = "roomMessageCountLoader"
	loggerKey                 key = "logger"
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
		return nil
	}

	return loader
}

func SetRoomMessageCountLoader(c context.Context, l *generated.RoomMessageCountLoader) context.Context {
	return context.WithValue(c, roomMessageCountLoaderKey, l)
}

func GetRoomMessageCountLoader(c context.Context) *generated.RoomMessageCountLoader {
	l := c.Value(roomMessageCountLoaderKey)
	if l == nil {
		return nil
	}

	loader, ok := l.(*generated.RoomMessageCountLoader)
	if !ok {
		return nil
	}

	return loader
}

func GetRequestID(c context.Context) string {
	return middleware.GetReqID(c)
}

func SetLogger(c context.Context, l alog.Logger) context.Context {
	return context.WithValue(c, loggerKey, l)
}

func GetLogger(c context.Context) alog.Logger {
	l := c.Value(loggerKey)
	if l == nil {
		return alog.DefaultLogger
	}

	if logger, ok := l.(alog.Logger); ok {
		return logger
	}

	return alog.DefaultLogger
}
