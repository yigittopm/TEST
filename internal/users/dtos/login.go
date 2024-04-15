package dtos

import (
	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
	"github.com/yigittopm/wl-auth/internal/authentication/entities/role"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID          uint        `json:"id"`
	Roles       []role.Role `json:"roles"`
	AccessToken string      `json:"accessToken"`
}

func (req LoginRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, validation.Length(3, 255), is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 255)),
	)
}
