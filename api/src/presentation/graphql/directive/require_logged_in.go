package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	graph "github.com/laster18/poi/api/src/presentation/graphql"
	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/aerrors"
)

func (d *Directive) RequireLoggedIn(
	ctx context.Context,
	obj interface{},
	next graphql.Resolver,
) (res interface{}, err error) {
	currentUser := acontext.GetUser(ctx)
	if currentUser == nil {
		graph.HandleErr(ctx, aerrors.Wrap(graph.ErrUnauthorized))
		return nil, nil
	}

	return next(ctx)
}
