package xerror

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrInvalidRequest      = errors.New("invalid request")
	ErrInternalError       = errors.New("internal server error")
	ErrAlreadyExists       = errors.New("already exists")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrUnauthenticated     = errors.New("unauthenticated request")
)
