package dtos

import "github.com/invopop/validation"

type UpdateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserResponse struct {
	ID string `json:"id"`
}

func (req UpdateUserRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.Required, validation.Length(3, 255)),
		validation.Field(&req.Email, validation.Required, validation.Length(8, 255)),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 255)),
	)
}
