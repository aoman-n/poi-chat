package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/repository"
	"github.com/laster18/poi/api/src/testutil"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupRepo(t *testing.T) (context.Context, *gorm.DB, domain.IMessageRepo) {
	t.Helper()

	ctx := testutil.NewMockCtx()
	rdb := testutil.SetupRDB(t)
	repo := repository.NewMessageRepo(rdb)
	return ctx, rdb, repo
}

func Test_MessageRepo_CountByRoomIDs(t *testing.T) {
	ctx, rdb, repo := setupRepo(t)

	// roomを2つ作成
	room1 := createRoom(t, rdb, 1)
	room2 := createRoom(t, rdb, 2)

	// room1へmessageを3つ保存
	createMessage(t, rdb, 1, room1.ID)
	createMessage(t, rdb, 2, room1.ID)
	createMessage(t, rdb, 3, room1.ID)
	// room2へmessageを5つ保存
	createMessage(t, rdb, 4, room2.ID)
	createMessage(t, rdb, 5, room2.ID)
	createMessage(t, rdb, 6, room2.ID)
	createMessage(t, rdb, 7, room2.ID)
	createMessage(t, rdb, 8, room2.ID)

	got, err := repo.CountByRoomIDs(ctx, []int{room1.ID, room2.ID})

	assert.NoError(t, err)
	assert.Equal(t, []int{3, 5}, got)
}

func createRoom(t *testing.T, rdb *gorm.DB, index int) *domain.Room {
	room := domain.Room{
		Name:            fmt.Sprintf("room%d", index),
		BackgroundURL:   fmt.Sprintf("https://example.com/%d", index),
		BackgroundColor: "#ffffff",
		UserCount:       0,
	}

	if err := rdb.Create(&room).Error; err != nil {
		t.Fatalf("failed to setup, create room error: %#v", err)
	}

	return &room
}

func createMessage(t *testing.T, rdb *gorm.DB, index, roomID int) *domain.Message {
	msg := domain.Message{
		UserUID:       "1",
		Body:          fmt.Sprintf("メッセージ%d", index),
		UserName:      fmt.Sprintf("名無しさん%d", index),
		UserAvatarURL: fmt.Sprintf("http://localhost:3000/image/%d", index),
		RoomID:        roomID,
	}

	if err := rdb.Create(&msg).Error; err != nil {
		t.Fatalf("failed to setup, create message error: %v", err)
	}

	return &msg
}
