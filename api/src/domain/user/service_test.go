package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/testutil"
	"github.com/laster18/poi/api/src/util/aerrors"
	"github.com/stretchr/testify/assert"
)

func setupSvc(t *testing.T) (context.Context, user.Service, *user.MockRepository) {
	ctrl := gomock.NewController(t)
	mockRepo := user.NewMockRepository(ctrl)
	svc := user.NewService(mockRepo)
	ctx := testutil.NewMockCtx()

	return ctx, svc, mockRepo
}

func TestUserService_SaveIfNotExists(t *testing.T) {
	ctx, svc, mockRepo := setupSvc(t)

	user := &user.User{
		ID:   0,
		UID:  "uuid",
		Name: "name",
	}

	t.Run("userが存在した場合にはRepositoryのSaveが呼ばれないこと", func(t *testing.T) {
		mockRepo.EXPECT().GetByUID(ctx, user.UID).Return(user, nil)
		mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Times(0)
		_, err := svc.FindOrCreate(ctx, user)
		assert.NoError(t, err)
	})

	t.Run("userが存在しなかった場合にはRepositoryのSaveが呼ばれること", func(t *testing.T) {
		errNotFound := aerrors.New("not_found").SetCode(aerrors.CodeNotFound)
		mockRepo.EXPECT().GetByUID(ctx, user.UID).Return(nil, errNotFound)
		mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Times(1)
		_, err := svc.FindOrCreate(ctx, user)
		assert.NoError(t, err)
	})

	t.Run("Repositoryで予期せぬエラーが発生した場合エラーを返すこと", func(t *testing.T) {
		e := errors.New("unexpected error")
		mockRepo.EXPECT().GetByUID(ctx, user.UID).Return(nil, e)
		_, err := svc.FindOrCreate(ctx, user)
		nextErr := errors.Unwrap(err)
		assert.ErrorIs(t, e, nextErr)
	})
}
