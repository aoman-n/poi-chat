package graphql

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/laster18/poi/api/src/util/aerrors"
)

type Prefix string

var (
	RoomPrefix       Prefix = "Room"
	MessagePrefix    Prefix = "Message"
	UserPrefix       Prefix = "User"
	RoomUserPrefix   Prefix = "RoomUser"
	GlobalUserPrefix Prefix = "GlobalUser"
)

var (
	// ex: "Room:<id>"
	roomIDFormat    = fmt.Sprintf("%s:%%s", RoomPrefix)
	messageIDFormat = fmt.Sprintf("%s:%%s", MessagePrefix)
	userIDFormat    = fmt.Sprintf("%s:%%s", UserPrefix)
)

func encodeID(prefix Prefix, id int) string {
	return fmt.Sprintf("%s:%d", string(prefix), id)
}

func encodeIDStr(prefix Prefix, idStr string) string {
	return fmt.Sprintf("%s:%s", string(prefix), idStr)
}

func RoomIDStr(id string) string {
	return encodeIDStr(RoomPrefix, id)
}

func MessageIDStr(id string) string {
	return encodeIDStr(MessagePrefix, id)
}

func RoomUserIDStr(id string) string {
	return encodeIDStr(RoomUserPrefix, id)
}

func GlobalUserIDStr(id string) string {
	return encodeIDStr(GlobalUserPrefix, id)
}

func decodeID(prefix Prefix, id string) (int, error) {
	idParts := strings.Split(id, ":")
	if !strings.HasPrefix(id, string(prefix)) || len(idParts) != 2 {
		return 0, aerrors.Errorf(invalidIDMsg, id).SetCode(aerrors.CodeBadParams)
	}

	retID, err := strconv.Atoi(idParts[1])
	if err != nil {
		return 0, aerrors.Errorf(invalidIDMsg, id).SetCode(aerrors.CodeBadParams)
	}

	return retID, nil
}

func decodeIDStr(prefix Prefix, id string) (string, error) {
	idParts := strings.Split(id, ":")
	if !strings.HasPrefix(id, string(prefix)) || len(idParts) != 2 {
		return "", aerrors.Errorf(invalidIDMsg, id).SetCode(aerrors.CodeBadParams)
	}

	return idParts[1], nil
}

func DecodeRoomID(graphID string) (int, error) {
	return decodeID(RoomPrefix, graphID)
}
