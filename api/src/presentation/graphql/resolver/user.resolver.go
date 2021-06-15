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
	"github.com/laster18/poi/api/src/domain"
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

	roomUser, err := r.roomUserRepo.Get(ctx, domainRoomID, currentUser.UID)
	if err != nil {
		log.Println("failed to get roomUser err:", err)
	}
	if roomUser == nil {
		roomUser = domain.NewDefaultRoomUser(domainRoomID, currentUser)
	}
	roomUser.SetPosition(input.X, input.Y)

	if err := r.roomUserRepo.Save(ctx, roomUser); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomUserRepo.Save"))
		return nil, nil
	}

	return presenter.ToMovePayload(roomUser), nil
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
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		graphql.HandleErr(ctx, aerrors.Wrap(graphql.ErrUnauthorized))
		return nil, nil
	}

	users, err := r.globalUserRepo.GetAll(ctx)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	return presenter.ToGlobalUsers(users), nil
}

func (r *roomResolver) Users(ctx context.Context, obj *model.Room) ([]*model.RoomUser, error) {
	roomID, _ := strconv.Atoi(obj.ID)

	users, err := r.roomUserRepo.GetByRoomID(ctx, roomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomUserRepo.GetByRoomID"))
		return nil, nil
	}

	return presenter.ToRoomUsers(users), nil
}

func (r *roomUserResolver) ID(ctx context.Context, obj *model.RoomUser) (string, error) {
	return graphql.RoomUserIDStr(obj.ID), nil
}

func (r *subscriptionResolver) ActedGlobalUserEvent(ctx context.Context) (<-chan model.GlobalUserEvent, error) {
	currentUser := acontext.GetUser(ctx)

	if currentUser == nil {
		graphql.HandleErr(ctx, aerrors.Wrap(graphql.ErrUnauthorized))
		return nil, nil
	}

	ch := make(chan model.GlobalUserEvent)
	r.globalUserSubscriber.AddCh(ch, currentUser.UID)

	newGlobalUser := &domain.GlobalUser{
		UID:       currentUser.UID,
		Name:      currentUser.Name,
		AvatarURL: currentUser.AvatarURL,
	}
	if err := r.globalUserRepo.Save(ctx, newGlobalUser); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	go func() {
		<-ctx.Done()
		r.globalUserSubscriber.RemoveCh(currentUser.UID)
		if err := r.globalUserRepo.Delete(context.Background(), currentUser.UID); err != nil {
			log.Println("failed to delete globalUser, err:", err)
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
	r.roomUserSubscriber.AddCh(ch, domainRoomID, currentUser.UID)

	newRoomUser := domain.NewDefaultRoomUser(domainRoomID, currentUser)
	if err := r.roomUserRepo.Save(ctx, newRoomUser); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomUserRepo.Save"))
		return nil, nil
	}

	go func() {
		<-ctx.Done()
		r.roomUserSubscriber.RemoveCh(domainRoomID, currentUser.UID)
		if err := r.roomUserRepo.Delete(context.Background(), newRoomUser); err != nil {
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

	roomUser, err := r.roomUserRepo.Get(ctx, domainRoomID, currentUser.UID)
	if err != nil {
		logger.Infof("failed to get roomUser err: %v", err)
	}
	if roomUser == nil {
		roomUser = domain.NewDefaultRoomUser(domainRoomID, currentUser)
	}
	roomUser.LastMessage = nil
	roomUser.LastEvent = domain.RemoveLastMessageEvent

	if err := r.roomUserRepo.Save(ctx, roomUser); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomUserRepo.Save"))
		return nil, nil
	}

	return presenter.ToRemoveLastMessagePayload(roomUser), nil
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

	roomUser, err := r.roomUserRepo.Get(ctx, domainRoomID, currentUser.UID)
	if err != nil {
		logger.Infof("failed to get roomUser err: %v", err)
	}
	if roomUser == nil {
		roomUser = domain.NewDefaultRoomUser(domainRoomID, currentUser)
	}
	roomUser.LastEvent = domain.ChangeBalloonPositionEvent

	var domainBalloonPos domain.BalloonPosition
	switch input.BalloonPosition {
	case model.BalloonPositionTopRight:
		domainBalloonPos = domain.TopRight
	case model.BalloonPositionTopLeft:
		domainBalloonPos = domain.TopLeft
	case model.BalloonPositionBottomRight:
		domainBalloonPos = domain.BottomRight
	case model.BalloonPositionBottomLeft:
		domainBalloonPos = domain.BottomLeft
	default:
		m := fmt.Sprintf("%s is unknown balloon position", string(input.BalloonPosition))
		graphql.HandleErr(ctx, aerrors.New(m).SetCode(aerrors.CodeBadParams).Message(m))
	}

	roomUser.BalloonPosition = domainBalloonPos

	if err := r.roomUserRepo.Save(ctx, roomUser); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomUserRepo.Save"))
		return nil, nil
	}

	return presenter.ToChangeBalloonPositionPayload(roomUser), nil
}

// ExitedPayload returns generated.ExitedPayloadResolver implementation.
func (r *Resolver) ExitedPayload() generated.ExitedPayloadResolver { return &exitedPayloadResolver{r} }

// GlobalUser returns generated.GlobalUserResolver implementation.
func (r *Resolver) GlobalUser() generated.GlobalUserResolver { return &globalUserResolver{r} }

// Me returns generated.MeResolver implementation.
func (r *Resolver) Me() generated.MeResolver { return &meResolver{r} }

// OfflinedPayload returns generated.OfflinedPayloadResolver implementation.
func (r *Resolver) OfflinedPayload() generated.OfflinedPayloadResolver {
	return &offlinedPayloadResolver{r}
}

// RoomUser returns generated.RoomUserResolver implementation.
func (r *Resolver) RoomUser() generated.RoomUserResolver { return &roomUserResolver{r} }

type exitedPayloadResolver struct{ *Resolver }
type globalUserResolver struct{ *Resolver }
type meResolver struct{ *Resolver }
type offlinedPayloadResolver struct{ *Resolver }
type roomUserResolver struct{ *Resolver }
