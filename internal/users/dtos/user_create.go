package dtos

import "github.com/invopop/validation"

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID        string `json:"id"`
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}

func (req CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.Required, validation.Length(3, 255)),
		validation.Field(&req.Email, validation.Required, validation.Length(8, 255)),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 255)),
	)
}
