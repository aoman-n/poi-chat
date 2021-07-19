package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/domain/user"
	"github.com/laster18/poi/api/src/presentation/graphql"
	"github.com/laster18/poi/api/src/presentation/graphql/presenter"
	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/aerrors"
)

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return graphql.UserIDStr(obj.ID), nil
}

func (r *mutationResolver) Move(ctx context.Context, input model.MoveInput) (*model.MovePayload, error) {
	currentUser := acontext.GetUser(ctx)
	currentUserStatus := acontext.GetUserStatus(ctx)

	roomRepo := r.repo.NewRoom()
	roomSvc := r.service.NewRoom()
	userStatusInRoom, err := roomSvc.FindOrNewUserStatus(ctx, currentUser, *currentUserStatus.EnteredRoomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomSvc.FindOrNewUserStatus"))
		return nil, nil
	}

	userStatusInRoom.SetPosition(input.X, input.Y)

	if err := roomRepo.SaveUserStatus(ctx, userStatusInRoom); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.SaveUserStatus"))
		return nil, nil
	}

	return presenter.ToMovePayload(userStatusInRoom), nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	currentUser := acontext.GetUser(ctx)
	return presenter.ToUser(currentUser), nil
}

func (r *queryResolver) OnlineUsers(ctx context.Context) ([]*model.User, error) {
	userRepo := r.repo.NewUser()
	users, err := userRepo.GetOnlineUsers(ctx)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	return presenter.ToUsers(users), nil
}

func (r *roomResolver) Users(ctx context.Context, obj *model.Room) ([]*model.RoomUser, error) {
	roomID, _ := strconv.Atoi(obj.ID)

	roomRepo := r.repo.NewRoom()
	userStatuses, err := roomRepo.GetUserStatuses(ctx, roomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.GetUsers"))
		return nil, nil
	}

	return presenter.ToRoomUsers(userStatuses), nil
}

func (r *roomUserResolver) User(ctx context.Context, obj *model.RoomUser) (*model.User, error) {
	uID, _ := strconv.Atoi(obj.User.ID)
	user, err := acontext.GetUserLoader(ctx).Load(uID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to get user on userLoader"))
		return nil, nil
	}

	return presenter.ToUser(user), nil
}

func (r *roomUserResolver) ID(ctx context.Context, obj *model.RoomUser) (string, error) {
	return graphql.RoomUserIDStr(obj.ID), nil
}

func (r *subscriptionResolver) ActedUserEvent(ctx context.Context) (<-chan model.UserEvent, error) {
	logger := acontext.GetLogger(ctx)
	currentUser := acontext.GetUser(ctx)

	userRepo := r.repo.NewUser()
	userSvc := r.service.NewUser()
	exists, err := userSvc.ExistsStatus(ctx, currentUser.UID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to userSvc.ExistsStatus"))
		return nil, nil
	}
	if !exists {
		status := user.NewStatus(currentUser)
		if err := userRepo.SaveStatus(ctx, currentUser.UID, status); err != nil {
			graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to userRepo.SaveStatus"))
			return nil, nil
		}
	}

	ch := make(chan model.UserEvent)
	r.globalUserSubscriber.AddCh(ch, currentUser.UID)

	go func() {
		<-ctx.Done()
		r.globalUserSubscriber.RemoveCh(currentUser.UID)

		if err := userRepo.DeleteStatus(context.Background(), currentUser.UID); err != nil {
			logger.WarnWithErr(err, "failed to delete globalUser")
		}
	}()

	return ch, nil
}

func (r *subscriptionResolver) ActedRoomUserEvent(
	ctx context.Context,
	roomID string,
) (<-chan model.RoomUserEvent, error) {
	logger := acontext.GetLogger(ctx)
	currentUser := acontext.GetUser(ctx)

	domainRoomID, err := graphql.DecodeRoomID(roomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "roomId is invalid format"))
		return nil, nil
	}

	// 対象roomの存在チェック
	roomRepo := r.repo.NewRoom()
	_, err = roomRepo.GetByID(ctx, domainRoomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.GetByID"))
		return nil, nil
	}

	ch := make(chan model.RoomUserEvent)
	r.roomUserSubscriber.AddCh(ch, domainRoomID, currentUser.UID)

	userRepo := r.repo.NewUser()
	userStatus := user.NewStatus(currentUser)
	userStatus.ChangeEnteredRoom(domainRoomID)
	if err := userRepo.SaveStatus(ctx, currentUser.UID, userStatus); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to userRepo.SaveStatus"))
		return nil, nil
	}

	roomSvc := r.service.NewRoom()
	userStatusInRoom, err := roomSvc.FindOrNewUserStatus(ctx, currentUser, domainRoomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomSvc.FindOrNewUserStatus"))
		return nil, nil
	}

	if err := roomRepo.SaveUserStatus(ctx, userStatusInRoom); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomUserRepo.Save"))
		return nil, nil
	}

	go func() {
		<-ctx.Done()
		r.roomUserSubscriber.RemoveCh(domainRoomID, currentUser.UID)
		if err := roomRepo.DeleteUserStatus(context.Background(), userStatusInRoom); err != nil {
			logger.Warnf("failed to roomRepo.DeleteUserStatus, err: %v", err)
		}

		// TODO: ステータスが存在する場合に更新するようにする - userSvc.UpdateStatusIfExits(...)
		userStatus.LeaveRoom()
		if err := userRepo.SaveStatus(context.Background(), currentUser.UID, userStatus); err != nil {
			logger.Warnf("failed to userRepo.SaveStatus, err: %v", err)
		}
	}()

	return ch, nil
}

