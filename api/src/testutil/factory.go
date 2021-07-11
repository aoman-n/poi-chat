package testutil

import (
	"fmt"
	"testing"

	"github.com/laster18/poi/api/src/domain/message"
	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/domain/user"
	"gorm.io/gorm"
)

func CreateUser(t *testing.T, rdb *gorm.DB, index int) *user.User {
	user := user.User{
		UID:       fmt.Sprintf("uid%d", index),
		Name:      fmt.Sprintf("名無し%d", index),
		AvatarURL: fmt.Sprintf("https://example.com/%d", index),
		Provider:  user.ProviderTwitter,
	}

	if err := rdb.Create(&user).Error; err != nil {
		t.Fatalf("failed to setup, create user error: %#v", err)
	}

	return &user
}

func CreateRoom(t *testing.T, rdb *gorm.DB, index int, uID int) *room.Room {
	room := room.Room{
		OwnerID:         uID,
		Name:            fmt.Sprintf("room%d", index),
		BackgroundURL:   fmt.Sprintf("https://example.com/%d", index),
		BackgroundColor: "#ffffff",
	}

	if err := rdb.Create(&room).Error; err != nil {
		t.Fatalf("failed to setup, create room error: %#v", err)
	}

	return &room
}

func CreateMessage(t *testing.T, rdb *gorm.DB, index, userID int, roomID int) *message.Message {
	msg := message.Message{
		UserID: userID,
		RoomID: roomID,
		Body:   fmt.Sprintf("メッセージ%d", index),
	}

	if err := rdb.Create(&msg).Error; err != nil {
		t.Fatalf("failed to setup, create message error: %v", err)
	}

	return &msg
}
