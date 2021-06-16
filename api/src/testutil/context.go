package testutil

import "context"

func NewMockCtx() context.Context {
	ctx := context.Background()
	// TODO: 必要なものを注入する
	return ctx
}
