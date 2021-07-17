package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/src/registry"
)

type Func func(
	ctx context.Context,
	obj interface{},
	next graphql.Resolver,
) (res interface{}, err error)

type Directive struct {
	repo registry.Repository
}

func New(repo registry.Repository) *generated.DirectiveRoot {
	d := Directive{repo}

	return &generated.DirectiveRoot{
		RequireEntered:  d.RequireEntered,
		RequireLoggedIn: d.RequireLoggedIn,
	}
}
