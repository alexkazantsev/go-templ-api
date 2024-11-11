package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type FindManyRequest struct {
	Name   string `form:"name"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

func (f FindManyRequest) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Name),
		validation.Field(&f.Limit, validation.Required, validation.Min(1), validation.Max(100)),
		validation.Field(&f.Offset, validation.Min(0)),
	)
}
