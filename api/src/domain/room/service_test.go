package room_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/testutil"
	"github.com/laster18/poi/api/src/util/aerrors"
	"github.com/stretchr/testify/assert"
)

func sertupSvc(t *testing.T) (context.Context, room.Service, *room.MockRepository) {
	ctrl := gomock.NewController(t)
	mockRepo := room.NewMockRepository(ctrl)
	svc := room.NewService(mockRepo)
	ctx := testutil.NewMockCtx()

	return ctx, svc, mockRepo
}

func TestRoomService_FindOrNewUserStatus(t *testing.T) {
	ctx, svc, mockRepo := sertupSvc(t)

	t.Run("userStatusが見つかった場合はそのuserStatusを返すこと", func(t *testing.T) {
		roomID := 2222
		userUID := "hogehoge_user_uidddd"

		user := &user.User{
			ID:        roomID,
			UID:       userUID,
			Name:      "hogehoge",
			AvatarURL: "http://localhost:3000/avatar.png",
			Provider:  user.ProviderTwitter,
		}
		userStatus := &room.UserStatus{
			RoomID:          roomID,
			UserUID:         userUID,
			X:               200,
			Y:               300,
			LastMessage:     nil,
			LastEvent:       room.AddMessageEvent,
			BalloonPosition: room.BalloonPositionBottomLeft,
		}

		mockRepo.EXPECT().GetUserStatus(ctx, roomID, userUID).Return(userStatus, nil)
		actual, err := svc.FindOrNewUserStatus(ctx, user, roomID)
		assert.NoError(t, err)
		assert.Equal(t, userStatus, actual)
	})

	t.Run("userStatusが見つからなかった場合はデフォルトのuserStatusを作成し返すこと", func(t *testing.T) {
		roomID := 2222
		userUID := "hogehoge_user_uidddd"

		user := &user.User{
			ID:        roomID,
			UID:       userUID,
			Name:      "hogehoge",
			AvatarURL: "http://localhost:3000/avatar.png",
			Provider:  user.ProviderTwitter,
		}

		defaultUserStatus := room.NewUserStatus(user, roomID)

		errNotFound := aerrors.New("not found").SetCode(aerrors.CodeNotFound)
		mockRepo.EXPECT().
			GetUserStatus(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, errNotFound)

		actual, err := svc.FindOrNewUserStatus(ctx, user, roomID)

		assert.NoError(t, err)
		assert.Equal(t, defaultUserStatus, actual)
	})
}
