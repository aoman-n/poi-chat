package graphql

import "errors"

var (
	errBadCredentials  = errors.New("Email/password combination don't work")
	errUnauthenticated = errors.New("Unauthenticated")
	errUnknown         = errors.New("Something went wrong")
	errRoomNotFound    = errors.New("Not found room")
	errUnexpected      = errors.New("Internal server error")
)
