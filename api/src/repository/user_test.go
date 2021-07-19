package repository_test

import (
	"testing"

	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/repository"
	"github.com/laster18/poi/api/src/testutil"
	"github.com/stretchr/testify/assert"
)

func setupUserRepository(t *testing.T) user.Repository {
	db := testutil.SetupRDB(t)
	redis := testutil.SetupRedis(t)
	repo := repository.NewUser(db, redis)

	return repo
}

func TestUser_Save_Get(t *testing.T) {
	mockCtx := testutil.NewMockCtx()
	repo := setupUserRepository(t)

	user := &user.User{
		ID:        0,
		UID:       "uid",
		Name:      "hoge",
		AvatarURL: "http://example.com/avatar.png",
		Provider:  user.ProviderTwitter,
	}

	err := repo.Save(mockCtx, user)
	assert.NoError(t, err)

	u, err := repo.Get(mockCtx, user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, u)

	uu, err := repo.GetByUID(mockCtx, user.UID)
	assert.NoError(t, err)
	assert.Equal(t, user, uu)
}

func TestUser_Save_Delete_Status(t *testing.T) {
	mockCtx := testutil.NewMockCtx()
	repo := setupUserRepository(t)

	// create user1
	user1 := &user.User{
		ID:        0,
		UID:       "uid1",
		Name:      "hoge1",
		AvatarURL: "http://example.com/avatar.png",
		Provider:  user.ProviderTwitter,
	}
	assert.NoError(t, repo.Save(mockCtx, user1))
	user1Status := user.NewStatus(user1)

	// create user1
	user2 := &user.User{
		ID:        0,
		UID:       "uid2",
		Name:      "hoge2",
		AvatarURL: "http://example.com/avatar.png",
		Provider:  user.ProviderTwitter,
	}
	assert.NoError(t, repo.Save(mockCtx, user2))
	user2Status := user.NewStatus(user2)

	// user1とuser2をonlineにする
	assert.NoError(t, repo.SaveStatus(mockCtx, user1.UID, user1Status))
	assert.NoError(t, repo.SaveStatus(mockCtx, user2.UID, user2Status))

	t.Run("オンラインにしたユーザーが取得されること", func(t *testing.T) {
		expectOnlineUsers := []*user.User{
			user1,
			user2,
		}

		onlineUsers, err := repo.GetOnlineUsers(mockCtx)
		assert.NoError(t, err)
		assert.Equal(t, expectOnlineUsers, onlineUsers)

	})

	t.Run("オフラインにしたユーザーは取得されないこと", func(t *testing.T) {
		// user1をofflineにする
		assert.NoError(t, repo.DeleteStatus(mockCtx, user1.UID))

		expectOnlineUsers := []*user.User{
			user2,
		}
		onlineUsers, err := repo.GetOnlineUsers(mockCtx)
		assert.NoError(t, err)
		assert.Equal(t, expectOnlineUsers, onlineUsers)
	})
}

func TestUser_Save_Get_Status(t *testing.T) {
	mockCtx := testutil.NewMockCtx()
	repo := setupUserRepository(t)

	userUID := "hoge_uid"
	roomID := 50000
	status := &user.Status{
		EnteredRoomID: &roomID,
		State:         user.StateNormal,
	}

	assert.NoError(t, repo.SaveStatus(mockCtx, userUID, status))

	actual, err := repo.GetStatus(mockCtx, userUID)
	assert.NoError(t, err)

	assert.Equal(t, status, actual)
}

func TestUser_Save_GetStatuses(t *testing.T) {
	mockCtx := testutil.NewMockCtx()
	repo := setupUserRepository(t)

	roomID := 100

	status1 := &user.Status{
		UserUID:       "user1",
		EnteredRoomID: &roomID,
		State:         user.StateNormal,
	}
	status2 := &user.Status{
		UserUID:       "user2",
		EnteredRoomID: &roomID,
		State:         user.StateNormal,
	}
	status3 := &user.Status{
		UserUID:       "user3",
		EnteredRoomID: &roomID,
		State:         user.StateNormal,
	}

	assert.NoError(t, repo.SaveStatus(mockCtx, status1.UserUID, status1))
	assert.NoError(t, repo.SaveStatus(mockCtx, status2.UserUID, status2))
	assert.NoError(t, repo.SaveStatus(mockCtx, status3.UserUID, status3))

	t.Run("正しくstatusのスライスを取得できること", func(t *testing.T) {
		uids := []string{status1.UserUID, status2.UserUID, status3.UserUID}

		actual, err := repo.GetStatuses(mockCtx, uids)
		assert.NoError(t, err)

		expect := []*user.Status{
			status1,
			status2,
			status3,
		}

		assert.Equal(t, expect, actual)
	})

	t.Run("存在しないuidがあった場合には見つかったものだけが返されること", func(t *testing.T) {
		uids := []string{"notFoundID1", status1.UserUID, status2.UserUID, "notFoundID2", status3.UserUID}

		actual, err := repo.GetStatuses(mockCtx, uids)
		assert.NoError(t, err)

		expect := []*user.Status{
			status1,
			status2,
			status3,
		}

		assert.Equal(t, expect, actual)
	})
}
