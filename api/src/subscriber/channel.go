package subscriber

import (
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
func MakeRoomUserKey(roomID, userID int) string {
	return fmt.Sprintf("%s:%d:%d", RoomUserChannel, roomID, userID)
}

var roomUserChannelChannelReg = regexp.MustCompile(RoomUserChannel + `:([\d]+):(\d+)`)

func destructRoomUserKey(key string) (roomID, userID int, err error) {
	matches := roomUserChannelChannelReg.FindStringSubmatch(key)
	roomID, err = strconv.Atoi(matches[1])
	if err != nil {
		return
	}
	userID, err = strconv.Atoi(matches[2])
	if err != nil {
		return
	}
	return
}

func removeKeyspacePrefix(key string) string {
	prefix := redis.KeySpace + ":"
	if strings.HasPrefix(key, prefix) {
		return strings.TrimPrefix(key, prefix)
	}
	return key
}
