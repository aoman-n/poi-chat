package subscriber

import (
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
	OnlineUserChannel     = "onlineUser"
	RoomUserStatusChannel = "roomUserStatus"
)

// --------------------------------------------
// --------------------------------------------
// For OnlineUser

// onlineUser:<userUID>
var onlineUserChannelReg = regexp.MustCompile(OnlineUserChannel + `:([a-zA-Z\d-]+)`)

func DestructOnlineUserKey(key string) (userUID string, err error) {
	matches := onlineUserChannelReg.FindStringSubmatch(key)

	if len(matches) == 0 {
		return "", fmt.Errorf(`onlineUserKey is invalid format", key: "%s"`, key)
	}

	userIDStr := matches[1]
	if userIDStr == "" {
		return "", fmt.Errorf(`onlineUserKey is invalid format", key: "%s"`, key)
	}

	userUID = userIDStr
	return
}

// MakeRoomUserKey "onlineUser:<userUID>"
func MakeOnlineUserKey(userUID string) string {
	return fmt.Sprintf("%s:%s", OnlineUserChannel, userUID)
}

// --------------------------------------------
// --------------------------------------------
// For RoomUserStatus

// MakeRoomUserStatusKey "roomUserStatus:<roomID>:<userUID>"
func MakeRoomUserStatusKey(roomID int, userUID string) string {
	return fmt.Sprintf("%s:%d:%s", RoomUserStatusChannel, roomID, userUID)
}

// roomUserStatus:<roomID>:<userUID>
var roomUserStatusChannelReg = regexp.MustCompile(RoomUserStatusChannel + `:(\d+):([a-zA-Z\d-]+)`)

func DestructRoomUserStatusKey(key string) (roomID int, userUID string, err error) {
	matches := roomUserStatusChannelReg.FindStringSubmatch(key)

	if len(matches) == 0 {
		return 0, "", fmt.Errorf(`roomUserStatus is invalid format", key: "%s"`, key)
	}

	roomIDStr := matches[1]
	userUIDStr := matches[2]
	if roomIDStr == "" || userUIDStr == "" {
		return 0, "", fmt.Errorf(`roomUserStatus is invalid format", key: "%s"`, key)
	}

	roomID, err = strconv.Atoi(roomIDStr)
	if err != nil {
		return 0, "", fmt.Errorf(`roomUserStatus is invalid format", key: "%s"`, key)
	}

	userUID = userUIDStr
	return
}
