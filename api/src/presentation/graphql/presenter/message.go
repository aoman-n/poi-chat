package presenter

import (
	"strconv"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/presentation/graphql"
)

func ToMessage(m *domain.Message) *model.Message {
	if m == nil {
		return nil
	}

	return &model.Message{
		ID:            strconv.Itoa(m.ID),
		UserID:        m.UserUID,
		UserName:      m.UserName,
		UserAvatarURL: m.UserAvatarURL,
		Body:          m.Body,
		CreatedAt:     m.CreatedAt,
	}
}

func ToSendMessagePayload(m *domain.Message) *model.SendMassagePaylaod {
	return &model.SendMassagePaylaod{
		Message: ToMessage(m),
	}
}

func ToMessageConnection(before *string, resp *domain.MessageListResp, totalCount int) *model.MessageConnection {
	// create pageInfo
	hasNextPage := false
	if before != nil {
		hasNextPage = true
	}
	startCursor, endCursor := graphql.MessageCursors(resp.List)

	pageInfo := &model.PageInfo{
		StartCursor:     startCursor,
		EndCursor:       endCursor,
		HasNextPage:     hasNextPage,
		HasPreviousPage: resp.HasPreviousPage,
	}

	// create nodes and edges
	nodes := make([]*model.Message, len(resp.List))
	edges := make([]*model.MessageEdge, len(resp.List))

	for i, message := range resp.List {
		nodes[i] = ToMessage(message)
		edges[i] = &model.MessageEdge{
			Cursor: *graphql.MessageCursor(message.GetID(), message.GetCreatedAtUnix()),
			Node:   nodes[i],
		}
	}

	return &model.MessageConnection{
		PageInfo:     pageInfo,
		Nodes:        nodes,
		Edges:        edges,
		MessageCount: totalCount,
	}
}
