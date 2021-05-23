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
		return nil, aerrors.Wrap(err, "failed to roomRepo.GetByName")
	}
	if dupRoom != nil {
		msg := fmt.Sprintf("%q is already exists", input.Name)
		return nil, aerrors.New(msg).Message(msg)
	}

	newRoom := domain.NewRoom(input.Name, "#20b2aa")
	if err := newRoom.Validate(); err != nil {
		return nil, aerrors.Wrap(err)
	}

	if err := r.roomRepo.Create(ctx, newRoom); err != nil {
		return nil, aerrors.Wrap(err, "failed to roomRepo.Create")
	}

	return toCreateRoomPayload(newRoom), nil
}

func (r *queryResolver) Rooms(
	ctx context.Context,
	first *int, after *string,
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
			return nil, aerrors.Wrap(err)
		}

		roomListReq.LastKnownID = id
		roomListReq.LastKnownUnix = unix
	}

	roomListResp, err := r.roomRepo.List(ctx, roomListReq)
	if err != nil {
		return nil, aerrors.Wrap(err, "failed to roomRepo.List")
	}
	roomCount, err := r.roomRepo.Count(ctx)
	if err != nil {
		return nil, aerrors.Wrap(err, "failed to roomRepo.Count")
	}

	// create pageInfo
	var pageInfo model.PageInfo
	pageInfo.HasNextPage = roomListResp.HasNext
	if after != nil {
		pageInfo.HasPreviousPage = true
	} else {
		pageInfo.HasPreviousPage = false
	}
	startCursor, endCursor := getRoomCursors(roomListResp.List)
	pageInfo.StartCursor = startCursor
	pageInfo.EndCursor = endCursor

	// serialize room model
	nodes := make([]*model.Room, len(roomListResp.List))
	for i, room := range roomListResp.List {
		nodes[i] = &model.Room{
			ID:   strconv.Itoa(int(room.ID)),
			Name: room.Name,
		}
	}

	// create connection
	edges := make([]*model.RoomEdge, len(roomListResp.List))
	for i, room := range roomListResp.List {
		edges[i] = &model.RoomEdge{
			Cursor: *encodeCursor(roomPrefix, room.ID, int(room.CreatedAt.Unix())),
			Node:   nodes[i],
		}
	}

	return &model.RoomConnection{
		PageInfo:  &pageInfo,
		Nodes:     nodes,
		Edges:     edges,
		RoomCount: roomCount,
	}, nil
}

func (r *queryResolver) Room(ctx context.Context, id string) (*model.Room, error) {
	roomID, err := decodeID(roomPrefix, id)
	if err != nil {
		return nil, aerrors.Wrap(err)
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
		return 0, err
	}

	count, err := acontext.GetRoomUserCountLoader(ctx).Load(id)
	if err != nil {
		return 0, aerrors.Wrap(err, "failed to roomUserCountLoader.Load")
	}

	return count, nil
}

// Room returns generated.RoomResolver implementation.
func (r *Resolver) Room() generated.RoomResolver { return &roomResolver{r} }

type roomResolver struct{ *Resolver }
