package presenter

import (
	"strconv"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/domain/user"
)

func ToMovePayload(ru *room.UserStatus) *model.MovePayload {
	return &model.MovePayload{
		// RoomUser: ToRoomUser(ru),
	}
}

func ToRemoveLastMessagePayload(us *room.UserStatus) *model.RemoveLastMessagePayload {
	return &model.RemoveLastMessagePayload{
		// RoomUser: ToRoomUser(ru),
	}
}

func ToChangeBalloonPositionPayload(us *room.UserStatus) *model.ChangeBalloonPositionPayload {
	return &model.ChangeBalloonPositionPayload{
		// RoomUser: ToRoomUser(ru),
	}
}

func ToRoomUser(u *room.UserStatus) *model.RoomUser {
	return &model.RoomUser{
		X: u.X,
		Y: u.Y,
		// LastMessage:     ToMessage(ru.LastMessage),
		BalloonPosition: ConvertBalloonPosition(u.BalloonPosition),
	}
}

func ToUser(u *user.User) *model.User {
	return &model.User{
		ID:         strconv.Itoa(u.ID),
		Name:       u.Name,
		AvatarURL:  u.AvatarURL,
		JoinedRoom: &model.Room{},
	}
}

func ToUsers(users []*user.User) []*model.User {
	ret := make([]*model.User, len(users))
	for i, u := range users {
		ret[i] = ToUser(u)
	}

	return ret
}

func ToRoomUsers(users []*user.User) []*model.RoomUser2 {
	roomUsers := make([]*model.RoomUser2, len(users))
	for i, u := range users {
		roomUser := &model.RoomUser2{
			ID:     strconv.Itoa(u.ID),
			User:   ToUser(u),
			Status: &model.RoomUserStatus{},
		}
		// roomUser := &model.RoomUser2{
		// 	ID:              r.UID,
		// 	Name:            r.Name,
		// 	AvatarURL:       r.AvatarURL,
		// 	X:               r.X,
		// 	Y:               r.Y,
		// 	BalloonPosition: ConvertBalloonPosition(r.BalloonPosition),
		// }
		// if r.LastMessage != nil {
		// 	roomUser.LastMessage = &model.Message{
		// 		ID:            strconv.Itoa(r.LastMessage.ID),
		// 		UserID:        r.LastMessage.UserUID,
		// 		UserName:      r.LastMessage.UserName,
		// 		UserAvatarURL: r.LastMessage.UserAvatarURL,
		// 		Body:          r.LastMessage.Body,
		// 		CreatedAt:     r.LastMessage.CreatedAt,
		// 	}
		// }
		roomUsers[i] = roomUser
	}

	return roomUsers
}

func ToJoinedPayload(ru *room.UserStatus) *model.JoinedPayload {
	return &model.JoinedPayload{
		// RoomUser: ToRoomUser(ru),
	}
}

func ToMovedPayload(ru *room.UserStatus) *model.MovedPayload {
	return &model.MovedPayload{
		// RoomUser: ToRoomUser(ru),
	}
}

func ToSentMassagePayload(ru *room.UserStatus) *model.SentMassagePayload {
	return &model.SentMassagePayload{
		// RoomUser: ToRoomUser(ru),
	}
}

func ToRemovedLastMessagePayload(ru *room.UserStatus) *model.RemovedLastMessagePayload {
	return &model.RemovedLastMessagePayload{
		// RoomUser: ToRoomUser(ru),
	}
}

func ToChangedBalloonPositionPayload(ru *room.UserStatus) *model.ChangedBalloonPositionPayload {
	return &model.ChangedBalloonPositionPayload{
		// RoomUser: ToRoomUser(ru),
	}
}

func ConvertBalloonPosition(p room.BalloonPosition) model.BalloonPosition {
	switch p {
	case room.BalloonPositionTopRight:
		return model.BalloonPositionTopRight
	case room.BalloonPositionTopLeft:
		return model.BalloonPositionTopLeft
	case room.BalloonPositionBottomRight:
		return model.BalloonPositionBottomRight
	case room.BalloonPositionBottomLeft:
		return model.BalloonPositionBottomLeft
	default:
		return model.BalloonPositionTopRight
	}
}
