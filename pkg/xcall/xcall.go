package xcall

import (
	"errors"
	"net/http"

	"github.com/alexkazantsev/go-templ-api/pkg/xerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Response[T any] struct {
	Message string            `json:"message"`
	Payload T                 `json:"payload,omitempty"`
	Error   string            `json:"error,omitempty"`
	Details map[string]string `json:"details,omitempty"`
}

func NewResponse[T any](message string, payload T, err error) *Response[T] {
	var (
		e validation.Errors

		details  = make(map[string]string)
		response = &Response[T]{
			Message: message,
			Payload: payload,
		}
	)

	if err != nil {
		if errors.As(err, &e) {
			for n, d := range e {
				details[n] = d.Error()
			}

			response.Details = details
		}

		response.Error = err.Error()
	}

	return response
}

func ResponseOK[T any](payload T) *Response[T] {
	return NewResponse("OK", payload, nil)
}

func ResponseFail[T any](err error) *Response[T] {
	var p T

	return NewResponse[T]("FAIL", p, err)
}

func CallS(fn func() error) (int, *Response[any]) {
	var err error

	if err = fn(); err == nil {
		return http.StatusOK, ResponseOK[any](nil)
	}

	return xerror.MapErrorToCode(err), ResponseFail[any](err)
}

func CallM[T any](fn func() (T, error)) (int, any) {
	var (
		err error
		p   T
	)

	if p, err = fn(); err == nil {
		return http.StatusOK, ResponseOK(p)
	}

	return xerror.MapErrorToCode(err), ResponseFail[T](err)
}
