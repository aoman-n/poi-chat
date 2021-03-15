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
)

func (r *mutationResolver) CreateRoom(ctx context.Context, input *model.CreateRoomInput) (*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Rooms(ctx context.Context, first *int, after *string, orderBy *model.RoomOrderField) (*model.RoomConnection, error) {
	roomListReq := &domain.RoomListReq{}

	if first != nil {
		roomListReq.Limit = *first
	}
	if after != nil {
		// afterCursorのformatチェック
		id, unix, err := decodeCursor(roomPrefix, after)
		if err != nil {
			return nil, err
		}

		roomListReq.LastKnownID = id
		roomListReq.LastKnownUnix = unix
	}

	roomListResp, err := r.roomRepo.List(ctx, roomListReq)
	if err != nil {
		return nil, err
	}
	roomCount, err := r.roomRepo.Count(ctx)
	if err != nil {
		return nil, err
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
			Cursor: *encodeCursor("Room:", room.ID, int(room.CreatedAt.Unix())),
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

func (r *queryResolver) RoomDetail(ctx context.Context, id string) (*model.RoomDetail, error) {
	roomID, err := decodeID(roomPrefix, id)
	if err != nil {
		return nil, err
	}

	room, err := r.roomRepo.GetByID(ctx, roomID)
	if err != nil {
		return nil, err
	}

	return &model.RoomDetail{
		ID:   strconv.Itoa(room.ID),
		Name: room.Name,
	}, nil
}

func (r *roomResolver) ID(ctx context.Context, obj *model.Room) (string, error) {
	return fmt.Sprintf(roomIDFormat, obj.ID), nil
}

func (r *roomDetailResolver) ID(ctx context.Context, obj *model.RoomDetail) (string, error) {
	return fmt.Sprintf(roomIDFormat, obj.ID), nil
}

func (r *subscriptionResolver) SubMessage(ctx context.Context, roomID string) (<-chan *model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

// Room returns generated.RoomResolver implementation.
func (r *Resolver) Room() generated.RoomResolver { return &roomResolver{r} }

// RoomDetail returns generated.RoomDetailResolver implementation.
func (r *Resolver) RoomDetail() generated.RoomDetailResolver { return &roomDetailResolver{r} }

type roomResolver struct{ *Resolver }
type roomDetailResolver struct{ *Resolver }
