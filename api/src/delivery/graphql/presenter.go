package graphql

import (
	"strconv"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
)

func toCreateRoomPayload(r *domain.Room) *model.CreateRoomPayload {
	return &model.CreateRoomPayload{
		Room: &model.Room{
			ID:        strconv.Itoa(r.ID),
			Name:      r.Name,
			UserCount: r.UserCount,
			BgColor:   r.BackgroundColor,
			BgURL:     r.BackgroundURL,
			CreatedAt: r.CreatedAt,
		},
	}
}

func toMovePayload(ru *domain.RoomUser) *model.MovePayload {
	return &model.MovePayload{
		RoomUser: toRoomUser(ru),
	}
}

func toRoomUser(ru *domain.RoomUser) *model.RoomUser {
	return &model.RoomUser{
		ID:          ru.UID,
		Name:        ru.Name,
		AvatarURL:   ru.AvatarURL,
		X:           ru.X,
		Y:           ru.Y,
		LastMessage: toMessage(ru.LastMessage),
	}
}

func toMessage(m *domain.Message) *model.Message {
	if m == nil {
		return nil
	}

	return &model.Message{
		ID:            strconv.Itoa(m.ID),
		UserID:        m.UserUID,
		UserName:      m.UserName,
		UserAvatarURL: m.UserAvatarURL,
		Body:          m.Body,
		CreatedAt:     m.CreatedAt,
	}
}

func toSendMessagePayload(m *domain.Message) *model.SendMassagePaylaod {
	return &model.SendMassagePaylaod{
		Message: toMessage(m),
	}
}

func toGlobalUsers(gs []*domain.GlobalUser) []*model.GlobalUser {
	os := make([]*model.GlobalUser, len(gs))
	for i, g := range gs {
		os[i] = &model.GlobalUser{
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

func toMessageConnection(before *string, resp *domain.MessageListResp, totalCount int) *model.MessageConnection {
	// create pageInfo
	hasNextPage := false
	if before != nil {
		hasNextPage = true
	}
	startCursor, endCursor := getMessageCursors(resp.List)

	pageInfo := &model.PageInfo{
		StartCursor:     startCursor,
		EndCursor:       endCursor,
		HasNextPage:     hasNextPage,
		HasPreviousPage: resp.HasPreviousPage,
	}

	// create nodes and edges
	nodes := make([]*model.Message, len(resp.List))
	edges := make([]*model.MessageEdge, len(resp.List))

	for i, message := range resp.List {
		nodes[i] = toMessage(message)
		edges[i] = &model.MessageEdge{
			Cursor: *encodeCursor(messagePrefix, message.GetID(), message.GetCreatedAtUnix()),
			Node:   nodes[i],
		}
	}

	return &model.MessageConnection{
		PageInfo:     pageInfo,
		Nodes:        nodes,
		Edges:        edges,
		MessageCount: totalCount,
	}
}

func toRoomConnection(after *string, resp *domain.RoomListResp, total int) *model.RoomConnection {
	// create pageInfo
	hasPrevious := false
	if after != nil {
		hasPrevious = true
	}
	startCursor, endCursor := getRoomCursors(resp.List)
	pageInfo := model.PageInfo{
		StartCursor:     startCursor,
		EndCursor:       endCursor,
		HasNextPage:     resp.HasNext,
		HasPreviousPage: hasPrevious,
	}

	// create nodes and edges
	nodes := make([]*model.Room, len(resp.List))
	edges := make([]*model.RoomEdge, len(resp.List))
	for i, room := range resp.List {
		nodes[i] = &model.Room{
			ID:        strconv.Itoa(int(room.ID)),
			Name:      room.Name,
			UserCount: 0,
			BgColor:   room.BackgroundColor,
			BgURL:     room.BackgroundURL,
			CreatedAt: room.CreatedAt,
		}
		edges[i] = &model.RoomEdge{
			Cursor: *encodeCursor(roomPrefix, room.ID, int(room.CreatedAt.Unix())),
			Node:   nodes[i],
		}
	}

	return &model.RoomConnection{
		PageInfo:  &pageInfo,
		Nodes:     nodes,
		Edges:     edges,
		RoomCount: total,
	}
}
