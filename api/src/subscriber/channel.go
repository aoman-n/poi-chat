package subscriber

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/laster18/poi/api/src/infra/redis"
)

func removeKeyspacePrefix(key string) string {
	prefix := redis.KeySpace + ":"
	if strings.HasPrefix(key, prefix) {
		return strings.TrimPrefix(key, prefix)
	}
	return key
}

const (
	GlobalUserChannel = "globalUser"
	RoomUserChannel   = "roomUser"
)

// --------------------------------------------
// --------------------------------------------
// For GlobalUser

// MakeRoomUserKey "globalUser:<userId>"
func MakeGlobalUserKey(userUID string) string {
	return fmt.Sprintf("%s:%s", GlobalUserChannel, userUID)
}

// globalUser:11111
var globalUserChannelReg = regexp.MustCompile(GlobalUserChannel + `:([a-zA-Z\d]+)`)

func destructGlobalUserKey(key string) (userUID string, err error) {
	matches := globalUserChannelReg.FindStringSubmatch(key)

	userUID = matches[1]
	if userUID == "" {
		return "", errors.New(
			fmt.Sprintf(`globalUserKey is invalid format, key: "%s"` + key),
		)
	}

	return
}

// --------------------------------------------
// --------------------------------------------
// For RoomUser

// roomUser:1:335902496
var roomUserChannelReg = regexp.MustCompile(RoomUserChannel + `:(\d+):([a-zA-Z\d]+)`)

func destructRoomUserKey(key string) (roomID int, userUID string, err error) {
	matches := roomUserChannelReg.FindStringSubmatch(key)

	roomIDStr := matches[1]
	userIDStr := matches[2]
	if roomIDStr == "" || userIDStr == "" {
		return 0, "", errors.New(
			fmt.Sprintf(`"roomUserKey is invalid format", key: "%s"` + key),
		)
	}

	roomID, err = strconv.Atoi(roomIDStr)
	if err != nil {
		return
	}
	userUID = userIDStr
	return
}

// MakeRoomUserKey "roomUser:<roomId>:<userId>"
func MakeRoomUserKey(roomID int, userUID string) string {
	return fmt.Sprintf("%s:%d:%s", RoomUserChannel, roomID, userUID)
}
