package subscriber

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/laster18/poi/api/src/infra/redis"
)

const (
	RoomUserChannel = "roomUser"
)

// MakeRoomUserKey "roomUser:<roomId>:<userId>"
func MakeRoomUserKey(roomID int, userUID string) string {
	return fmt.Sprintf("%s:%d:%s", RoomUserChannel, roomID, userUID)
}

var roomUserChannelChannelReg = regexp.MustCompile(RoomUserChannel + `:([\d]+):(\d+)`)

// roomUser:1:User:335902496

func destructRoomUserKey(key string) (roomID int, userUID string, err error) {
	matches := roomUserChannelChannelReg.FindStringSubmatch(key)

	roomIDStr := matches[1]
	userIDStr := matches[2]
	if roomIDStr == "" || userIDStr == "" {
		return 0, "", errors.New("roomUserKey is invalid format")
	}

	roomID, err = strconv.Atoi(roomIDStr)
	if err != nil {
		return
	}
	userUID = userIDStr
	return
}

func removeKeyspacePrefix(key string) string {
	prefix := redis.KeySpace + ":"
	if strings.HasPrefix(key, prefix) {
		return strings.TrimPrefix(key, prefix)
	}
	return key
}
