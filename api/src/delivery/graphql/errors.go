package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/iancoleman/strcase"
	"github.com/laster18/poi/api/src/util/aerrors"
	"github.com/laster18/poi/api/src/util/validator"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// var (
// 	errBadCredentials  = errors.New("Email/password combination don't work")
// 	errUnauthenticated = errors.New("Unauthenticated")
// 	errUnknown         = errors.New("Something went wrong")
// 	errRoomNotFound    = errors.New("Not found room")
// 	errUnexpected      = errors.New("Internal server error")
// )

var (
	invalidIDMsg = "invalid id format: %s"
)

type GraphErrCode string

const (
	CodeUnauthorized GraphErrCode = "AUTHENTICATION_ERROR"
	CodeInternal     GraphErrCode = "INTERNAL_SERVER_ERROR"
	CodeUserInput    GraphErrCode = "USER_INPUT_ERROR"
	CodeNotFound     GraphErrCode = "NOT_FOUND_ERROR"
)

const (
	UnauthoziedErrMsg = "unauthorized"
	InternalErrMsg    = "internal server error"
)

func addErr(ctx context.Context, message string, code GraphErrCode) {
	graphql.AddError(ctx, &gqlerror.Error{
		Message:    message,
		Path:       graphql.GetPath(ctx),
		Extensions: map[string]interface{}{"code": code},
	})
}

func addValidationErr(ctx context.Context, vErr *validator.ErrValidation) {
	for fieldName, errString := range vErr.GetErrFields() {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: errString,
			Path:    graphql.GetPath(ctx),
			Extensions: map[string]interface{}{
				"code":      CodeUserInput,
				"attribute": strcase.ToLowerCamel(fieldName),
			},
		})
	}
}

func handleErr(ctx context.Context, err error) {
	e := aerrors.AsErrApp(err)
	if e == nil {
		// TODO: unexpected error handling
		return
	}

	switch e.Code() {
	case aerrors.CodeNotFound:
		addErr(ctx, getInfoMsg(e), CodeNotFound)
	case aerrors.CodeUnauthorized:
		addErr(ctx, getInfoMsg(e), CodeUnauthorized)
	case aerrors.CodeBadParams:
		addErr(ctx, getInfoMsg(e), CodeUserInput)
	case aerrors.CodeDatabase, aerrors.CodeRedis, aerrors.CodeInternal:
		addErr(ctx, getInfoMsg(e), CodeInternal)
	case aerrors.CodeUnknown:
		fallthrough
	default:
		addErr(ctx, getInfoMsg(e), CodeInternal)
	}
}

var outputMsgMap = map[aerrors.Code]string{
	aerrors.CodeNotFound:     "",
	aerrors.CodeBadParams:    "",
	aerrors.CodeUnauthorized: "",
	aerrors.CodeDatabase:     "",
	aerrors.CodeRedis:        "",
	aerrors.CodeInternal:     "",
	aerrors.CodeUnknown:      "",
}

func getInfoMsg(e *aerrors.ErrApp) string {
	if e.InfoMsg() != "" {
		return e.InfoMsg()
	}
	if msg, ok := outputMsgMap[e.Code()]; ok {
		return msg
	}

	return "internal server error"
}
