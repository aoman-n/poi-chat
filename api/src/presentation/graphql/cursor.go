package graphql

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/domain/message"
	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/util/aerrors"
)

var (
	// ex: "Message:<id>:<unixTimestamp>"
	RoomCursorFormat    = fmt.Sprintf("%s:%%s:%%s", RoomPrefix)
	MessageCursorFormat = fmt.Sprintf("%s:%%s:%%s", MessagePrefix)
)

var (
	invalidIDMsg = "invalid id format: %s"
)

func encodeCursors(nodes []domain.INode, prefix Prefix) (startCursor *string, endCursor *string) {
	if len(nodes) == 0 {
		return nil, nil
	}

	startNode := nodes[0]
	endNode := nodes[len(nodes)-1]

	return encodeCursor(prefix, startNode.GetID(), startNode.GetCreatedAtUnix()),
		encodeCursor(prefix, endNode.GetID(), endNode.GetCreatedAtUnix())
}

func encodeCursor(prefix Prefix, id, unix int) *string {
	cursor := fmt.Sprintf("%s:%d:%d", string(prefix), id, unix)
	return &cursor
}

func MessageCursor(id, unix int) *string {
	return encodeCursor(MessagePrefix, id, unix)
}

func RoomCursor(id, unix int) *string {
	return encodeCursor(RoomPrefix, id, unix)
}

func RoomCursors(rooms []*room.Room) (startCursor *string, endCursor *string) {
	var nodes = make([]domain.INode, len(rooms))
	for i, room := range rooms {
		nodes[i] = room
	}

	return encodeCursors(nodes, RoomPrefix)
}

func MessageCursors(messages []*message.Message) (startCursor *string, endCursor *string) {
	var nodes = make([]domain.INode, len(messages))
	for i, msg := range messages {
		nodes[i] = msg
	}

	return encodeCursors(nodes, MessagePrefix)
}

func decodeCursor(cursorPrefix Prefix, cursor *string) (objID int, objUnix int, err error) {
	cursorParts := strings.Split(*cursor, ":")
	if !strings.HasPrefix(*cursor, string(cursorPrefix)) || len(cursorParts) != 3 {
		return 0, 0, aerrors.Errorf(invalidIDMsg, *cursor).SetCode(aerrors.CodeBadParams)
	}

	id, err := strconv.Atoi(cursorParts[1])
	if err != nil {
		return 0, 0, aerrors.Errorf(invalidIDMsg, *cursor).SetCode(aerrors.CodeBadParams)
	}

	unix, err := strconv.Atoi(cursorParts[2])
	if err != nil {
		return 0, 0, aerrors.Errorf(invalidIDMsg, *cursor).SetCode(aerrors.CodeBadParams)
	}

	return id, unix, nil
}

func DecodeRoomCursor(cursor *string) (objID int, objUnix int, err error) {
	return decodeCursor(RoomPrefix, cursor)
}

func DecodeMessageCursor(cursor *string) (objID int, objUnix int, err error) {
	return decodeCursor(MessagePrefix, cursor)
}
