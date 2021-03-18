package graphql

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/laster18/poi/api/src/domain"
)

type Prefix string

var (
	roomPrefix    Prefix = "Room:"
	messagePrefix Prefix = "Message:"
)

var (
	// ex: "Room:<id>"
	roomIDFormat    = fmt.Sprintf("%s%%s", roomPrefix)
	messageIDFormat = fmt.Sprintf("%s%%s", messagePrefix)
	// ex: "Room:<id>:<unixTimestamp>"
	roomCursorFormat    = fmt.Sprintf("%s%%s:%%s", roomPrefix)
	messageCursorFormat = fmt.Sprintf("%s%%s:%%s", messagePrefix)
)

func getCursors(nodes []domain.INode, prefix Prefix) (startCursor *string, endCursor *string) {
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

	return getCursors(nodes, messagePrefix)
}

func encodeCursor(prefix Prefix, id, unix int) *string {
	cursor := fmt.Sprintf(string(prefix)+"%d:%d", id, unix)
	return &cursor
}

func decodeCursor(cursorPrefix Prefix, cursor *string) (objID int, objUnix int, err error) {
	cursorParts := strings.Split(*cursor, ":")
	if !strings.HasPrefix(*cursor, string(cursorPrefix)) || len(cursorParts) != 3 {
		return 0, 0, fmt.Errorf(invalidIDMsg, *cursor)
	}

	id, err := strconv.Atoi(cursorParts[1])
	if err != nil {
		return 0, 0, fmt.Errorf(invalidIDMsg, *cursor)
	}

	unix, err := strconv.Atoi(cursorParts[2])
	if err != nil {
		return 0, 0, fmt.Errorf(invalidIDMsg, *cursor)
	}

	return id, unix, nil
}

func encodeID(prefix Prefix, id int) string {
	return fmt.Sprintf(string(prefix)+"%d", id)
}

func decodeID(prefix Prefix, id string) (int, error) {
	idParts := strings.Split(id, ":")
	if !strings.HasPrefix(id, string(prefix)) || len(idParts) != 2 {
		return 0, fmt.Errorf(invalidIDMsg, id)
	}

	retID, err := strconv.Atoi(idParts[1])
	if err != nil {
		return 0, fmt.Errorf(invalidIDMsg, id)
	}

	return retID, nil
}
