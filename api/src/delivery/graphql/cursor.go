package graphql

import (
	"fmt"

	"github.com/laster18/poi/api/src/domain"
)

func getRoomCursors(rooms []*domain.Room) (startCursor *string, endCursor *string) {
	if len(rooms) == 0 {
		return nil, nil
	}

	startRoom := rooms[0]
	endRoom := rooms[len(rooms)-1]

	return encodeCursor("Room:", startRoom.ID, int(startRoom.CreatedAt.Unix())),
		encodeCursor("Room:", endRoom.ID, int(endRoom.CreatedAt.Unix()))
}

func encodeCursor(prefix string, id, unix int) *string {
	cursor := fmt.Sprintf(prefix+"%d:%d", id, unix)
	return &cursor
}

func decodeCursor() {
}
