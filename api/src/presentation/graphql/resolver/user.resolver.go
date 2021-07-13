package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/presentation/graphql"
	"github.com/laster18/poi/api/src/presentation/graphql/presenter"
	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/aerrors"
)

func (r *exitedPayloadResolver) UserID(ctx context.Context, obj *model.ExitedPayload) (string, error) {
	return graphql.RoomUserIDStr(obj.UserID), nil
}

func (r *globalUserResolver) ID(ctx context.Context, obj *model.GlobalUser) (string, error) {
	return graphql.GlobalUserIDStr(obj.ID), nil
}

func (r *globalUserResolver) Joined(ctx context.Context, obj *model.GlobalUser) (*model.Room, error) {
	panic(fmt.Errorf("not impelemented"))
}

func (r *meResolver) ID(ctx context.Context, obj *model.Me) (string, error) {
	// return encodeIDStr(globalUserPrefix, obj.ID), nil
	return graphql.GlobalUserIDStr(obj.ID), nil
}

func (r *mutationResolver) Move(ctx context.Context, input model.MoveInput) (*model.MovePayload, error) {
	logger := acontext.GetLogger(ctx)
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		graphql.HandleErr(ctx, aerrors.Wrap(graphql.ErrUnauthorized))
		return nil, nil
	}

	domainRoomID, err := graphql.DecodeRoomID(input.RoomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	roomRepo := r.repo.NewRoom()
	userStatusInRoom, err := roomRepo.GetUserStatus(ctx, domainRoomID, currentUser.UID)
	if err != nil {
		logger.Infof("failed to get userStatusInRoom err: %v", err)
	}
	if userStatusInRoom == nil {
		userStatusInRoom = room.NewUserStatus(currentUser)
	}
	userStatusInRoom.SetPosition(input.X, input.Y)

	if err := roomRepo.SaveUserStatus(ctx, userStatusInRoom); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.SaveUserStatus"))
		return nil, nil
	}

	return presenter.ToMovePayload(userStatusInRoom), nil
}

func (r *offlinedPayloadResolver) UserID(ctx context.Context, obj *model.OfflinedPayload) (string, error) {
	return graphql.GlobalUserIDStr(obj.UserID), nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.Me, error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		graphql.HandleErr(ctx, aerrors.Wrap(graphql.ErrUnauthorized))
		return nil, nil
	}

	me := &model.Me{
		ID:        currentUser.UID,
		Name:      currentUser.Name,
		AvatarURL: currentUser.AvatarURL,
	}

	return me, nil
}

func (r *queryResolver) GlobalUsers(ctx context.Context) ([]*model.GlobalUser, error) {
	panic("not impl")
}

func (r *queryResolver) OnlineUsers(ctx context.Context) ([]*model.User, error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		graphql.HandleErr(ctx, aerrors.Wrap(graphql.ErrUnauthorized))
		return nil, nil
	}

	userRepo := r.repo.NewUser()
	users, err := userRepo.GetOnlineUsers(ctx)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	return presenter.ToUsers(users), nil
}

func (r *roomResolver) Users(ctx context.Context, obj *model.Room) ([]*model.RoomUser2, error) {
	roomID, _ := strconv.Atoi(obj.ID)

	roomRepo := r.repo.NewRoom()
	users, err := roomRepo.GetUsers(ctx, roomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.GetUsers"))
		return nil, nil
	}

	return presenter.ToRoomUsers(users), nil
}

func (r *roomUserResolver) ID(ctx context.Context, obj *model.RoomUser) (string, error) {
	return graphql.RoomUserIDStr(obj.ID), nil
}

func (r *subscriptionResolver) ActedGlobalUserEvent(ctx context.Context) (<-chan model.GlobalUserEvent, error) {
	logger := acontext.GetLogger(ctx)
	currentUser := acontext.GetUser(ctx)

	if currentUser == nil {
		graphql.HandleErr(ctx, aerrors.Wrap(graphql.ErrUnauthorized))
		return nil, nil
	}

	ch := make(chan model.GlobalUserEvent)
	// TODO: subscriber修正
	r.globalUserSubscriber.AddCh(ch, currentUser.UID)

	userRepo := r.repo.NewUser()
	if err := userRepo.Online(ctx, currentUser); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	go func() {
		<-ctx.Done()
		// TODO: subscriber修正
		r.globalUserSubscriber.RemoveCh(currentUser.UID)

		if err := userRepo.Offline(context.Background(), currentUser); err != nil {
			logger.WarnWithErr(err, "failed to delete globalUser")
		}
	}()

	return ch, nil
}

