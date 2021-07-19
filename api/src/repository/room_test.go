package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/repository"
	"github.com/laster18/poi/api/src/testutil"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupRoomRepo(t *testing.T) (context.Context, *gorm.DB, *redis.Client, room.Repository) {
	ctx := testutil.NewMockCtx()
	rdb := testutil.SetupRDB(t)
	redis := testutil.SetupRedis(t)
	repo := repository.NewRoom(rdb, redis)

	return ctx, rdb, redis, repo
}

func Test_Room_GetByID(t *testing.T) {
	ctx, rdb, _, repo := setupRoomRepo(t)

	user := testutil.CreateUser(t, rdb, 1)
	room := testutil.CreateRoom(t, rdb, 1, user.ID)

	gettedRoom, err := repo.GetByID(ctx, room.ID)
	assert.NoError(t, err)
	room.CreatedAt = time.Time{}
	gettedRoom.CreatedAt = time.Time{}
	assert.Equal(t, room, gettedRoom)
}

func Test_Room_GetByIDs(t *testing.T) {
	ctx, rdb, _, repo := setupRoomRepo(t)

	user := testutil.CreateUser(t, rdb, 1)
	room1 := testutil.CreateRoom(t, rdb, 1, user.ID)
	room2 := testutil.CreateRoom(t, rdb, 2, user.ID)
	room3 := testutil.CreateRoom(t, rdb, 3, user.ID)
	room1.CreatedAt = time.Time{}
	room2.CreatedAt = time.Time{}
	room3.CreatedAt = time.Time{}

	gettedRooms, err := repo.GetByIDs(ctx, []int{room1.ID, 200, room2.ID, room3.ID})
	assert.NoError(t, err)

	expect := []*room.Room{room1, room2, room3}

	for _, room := range gettedRooms {
		room.CreatedAt = time.Time{}
	}

	assert.Equal(t, expect, gettedRooms)
}

func Test_Room_GetByName(t *testing.T) {
	ctx, rdb, _, repo := setupRoomRepo(t)

	user := testutil.CreateUser(t, rdb, 1)
	room := testutil.CreateRoom(t, rdb, 1, user.ID)

	gettedRoom, err := repo.GetByName(ctx, room.Name)
	assert.NoError(t, err)
	room.CreatedAt = time.Time{}
	gettedRoom.CreatedAt = time.Time{}
	assert.Equal(t, room, gettedRoom)
}

func Test_Room_GetAll(t *testing.T) {
	t.Skip()
}

func Test_Room_List(t *testing.T) {
	t.Skip()
}

func Test_Room_Create(t *testing.T) {
	ctx, rdb, _, repo := setupRoomRepo(t)

	user := testutil.CreateUser(t, rdb, 1)

	room := &room.Room{
		ID:              0,
		OwnerID:         user.ID,
		Name:            "hoge",
		BackgroundURL:   "http://example.com/bg.png",
		BackgroundColor: "#ffffff",
	}

	assert.NoError(t, repo.Create(ctx, room))
}

func Test_Room_Count(t *testing.T) {
	ctx, rdb, _, repo := setupRoomRepo(t)

	user := testutil.CreateUser(t, rdb, 1)
	for i := range make([]int, 5) {
		testutil.CreateRoom(t, rdb, i, user.ID)
	}

	roomCount, err := repo.Count(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 5, roomCount)
}

func createUserAndStatus(ctx context.Context, t *testing.T, rdb *gorm.DB, repo room.Repository, roomID, index int) {
	user := testutil.CreateUser(t, rdb, index)
	userStatus := &room.UserStatus{
		RoomID:          roomID,
		UserID:          index * 100,
		UserUID:         user.UID,
		X:               10 * index,
		Y:               20 * index,
		LastMessage:     nil,
		LastEvent:       room.AddMessageEvent,
		BalloonPosition: room.BalloonPositionBottomLeft,
	}
	err := repo.SaveUserStatus(ctx, userStatus)
	if err != nil {
		t.Fatalf("failed to setup user status, err: %v", err)
	}
}

