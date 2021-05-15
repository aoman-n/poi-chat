package graphql

import (
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
