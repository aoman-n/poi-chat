package aerrors

type Code string

const (
	CodeNotFound     Code = "not_found"
	CodeBadParams    Code = "bad_params"
	CodeUnauthorized Code = "unauthorized"
	CodeDuplicated   Code = "duplicated"

	CodeDatabase Code = "database_error"
	CodeRedis    Code = "redis_error"
	CodeInternal Code = "internal_error"

	CodeUnknown Code = "unknown"
)
