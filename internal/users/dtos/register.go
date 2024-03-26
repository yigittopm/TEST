package dtos

import (
	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID uint `json:"id"`
}

func (req RegisterRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, validation.Length(8, 255), is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 255)),
	)
}