func (r *subscriptionResolver) ActedRoomUserEvent(
	ctx context.Context,
	roomID string,
) (<-chan model.RoomUserEvent, error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		graphql.HandleErr(ctx, aerrors.Wrap(graphql.ErrUnauthorized))
		return nil, nil
	}

	domainRoomID, err := graphql.DecodeRoomID(roomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "roomId is invalid format"))
		return nil, nil
	}

	// TODO: roomの存在チェック

	ch := make(chan model.RoomUserEvent)
	// TODO: subscriber修正
	r.roomUserSubscriber.AddCh(ch, domainRoomID, currentUser.UID)

	roomRepo := r.repo.NewRoom()
	userStatusInRoom := room.NewUserStatus(currentUser)

	if err := roomRepo.SaveUserStatus(ctx, userStatusInRoom); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomUserRepo.Save"))
		return nil, nil
	}

	go func() {
		<-ctx.Done()
		// TODO: subscriber修正
		r.roomUserSubscriber.RemoveCh(domainRoomID, currentUser.UID)
		if err := roomRepo.DeleteUserStatus(ctx, userStatusInRoom); err != nil {
			log.Println("failed to delete roomUser, err:", err)
		}
	}()

	return ch, nil
}

func (r *Resolver) RemoveLastMessage(
	ctx context.Context,
	input model.RemoveLastMessageInput,
) (*model.RemoveLastMessagePayload, error) {
	logger := acontext.GetLogger(ctx)

	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		graphql.HandleErr(ctx, aerrors.Wrap(graphql.ErrUnauthorized))
		return nil, nil
	}

	domainRoomID, err := graphql.DecodeRoomID(input.RoomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "roomId is invalid format"))
		return nil, nil
	}

	roomRepo := r.repo.NewRoom()

	userStatusInRoom, err := roomRepo.GetUserStatus(ctx, domainRoomID, currentUser.UID)
	if err != nil {
		logger.Infof("failed to get roomUser err: %v", err)
	}
	if userStatusInRoom == nil {
		userStatusInRoom = room.NewUserStatus(currentUser)
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
	logger := acontext.GetLogger(ctx)

	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		graphql.HandleErr(ctx, aerrors.Wrap(graphql.ErrUnauthorized))
		return nil, nil
	}

	domainRoomID, err := graphql.DecodeRoomID(input.RoomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "roomId is invalid format"))
		return nil, nil
	}

	roomRepo := r.repo.NewRoom()
	userStatusInRoom, err := roomRepo.GetUserStatus(ctx, domainRoomID, currentUser.UID)
	if err != nil {
		logger.Infof("failed to get roomUser err: %v", err)
	}
	if userStatusInRoom == nil {
		userStatusInRoom = room.NewUserStatus(currentUser)
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

func (r *userResolver) JoinedRoom(ctx context.Context, obj *model.User) (*model.Room, error) {
	panic("not implemented")
}

func (r *Resolver) ExitedPayload() generated.ExitedPayloadResolver { return &exitedPayloadResolver{r} }
func (r *Resolver) GlobalUser() generated.GlobalUserResolver       { return &globalUserResolver{r} }
func (r *Resolver) Me() generated.MeResolver                       { return &meResolver{r} }
func (r *Resolver) OfflinedPayload() generated.OfflinedPayloadResolver {
	return &offlinedPayloadResolver{r}
}
func (r *Resolver) RoomUser() generated.RoomUserResolver { return &roomUserResolver{r} }
func (r *Resolver) User() generated.UserResolver         { return &userResolver{r} }

type exitedPayloadResolver struct{ *Resolver }
type globalUserResolver struct{ *Resolver }
type meResolver struct{ *Resolver }
type offlinedPayloadResolver struct{ *Resolver }
type roomUserResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
