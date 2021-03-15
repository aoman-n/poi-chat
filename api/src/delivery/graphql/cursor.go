package graphql

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/laster18/poi/api/src/domain"
)

func getCursors(nodes []domain.INode, prefix CursorPrefix) (startCursor *string, endCursor *string) {
	if len(nodes) == 0 {
		return nil, nil
	}

	startNode := nodes[0]
	endNode := nodes[len(nodes)-1]

	return encodeCursor(prefix, startNode.GetID(), startNode.GetCreatedAtUnix()),
		encodeCursor(prefix, endNode.GetID(), endNode.GetCreatedAtUnix())
}

func getRoomCursors(rooms []*domain.Room) (startCursor *string, endCursor *string) {
	var nodes = make([]domain.INode, len(rooms))
	for i, room := range rooms {
		nodes[i] = room
	}

	return getCursors(nodes, roomPrefix)
}

func getMessageCursors(messages []*domain.Message) (startCursor *string, endCursor *string) {
	var nodes = make([]domain.INode, len(messages))
	for i, msg := range messages {
		nodes[i] = msg
	}

	return getCursors(nodes, roomPrefix)
}

func encodeCursor(prefix CursorPrefix, id, unix int) *string {
	cursor := fmt.Sprintf(string(prefix)+"%d:%d", id, unix)
	return &cursor
}

func decodeCursor() {
}

type CursorPrefix string

var (
	roomPrefix    CursorPrefix = "Room:"
	messagePrefix CursorPrefix = "Message:"
)

func getListParts(cursorPrefix CursorPrefix, cursor *string) (lastKnownID int, lastKnownUnix int, err error) {
	idParts := strings.Split(*cursor, ":")
	if !strings.HasPrefix(*cursor, "Room:") || len(idParts) != 3 {
		return 0, 0, fmt.Errorf(invalidIDMsg, *cursor)
	}

	id, err := strconv.Atoi(idParts[1])
	if err != nil {
		return 0, 0, fmt.Errorf(invalidIDMsg, *cursor)
	}

	unix, err := strconv.Atoi(idParts[2])
	if err != nil {
		return 0, 0, fmt.Errorf(invalidIDMsg, *cursor)
	}

	return id, unix, nil
}
