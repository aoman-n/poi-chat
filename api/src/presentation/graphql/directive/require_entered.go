package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (d *Directive) RequireEntered(
	ctx context.Context,
	obj interface{},
	next graphql.Resolver,
) (res interface{}, err error) {
	return nil, nil
}
