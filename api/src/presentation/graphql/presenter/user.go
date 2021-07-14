package presenter

import (
	"strconv"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/domain/user"
)

func ToMovePayload(us *room.UserStatus) *model.MovePayload {
	return &model.MovePayload{
		RoomUser: ToRoomUser(us),
	}
}

func ToRemoveLastMessagePayload(us *room.UserStatus) *model.RemoveLastMessagePayload {
	return &model.RemoveLastMessagePayload{
		RoomUser: ToRoomUser(us),
	}
}

func ToChangeBalloonPositionPayload(us *room.UserStatus) *model.ChangeBalloonPositionPayload {
	return &model.ChangeBalloonPositionPayload{
		RoomUser: ToRoomUser(us),
	}
}

func ToRoomUser(u *room.UserStatus) *model.RoomUser {
	return &model.RoomUser{
		ID:              u.UserUID,
		X:               u.X,
		Y:               u.Y,
		LastMessage:     ToMessage(u.LastMessage),
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

func ToRoomUsers(statuses []*room.UserStatus) []*model.RoomUser {
	roomUsers := make([]*model.RoomUser, len(statuses))
	for i, s := range statuses {
		roomUsers[i] = ToRoomUser(s)
	}

	return roomUsers
}

func ToOnlinedPayload(u *user.User) *model.OnlinedPayload {
	return &model.OnlinedPayload{
		User: ToUser(u),
	}
}

func ToOfflinedPayload(u *user.User) *model.OfflinedPayload {
	return &model.OfflinedPayload{
		User: ToUser(u),
	}
}

func ToJoinedPayload(us *room.UserStatus) *model.JoinedPayload {
	return &model.JoinedPayload{
		RoomUser: ToRoomUser(us),
	}
}

func ToMovedPayload(us *room.UserStatus) *model.MovedPayload {
	return &model.MovedPayload{
		RoomUser: ToRoomUser(us),
	}
}

func ToSentMassagePayload(us *room.UserStatus) *model.SentMassagePayload {
	return &model.SentMassagePayload{
		RoomUser: ToRoomUser(us),
	}
}

func ToRemovedLastMessagePayload(us *room.UserStatus) *model.RemovedLastMessagePayload {
	return &model.RemovedLastMessagePayload{
		RoomUser: ToRoomUser(us),
	}
}

func ToChangedBalloonPositionPayload(us *room.UserStatus) *model.ChangedBalloonPositionPayload {
	return &model.ChangedBalloonPositionPayload{
		RoomUser: ToRoomUser(us),
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
