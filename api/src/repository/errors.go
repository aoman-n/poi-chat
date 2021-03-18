package repository

import "errors"

var (
	errZeroID = errors.New("id is zero value")
)

type NotFoundErr struct {
	Msg string
}

func (e *NotFoundErr) Error() string {
	return e.Msg
}

func NewNotFoundErr(msg string) *NotFoundErr {
	return &NotFoundErr{msg}
}
