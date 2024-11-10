package xerror

import (
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func MapErrorToCode(err error) int {
	var (
		code int
		e    validation.Errors
	)

	switch {
	case errors.As(err, &e) || errors.Is(err, ErrInvalidRequest):
		code = http.StatusBadRequest
	case errors.Is(err, ErrNotFound):
		code = http.StatusNotFound
	case errors.Is(err, ErrAlreadyExists):
		code = http.StatusConflict
	case errors.Is(err, ErrUnprocessableEntity):
		code = http.StatusUnprocessableEntity
	case errors.Is(err, ErrUnauthenticated):
		code = http.StatusUnauthorized
	default:
		code = http.StatusInternalServerError
	}

	return code
}
