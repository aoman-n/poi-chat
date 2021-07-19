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

// onlineUser:<userID>
var onlineUserChannelReg = regexp.MustCompile(OnlineUserChannel + `:(\d+)`)

func DestructOnlineUserKey(key string) (userID int, err error) {
	matches := onlineUserChannelReg.FindStringSubmatch(key)

	if len(matches) == 0 {
		return 0, fmt.Errorf(`onlineUserKey is invalid format", key: "%s"`, key)
	}

	userIDStr := matches[1]
	if userIDStr == "" {
		return 0, fmt.Errorf(`onlineUserKey is invalid format", key: "%s"`, key)
	}

	userID, err = strconv.Atoi(userIDStr)
	if err != nil {
		return 0, fmt.Errorf(`onlineUserKey is invalid format", key: "%s"`, key)
	}

	return
}

// MakeRoomUserKey "onlineUser:<userID>"
func MakeOnlineUserKey(userID int) string {
	return fmt.Sprintf("%s:%d", OnlineUserChannel, userID)
}

// --------------------------------------------
// --------------------------------------------
// For RoomUserStatus

// MakeRoomUserStatusKey "roomUserStatus:<roomID>:<userID>"
func MakeRoomUserStatusKey(roomID int, userID int) string {
	return fmt.Sprintf("%s:%d:%d", RoomUserStatusChannel, roomID, userID)
}

// roomUserStatus:<roomID>:<userID>
var roomUserStatusChannelReg = regexp.MustCompile(RoomUserStatusChannel + `:(\d+):(\d+)`)

func DestructRoomUserStatusKey(key string) (roomID int, userID int, err error) {
	matches := roomUserStatusChannelReg.FindStringSubmatch(key)

	if len(matches) == 0 {
		return 0, 0, fmt.Errorf(`roomUserStatus is invalid format", key: "%s"`, key)
	}

	roomIDStr := matches[1]
	userIDStr := matches[2]
	if roomIDStr == "" || userIDStr == "" {
		return 0, 0, fmt.Errorf(`roomUserStatus is invalid format", key: "%s"`, key)
	}

	roomID, err = strconv.Atoi(roomIDStr)
	if err != nil {
		return 0, 0, fmt.Errorf(`roomUserStatus is invalid format", key: "%s"`, key)
	}

	userID, err = strconv.Atoi(userIDStr)
	if err != nil {
		return 0, 0, fmt.Errorf(`roomUserStatus is invalid format", key: "%s"`, key)
	}

	return
}
