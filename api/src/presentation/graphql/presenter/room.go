package presenter

import (
	"strconv"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/presentation/graphql"
)

func ToCreateRoomPayload(r *domain.Room) *model.CreateRoomPayload {
	return &model.CreateRoomPayload{
		Room: &model.Room{
			ID:        strconv.Itoa(r.ID),
			Name:      r.Name,
			BgColor:   r.BackgroundColor,
			BgURL:     r.BackgroundURL,
			CreatedAt: r.CreatedAt,
		},
	}
}

func ToRoomConnection(after *string, resp *domain.RoomListResp, total int) *model.RoomConnection {
	// create pageInfo
	hasPrevious := false
	if after != nil {
		hasPrevious = true
	}
	startCursor, endCursor := graphql.RoomCursors(resp.List)
	pageInfo := model.PageInfo{
		StartCursor:     startCursor,
		EndCursor:       endCursor,
		HasNextPage:     resp.HasNext,
		HasPreviousPage: hasPrevious,
	}

	// create nodes and edges
	nodes := make([]*model.Room, len(resp.List))
	edges := make([]*model.RoomEdge, len(resp.List))
	for i, room := range resp.List {
		nodes[i] = &model.Room{
			ID:        strconv.Itoa(int(room.ID)),
			Name:      room.Name,
			BgColor:   room.BackgroundColor,
			BgURL:     room.BackgroundURL,
			CreatedAt: room.CreatedAt,
		}
		edges[i] = &model.RoomEdge{
			Cursor: *graphql.RoomCursor(room.ID, int(room.CreatedAt.Unix())),
			Node:   nodes[i],
		}
	}

	return &model.RoomConnection{
		PageInfo:  &pageInfo,
		Nodes:     nodes,
		Edges:     edges,
		RoomCount: total,
	}
}
