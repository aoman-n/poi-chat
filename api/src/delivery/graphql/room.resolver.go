package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/aerrors"
)

func (r *mutationResolver) CreateRoom(
	ctx context.Context,
	input *model.CreateRoomInput,
) (*model.CreateRoomPayload, error) {
	dupRoom, err := r.roomRepo.GetByName(ctx, input.Name)
	if err != nil {
		handleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.GetByName"))
		return nil, nil
	}
	if dupRoom != nil {
		msg := fmt.Sprintf("%q is already exists", input.Name)
		handleErr(ctx, aerrors.New(msg).Message(msg))
		return nil, nil
	}

	newRoom := domain.NewRoom(input.Name, "#20b2aa")
	if err := newRoom.Validate(); err != nil {
		handleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	if err := r.roomRepo.Create(ctx, newRoom); err != nil {
		handleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.Create"))
		return nil, nil
	}

	return toCreateRoomPayload(newRoom), nil
}

func (r *queryResolver) Rooms(
	ctx context.Context,
	first *int,
	after *string,
	orderBy *model.RoomOrderField,
) (*model.RoomConnection, error) {
	roomListReq := &domain.RoomListReq{}

	if first != nil {
		roomListReq.Limit = *first
	}
	if after != nil {
		// afterCursorのformatチェック
		id, unix, err := decodeCursor(roomPrefix, after)
		if err != nil {
			handleErr(ctx, aerrors.Wrap(err))
			return nil, nil
		}

		roomListReq.LastKnownID = id
		roomListReq.LastKnownUnix = unix
	}

	roomListResp, err := r.roomRepo.List(ctx, roomListReq)
	if err != nil {
		handleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.List"))
		return nil, nil
	}
	totalCount, err := r.roomRepo.Count(ctx)
	if err != nil {
		handleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.Count"))
		return nil, nil
	}

	return toRoomConnection(after, roomListResp, totalCount), nil
}

func (r *queryResolver) Room(ctx context.Context, id string) (*model.Room, error) {
	roomID, err := decodeID(roomPrefix, id)
	if err != nil {
		handleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	room, err := r.roomRepo.GetByID(ctx, roomID)
	if err != nil {
		handleErr(ctx, aerrors.Wrap(err, "failed to roomRepo.GetByID"))
		return nil, nil
	}

	return &model.Room{
		ID:   strconv.Itoa(room.ID),
		Name: room.Name,
	}, nil
}

func (r *roomResolver) ID(ctx context.Context, obj *model.Room) (string, error) {
	return fmt.Sprintf(roomIDFormat, obj.ID), nil
}

func (r *roomResolver) UserCount(ctx context.Context, obj *model.Room) (int, error) {
	id, err := strconv.Atoi(obj.ID)
	if err != nil {
		handleErr(ctx, aerrors.Wrap(err))
		return 0, nil
	}

	count, err := acontext.GetRoomUserCountLoader(ctx).Load(id)
	if err != nil {
		handleErr(ctx, aerrors.Wrap(err, "failed to roomUserCountLoader.Load"))
		return 0, nil
	}

	return count, nil
}

// Room returns generated.RoomResolver implementation.
func (r *Resolver) Room() generated.RoomResolver { return &roomResolver{r} }

type roomResolver struct{ *Resolver }
