package repository

type NotFoundErr struct {
	Msg string
}

func (e *NotFoundErr) Error() string {
	return e.Msg
}

func NewNotFoundErr(msg string) *NotFoundErr {
	return &NotFoundErr{msg}
}
