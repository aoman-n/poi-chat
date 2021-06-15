package presenter

import (
	"strconv"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
)

func ToMovePayload(ru *domain.RoomUser) *model.MovePayload {
	return &model.MovePayload{
		RoomUser: ToRoomUser(ru),
	}
}

func ToRemoveLastMessagePayload(ru *domain.RoomUser) *model.RemoveLastMessagePayload {
	return &model.RemoveLastMessagePayload{
		RoomUser: ToRoomUser(ru),
	}
}

func ToChangeBalloonPositionPayload(ru *domain.RoomUser) *model.ChangeBalloonPositionPayload {
	return &model.ChangeBalloonPositionPayload{
		RoomUser: ToRoomUser(ru),
	}
}

func ToRoomUser(ru *domain.RoomUser) *model.RoomUser {
	return &model.RoomUser{
		ID:              ru.UID,
		Name:            ru.Name,
		AvatarURL:       ru.AvatarURL,
		X:               ru.X,
		Y:               ru.Y,
		LastMessage:     ToMessage(ru.LastMessage),
		BalloonPosition: ConvertBalloonPosition(ru.BalloonPosition),
	}
}

func ToGlobalUsers(gs []*domain.GlobalUser) []*model.GlobalUser {
	os := make([]*model.GlobalUser, len(gs))
	for i, g := range gs {
		os[i] = &model.GlobalUser{
			ID:        g.UID,
			Name:      g.Name,
			AvatarURL: g.AvatarURL,
			Joined:    nil, // TODO: ちゃんと返す
		}
	}

	return os
}

func ToRoomUsers(rus []*domain.RoomUser) []*model.RoomUser {
	roomUsers := make([]*model.RoomUser, len(rus))
	for i, r := range rus {
		roomUser := &model.RoomUser{
			ID:              r.UID,
			Name:            r.Name,
			AvatarURL:       r.AvatarURL,
			X:               r.X,
			Y:               r.Y,
			BalloonPosition: ConvertBalloonPosition(r.BalloonPosition),
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

func ToJoinedPayload(ru *domain.RoomUser) *model.JoinedPayload {
	return &model.JoinedPayload{
		RoomUser: ToRoomUser(ru),
	}
}

func ToMovedPayload(ru *domain.RoomUser) *model.MovedPayload {
	return &model.MovedPayload{
		RoomUser: ToRoomUser(ru),
	}
}

func ToSentMassagePayload(ru *domain.RoomUser) *model.SentMassagePayload {
	return &model.SentMassagePayload{
		RoomUser: ToRoomUser(ru),
	}
}

func ToRemovedLastMessagePayload(ru *domain.RoomUser) *model.RemovedLastMessagePayload {
	return &model.RemovedLastMessagePayload{
		RoomUser: ToRoomUser(ru),
	}
}

func ToChangedBalloonPositionPayload(ru *domain.RoomUser) *model.ChangedBalloonPositionPayload {
	return &model.ChangedBalloonPositionPayload{
		RoomUser: ToRoomUser(ru),
	}
}

func ConvertBalloonPosition(p domain.BalloonPosition) model.BalloonPosition {
	switch p {
	case domain.TopRight:
		return model.BalloonPositionTopRight
	case domain.TopLeft:
		return model.BalloonPositionTopLeft
	case domain.BottomRight:
		return model.BalloonPositionBottomRight
	case domain.BottomLeft:
		return model.BalloonPositionBottomLeft
	default:
		return model.BalloonPositionTopRight
	}
}