func Test_Room_CountUserByRoomIDs(t *testing.T) {
	ctx, rdb, _, repo := setupRoomRepo(t)

	// roomID1に5userStatusを保存
	roomID1 := 1000
	createUserAndStatus(ctx, t, rdb, repo, roomID1, 1)
	createUserAndStatus(ctx, t, rdb, repo, roomID1, 2)
	createUserAndStatus(ctx, t, rdb, repo, roomID1, 3)
	createUserAndStatus(ctx, t, rdb, repo, roomID1, 4)
	createUserAndStatus(ctx, t, rdb, repo, roomID1, 5)

	// roomID2に2userStatusを保存
	roomID2 := 2000
	createUserAndStatus(ctx, t, rdb, repo, roomID2, 6)
	createUserAndStatus(ctx, t, rdb, repo, roomID2, 7)

	// roomID3に3userStatusを保存
	roomID3 := 3000
	createUserAndStatus(ctx, t, rdb, repo, roomID3, 8)
	createUserAndStatus(ctx, t, rdb, repo, roomID3, 9)
	createUserAndStatus(ctx, t, rdb, repo, roomID3, 10)

	roomIDs := []int{roomID1, 3333, roomID2, roomID3, 222}
	expected := []int{5, 0, 2, 3, 0}

	actual, err := repo.CountUserByRoomIDs(ctx, roomIDs)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Room_CountMessageByRoomIDs(t *testing.T) {
	ctx, rdb, _, repo := setupRoomRepo(t)

	user := testutil.CreateUser(t, rdb, 1)

	room1 := testutil.CreateRoom(t, rdb, 1, user.ID)
	room2 := testutil.CreateRoom(t, rdb, 2, user.ID)
	room3 := testutil.CreateRoom(t, rdb, 3, user.ID)
	room4 := testutil.CreateRoom(t, rdb, 3, user.ID)

	// room1にmessageを5件作成
	for i := range make([]int, 5) {
		testutil.CreateMessage(t, rdb, i, user.ID, room1.ID)
	}

	// room2にmessageを3件作成
	for i := range make([]int, 3) {
		testutil.CreateMessage(t, rdb, i, user.ID, room2.ID)
	}

	// room4にmessageを10件作成
	for i := range make([]int, 10) {
		testutil.CreateMessage(t, rdb, i, user.ID, room4.ID)
	}

	roomIDs := []int{room1.ID, room2.ID, room3.ID, room4.ID}
	expected := []int{5, 3, 0, 10}

	actual, err := repo.CountMessageByRoomIDs(ctx, roomIDs)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Room_GetUsers(t *testing.T) {
	ctx, rdb, _, repo := setupRoomRepo(t)

	user1 := testutil.CreateUser(t, rdb, 1)
	user2 := testutil.CreateUser(t, rdb, 2)
	user3 := testutil.CreateUser(t, rdb, 3)

	userStatus1 := &room.UserStatus{
		RoomID:          1000,
		UserUID:         user1.UID,
		UserID:          user1.ID,
		X:               100,
		Y:               200,
		LastMessage:     nil,
		LastEvent:       room.AddMessageEvent,
		BalloonPosition: room.BalloonPositionBottomLeft,
	}
	userStatus2 := &room.UserStatus{
		RoomID:          1000,
		UserUID:         user2.UID,
		UserID:          user2.ID,
		X:               100,
		Y:               200,
		LastMessage:     nil,
		LastEvent:       room.AddMessageEvent,
		BalloonPosition: room.BalloonPositionBottomLeft,
	}
	userStatus3 := &room.UserStatus{
		RoomID:          1000,
		UserUID:         user3.UID,
		UserID:          user3.ID,
		X:               100,
		Y:               200,
		LastMessage:     nil,
		LastEvent:       room.AddMessageEvent,
		BalloonPosition: room.BalloonPositionBottomLeft,
	}
	roomID := 1000

	assert.NoError(t, repo.SaveUserStatus(ctx, userStatus1))
	assert.NoError(t, repo.SaveUserStatus(ctx, userStatus2))
	assert.NoError(t, repo.SaveUserStatus(ctx, userStatus3))

	actual, err := repo.GetUsers(ctx, roomID)
	assert.NoError(t, err)

	expectUsers := []*user.User{
		user1,
		user2,
		user3,
	}

	assert.Equal(t, expectUsers, actual)
}

func Test_Room_Save_GetUserStatus(t *testing.T) {
	ctx, _, _, repo := setupRoomRepo(t)

	roomID := 1000
	userStatus := &room.UserStatus{
		UserID:          10000,
		RoomID:          roomID,
		UserUID:         "uid_hoge",
		X:               100,
		Y:               200,
		LastMessage:     nil,
		LastEvent:       room.AddMessageEvent,
		BalloonPosition: room.BalloonPositionBottomLeft,
	}

	assert.NoError(t, repo.SaveUserStatus(ctx, userStatus))

	actual, err := repo.GetUserStatus(ctx, roomID, userStatus.UserID)
	assert.NoError(t, err)
	assert.Equal(t, userStatus, actual)
}

func Test_Room_Delete(t *testing.T) {
	ctx, rdb, _, repo := setupRoomRepo(t)

	user1 := testutil.CreateUser(t, rdb, 1)
	user2 := testutil.CreateUser(t, rdb, 2)
	user3 := testutil.CreateUser(t, rdb, 3)

	roomID := 1000
	userStatus1 := &room.UserStatus{
		RoomID:          roomID,
		UserUID:         user1.UID,
		UserID:          user1.ID,
		X:               100,
		Y:               200,
		LastMessage:     nil,
		LastEvent:       room.AddMessageEvent,
		BalloonPosition: room.BalloonPositionBottomLeft,
	}
	userStatus2 := &room.UserStatus{
		RoomID:          roomID,
		UserUID:         user2.UID,
		UserID:          user2.ID,
		X:               100,
		Y:               200,
		LastMessage:     nil,
		LastEvent:       room.AddMessageEvent,
		BalloonPosition: room.BalloonPositionBottomLeft,
	}
	userStatus3 := &room.UserStatus{
		RoomID:          roomID,
		UserUID:         user3.UID,
		UserID:          user3.ID,
		X:               100,
		Y:               200,
		LastMessage:     nil,
		LastEvent:       room.AddMessageEvent,
		BalloonPosition: room.BalloonPositionBottomLeft,
	}

	assert.NoError(t, repo.SaveUserStatus(ctx, userStatus1))
	assert.NoError(t, repo.SaveUserStatus(ctx, userStatus2))
	assert.NoError(t, repo.SaveUserStatus(ctx, userStatus3))

	// user1とuser2のstatusを削除
	assert.NoError(t, repo.DeleteUserStatus(ctx, userStatus1))
	assert.NoError(t, repo.DeleteUserStatus(ctx, userStatus2))

	actual, err := repo.GetUsers(ctx, roomID)
	assert.NoError(t, err)

	expectUsers := []*user.User{
		user3,
	}

	assert.Equal(t, expectUsers, actual)
}

func Test_Room_GetUserStatuses(t *testing.T) {
	ctx, _, _, repo := setupRoomRepo(t)

	roomID := 1000
	userStatus1 := &room.UserStatus{
		RoomID:          roomID,
		UserID:          100,
		UserUID:         "uid_hoge1",
		X:               100,
		Y:               200,
		LastMessage:     nil,
		LastEvent:       room.AddMessageEvent,
		BalloonPosition: room.BalloonPositionBottomLeft,
	}
	userStatus2 := &room.UserStatus{
		RoomID:          roomID,
		UserID:          200,
		UserUID:         "uid_hoge2",
		X:               100,
		Y:               200,
		LastMessage:     nil,
		LastEvent:       room.AddMessageEvent,
		BalloonPosition: room.BalloonPositionBottomLeft,
	}
	userStatus3 := &room.UserStatus{
		RoomID:          roomID,
		UserID:          300,
		UserUID:         "uid_hoge3",
		X:               100,
		Y:               200,
		LastMessage:     nil,
		LastEvent:       room.AddMessageEvent,
		BalloonPosition: room.BalloonPositionBottomLeft,
	}

	assert.NoError(t, repo.SaveUserStatus(ctx, userStatus1))
	assert.NoError(t, repo.SaveUserStatus(ctx, userStatus2))
	assert.NoError(t, repo.SaveUserStatus(ctx, userStatus3))

	actual, err := repo.GetUserStatuses(ctx, roomID)
	assert.NoError(t, err)

	expected := []*room.UserStatus{
		userStatus1,
		userStatus2,
		userStatus3,
	}

	assert.ElementsMatch(t, expected, actual)
}