func (r *Resolver) RemoveLastMessage(
	ctx context.Context,
	input model.RemoveLastMessageInput,
) (*model.RemoveLastMessagePayload, error) {
	currentUser := acontext.GetUser(ctx)
	currentUserStatus := acontext.GetUserStatus(ctx)

	roomRepo := r.repo.NewRoom()
	roomSvc := r.service.NewRoom()
	userStatusInRoom, err := roomSvc.FindOrNewUserStatus(ctx, currentUser, *currentUserStatus.EnteredRoomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomSvc.FindOrNewUserStatus"))
		return nil, nil
	}
	userStatusInRoom.RemoveMessgae()

	if err := roomRepo.SaveUserStatus(ctx, userStatusInRoom); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.SaveUserStatus"))
		return nil, nil
	}

	return presenter.ToRemoveLastMessagePayload(userStatusInRoom), nil
}

func (r *Resolver) ChangeBalloonPosition(
	ctx context.Context,
	input model.ChangeBalloonPositionInput,
) (*model.ChangeBalloonPositionPayload, error) {
	currentUser := acontext.GetUser(ctx)
	currentUserStatus := acontext.GetUserStatus(ctx)

	roomRepo := r.repo.NewRoom()
	roomSvc := r.service.NewRoom()
	userStatusInRoom, err := roomSvc.FindOrNewUserStatus(ctx, currentUser, *currentUserStatus.EnteredRoomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomSvc.FindOrNewUserStatus"))
		return nil, nil
	}

	var domainBalloonPos room.BalloonPosition
	switch input.BalloonPosition {
	case model.BalloonPositionTopRight:
		domainBalloonPos = room.BalloonPositionTopLeft
	case model.BalloonPositionTopLeft:
		domainBalloonPos = room.BalloonPositionTopLeft
	case model.BalloonPositionBottomRight:
		domainBalloonPos = room.BalloonPositionBottomRight
	case model.BalloonPositionBottomLeft:
		domainBalloonPos = room.BalloonPositionBottomLeft
	default:
		m := fmt.Sprintf("%s is unknown balloon position", string(input.BalloonPosition))
		graphql.HandleErr(ctx, aerrors.New(m).SetCode(aerrors.CodeBadParams).Message(m))
	}

	userStatusInRoom.ChangeBalloonPosition(domainBalloonPos)

	if err := roomRepo.SaveUserStatus(ctx, userStatusInRoom); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.SaveUserStatus"))
		return nil, nil
	}

	return presenter.ToChangeBalloonPositionPayload(userStatusInRoom), nil
}

func (r *userResolver) EnteredRoom(ctx context.Context, obj *model.User) (*model.Room, error) {
	// TODO: IDを渡せるようにする
	userStatus, err := acontext.GetUserStatusLoader(ctx).Load(obj.ID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to load user status from dataloader"))
		return nil, nil
	}

	if userStatus.EnteredRoomID == nil {
		return nil, nil
	}

	room, err := acontext.GetRoomLoader(ctx).Load(*userStatus.EnteredRoomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to get room from dataloader"))
		return nil, nil
	}

	return presenter.ToRoom(room), nil
}

func (r *exitedPayloadResolver) UserID(ctx context.Context, obj *model.ExitedPayload) (string, error) {
	return graphql.UserIDStr(obj.UserID), nil
}

func (r *Resolver) ExitedPayload() generated.ExitedPayloadResolver { return &exitedPayloadResolver{r} }
func (r *Resolver) RoomUser() generated.RoomUserResolver           { return &roomUserResolver{r} }
func (r *Resolver) User() generated.UserResolver                   { return &userResolver{r} }

type exitedPayloadResolver struct{ *Resolver }
type roomUserResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
