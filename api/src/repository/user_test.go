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

func TestUser_Online_Offline(t *testing.T) {
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

	// create user1
	user2 := &user.User{
		ID:        0,
		UID:       "uid2",
		Name:      "hoge2",
		AvatarURL: "http://example.com/avatar.png",
		Provider:  user.ProviderTwitter,
	}
	assert.NoError(t, repo.Save(mockCtx, user2))

	// user1とuser2をonlineにする
	assert.NoError(t, repo.Online(mockCtx, user1))
	assert.NoError(t, repo.Online(mockCtx, user2))

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
		assert.NoError(t, repo.Offline(mockCtx, user1))

		expectOnlineUsers := []*user.User{
			user2,
		}
		onlineUsers, err := repo.GetOnlineUsers(mockCtx)
		assert.NoError(t, err)
		assert.Equal(t, expectOnlineUsers, onlineUsers)
	})
}
