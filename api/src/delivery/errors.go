package delivery

import "errors"

var (
	errUnauthenticated = errors.New("Unauthenticated")
	errUnknown         = errors.New("Something went wrong")
)
