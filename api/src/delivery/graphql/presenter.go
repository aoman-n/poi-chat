package graphql

import (
	"strconv"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
)

func toMovePayload(ru *domain.RoomUser) *model.MovePayload {
	return &model.MovePayload{
		UserID: ru.UID,
		X:      ru.X,
		Y:      ru.Y,
	}
}

func toMessage(m *domain.Message) *model.Message {
	return &model.Message{
		ID:            strconv.Itoa(m.ID),
		UserID:        m.UserUID,
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
			ID:        g.UID,
			Name:      g.Name,
			AvatarURL: g.AvatarURL,
		}
	}

	return os
}

func toRoomUsers(rus []*domain.RoomUser) []*model.RoomUser {
	roomUsers := make([]*model.RoomUser, len(rus))
	for i, r := range rus {
		roomUser := &model.RoomUser{
			ID:        r.UID,
			Name:      r.Name,
			AvatarURL: r.AvatarURL,
			X:         r.X,
			Y:         r.Y,
		}
		if r.LastMessage != nil {
			roomUser.LastMessage = &model.Message{
				ID:            strconv.Itoa(r.LastMessage.ID),
				UserID:        r.LastMessage.UserUID,
				UserName:      r.LastMessage.UserName,
				UserAvatarURL: r.LastMessage.UserAvatarURL,
				Body:          r.LastMessage.Body,
				CreatedAt:     r.LastMessage.CreatedAt,
			}
		}
		roomUsers[i] = roomUser
	}

	return roomUsers
}
