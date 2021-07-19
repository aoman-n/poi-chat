package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	graph "github.com/laster18/poi/api/src/presentation/graphql"
	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/aerrors"
)

func (d *Directive) RequireEntered(
	ctx context.Context,
	obj interface{},
	next graphql.Resolver,
) (res interface{}, err error) {
	currentUser := acontext.GetUser(ctx)
	userRepo := d.repo.NewUser()

	userStatus, err := userRepo.GetStatus(ctx, currentUser.UID)
	if err != nil {
		errApp := aerrors.AsErrApp(err)
		if errApp != nil {
			if errApp.Code() == aerrors.CodeNotFound {
				graph.HandleErr(ctx, aerrors.Wrap(err).SetCode(aerrors.CodeRequireEntered).Message("require entered the room"))
				return nil, nil
			}
		}

		graph.HandleErr(ctx, aerrors.Wrap(err))
		return nil, nil
	}

	if userStatus.EnteredRoomID == nil {
		msg := "require entered the room"
		graph.HandleErr(ctx, aerrors.New(msg).SetCode(aerrors.CodeRequireEntered).Message(msg))
		return nil, nil
	}

	newCtx := acontext.SetUserStatus(ctx, userStatus)

	return next(newCtx)
}
