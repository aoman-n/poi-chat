package graphql

import (
	"strconv"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
)

func toMovePayload(ru *domain.RoomUser) *model.MovePayload {
	return &model.MovePayload{
		UserID: encodeIDStr(roomUserPrefix, ru.UID),
		X:      ru.X,
		Y:      ru.Y,
	}
}

func toMessage(m *domain.Message) *model.Message {
	return &model.Message{
		ID:            strconv.Itoa(m.ID),
		UserID:        encodeIDStr(userPrefix, m.UserUID),
		UserName:      m.UserName,
		UserAvatarURL: m.UserAvatarURL,
		Body:          m.Body,
		CreatedAt:     m.CreatedAt,
	}
}

func toOnlineUsers(gs []*domain.GlobalUser) []*model.OnlineUser {
	os := make([]*model.OnlineUser, len(gs))
	for i, g := range gs {
		os[i] = &model.OnlineUser{
			ID:        encodeIDStr(userPrefix, g.UID),
			Name:      g.Name,
			AvatarURL: g.AvatarURL,
		}
	}

	return os
}
