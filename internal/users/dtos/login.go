package dtos

import (
	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID          uint   `json:"id"`
	AccessToken string `json:"accessToken"`
}

func (req LoginRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, validation.Length(3, 255), is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 255)),
	)
}
