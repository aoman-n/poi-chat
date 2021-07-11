package repository_test

import (
	"context"
	"testing"

	"github.com/laster18/poi/api/src/domain/message"
	"github.com/laster18/poi/api/src/repository"
	"github.com/laster18/poi/api/src/testutil"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupMessageRepo(t *testing.T) (context.Context, *gorm.DB, message.Repository) {
	t.Helper()

	ctx := testutil.NewMockCtx()
	rdb := testutil.SetupRDB(t)
	repo := repository.NewMessage(rdb)
	return ctx, rdb, repo
}

func Test_Message_Create(t *testing.T) {
	ctx, rdb, repo := setupMessageRepo(t)

	user := testutil.CreateUser(t, rdb, 1)
	room := testutil.CreateRoom(t, rdb, 1, user.ID)

	message := &message.Message{
		UserID: user.ID,
		RoomID: room.ID,
		Body:   "message!",
	}

	assert.NoError(t, repo.Create(ctx, message))
}

func Test_Message_Count(t *testing.T) {
	ctx, rdb, repo := setupMessageRepo(t)

	user := testutil.CreateUser(t, rdb, 1)
	room := testutil.CreateRoom(t, rdb, 1, user.ID)

	// messageを10件作成
	for i := range make([]int, 10) {
		testutil.CreateMessage(t, rdb, i, user.ID, room.ID)
	}

	messageCount, err := repo.Count(ctx, room.ID)
	assert.NoError(t, err)
	assert.Equal(t, 10, messageCount)
}

// func Test_MessageRepo_CountByRoomIDs(t *testing.T) {
// 	ctx, rdb, repo := setupRepo(t)

// 	// userを1つ作成
// 	user := testutil.CreateUser(t, rdb, 1)

// 	// roomを2つ作成
// 	room1 := testutil.CreateRoom(t, rdb, 1)
// 	room2 := testutil.CreateRoom(t, rdb, 2)

// 	// room1へmessageを3つ保存
// 	testutil.CreateMessage(t, rdb, 1, user.ID, room1.ID)
// 	testutil.CreateMessage(t, rdb, 2, user.ID, room1.ID)
// 	testutil.CreateMessage(t, rdb, 3, user.ID, room1.ID)
// 	// room2へmessageを5つ保存
// 	testutil.CreateMessage(t, rdb, 4, user.ID, room2.ID)
// 	testutil.CreateMessage(t, rdb, 5, user.ID, room2.ID)
// 	testutil.CreateMessage(t, rdb, 6, user.ID, room2.ID)
// 	testutil.CreateMessage(t, rdb, 7, user.ID, room2.ID)
// 	testutil.CreateMessage(t, rdb, 8, user.ID, room2.ID)

// 	// got, err := repo.CountByRoomIDs(ctx, []int{room1.ID, room2.ID})

// 	// assert.NoError(t, err)
// 	// assert.Equal(t, []int{3, 5}, got)
// }
