package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/iancoleman/strcase"
	"github.com/laster18/poi/api/src/util/acontext"
	"github.com/laster18/poi/api/src/util/aerrors"
	"github.com/laster18/poi/api/src/util/validator"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var (
	invalidIDMsg = "invalid id format: %s"
)

// GraphErrCode: GraphQLのerrorレスポンスのextensionsフィールドに入れるエラーコード
type GraphErrCode string

const (
	codeUnauthorized GraphErrCode = "AUTHENTICATION_ERROR"
	codeInternal     GraphErrCode = "INTERNAL_SERVER_ERROR"
	codeUserInput    GraphErrCode = "USER_INPUT_ERROR"
	codeNotFound     GraphErrCode = "NOT_FOUND_ERROR"
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
				"code":      codeUserInput,
				"attribute": strcase.ToLowerCamel(fieldName),
			},
		})
	}
}

// GraphErrCode: ErrAppのinfoMsgが存在しなかった場合にユーザーへ返すエラーメッセージ
const (
	notFoundErrMsg    = "not found error"
	badrequestErrMsg  = "bad request error, please check request parameters"
	unauthoziedErrMsg = "unauthorized error"
	internalErrMsg    = "internal server error"
)

var errUnauthorized = aerrors.New("unauthorized").SetCode(aerrors.CodeUnauthorized)

func handleErr(ctx context.Context, err error) {
	logger := acontext.GetLogger(ctx)

	e := aerrors.AsErrApp(err)
	if e == nil {
		// TODO: unexpected error handling
		addErr(ctx, "server errror", codeInternal)
		return
	}

	logger.WarnWithErr(err, "handlerErr")

	switch e.Code() {
	case aerrors.CodeNotFound:
		addErr(ctx, getInfoMsg(e), codeNotFound)
	case aerrors.CodeUnauthorized:
		addErr(ctx, getInfoMsg(e), codeUnauthorized)
	case aerrors.CodeBadParams, aerrors.CodeDuplicated:
		addErr(ctx, getInfoMsg(e), codeUserInput)
	case aerrors.CodeDatabase, aerrors.CodeRedis, aerrors.CodeInternal:
		addErr(ctx, getInfoMsg(e), codeInternal)
	case aerrors.CodeUnknown:
		fallthrough
	default:
		addErr(ctx, getInfoMsg(e), codeInternal)
	}

	return
}

var errInfoMsgMap = map[aerrors.Code]string{
	aerrors.CodeNotFound:     notFoundErrMsg,
	aerrors.CodeBadParams:    badrequestErrMsg,
	aerrors.CodeUnauthorized: unauthoziedErrMsg,
	aerrors.CodeDatabase:     internalErrMsg,
	aerrors.CodeRedis:        internalErrMsg,
	aerrors.CodeInternal:     internalErrMsg,
	aerrors.CodeUnknown:      internalErrMsg,
}

func getInfoMsg(e *aerrors.ErrApp) string {
	if e.InfoMsg() != "" {
		return e.InfoMsg()
	}
	if msg, ok := errInfoMsgMap[e.Code()]; ok {
		return msg
	}

	return internalErrMsg
}
