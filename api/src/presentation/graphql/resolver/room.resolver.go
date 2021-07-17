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
	"github.com/laster18/poi/api/src/presentation/graphql"
	"github.com/laster18/poi/api/src/presentation/graphql/presenter"
	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/aerrors"
)

func (r *mutationResolver) CreateRoom(
	ctx context.Context,
	input *model.CreateRoomInput,
) (*model.CreateRoomPayload, error) {
	currentUser := acontext.GetUser(ctx)
	roomRepo := r.repo.NewRoom()
	roomService := r.service.NewRoom()

	exists, err := roomService.ExistsRoom(ctx, input.Name)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomService.ExistsRoom"))
		return nil, nil
	}
	if exists {
		msg := fmt.Sprintf("%q is already exists", input.Name)
		graphql.HandleErr(ctx, aerrors.New(msg).Message("already exists room name"))
		return nil, nil
	}

	newRoom := room.New(currentUser.ID, input.Name, input.BgColor, input.BgURL)
	if err := newRoom.Validate(); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	if err := roomRepo.Create(ctx, newRoom); err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.Create"))
		return nil, nil
	}

	return presenter.ToCreateRoomPayload(newRoom), nil
}

func (r *queryResolver) Rooms(
	ctx context.Context,
	first *int,
	after *string,
	orderBy *model.RoomOrderField,
) (*model.RoomConnection, error) {
	roomListReq := &room.ListReq{}

	if first != nil {
		roomListReq.Limit = *first
	}
	if after != nil {
		// afterCursorのformatチェック
		id, unix, err := graphql.DecodeRoomCursor(after)
		if err != nil {
			graphql.HandleErr(ctx, aerrors.Wrap(err))
			return nil, nil
		}

		roomListReq.LastKnownID = id
		roomListReq.LastKnownUnix = unix
	}

	roomRepo := r.repo.NewRoom()
	roomListResp, err := roomRepo.List(ctx, roomListReq)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.List"))
		return nil, nil
	}
	totalCount, err := roomRepo.Count(ctx)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.Count"))
		return nil, nil
	}

	return presenter.ToRoomConnection(after, roomListResp, totalCount), nil
}

func (r *queryResolver) Room(ctx context.Context, id string) (*model.Room, error) {
	roomID, err := graphql.DecodeRoomID(id)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	roomRepo := r.repo.NewRoom()
	room, err := roomRepo.GetByID(ctx, roomID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.GetByID"))
		return nil, nil
	}

	return presenter.ToRoom(room), nil
}

func (r *roomResolver) ID(ctx context.Context, obj *model.Room) (string, error) {
	return graphql.RoomIDStr(obj.ID), nil
}

func (r *roomResolver) TotalUserCount(ctx context.Context, obj *model.Room) (int, error) {
	id, err := strconv.Atoi(obj.ID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err))
		return 0, nil
	}

	count, err := acontext.GetRoomUserCountLoader(ctx).Load(id)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomUserCountLoader.Load"))
		return 0, nil
	}

	return count, nil
}

func (r *roomResolver) TotalMessageCount(ctx context.Context, obj *model.Room) (int, error) {
	id, err := strconv.Atoi(obj.ID)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err))
		return 0, nil
	}

	count, err := acontext.GetRoomMessageCountLoader(ctx).Load(id)
	if err != nil {
		graphql.HandleErr(ctx, aerrors.Wrap(err, "failed to roomMessageCountLoader.Load"))
		return 0, nil
	}

	return count, nil
}

// Room returns generated.RoomResolver implementation.
func (r *Resolver) Room() generated.RoomResolver { return &roomResolver{r} }

type roomResolver struct{ *Resolver }
