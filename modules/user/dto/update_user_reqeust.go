package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type UpdateUserRequest struct {
	ID    uuid.UUID
	Name  string
	Email string
}

func (u UpdateUserRequest) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.ID, validation.Required, is.UUIDv4),
		validation.Field(&u.Name, validation.Required, validation.Length(2, 255)),
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}
