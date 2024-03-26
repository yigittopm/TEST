package dtos

import "github.com/invopop/validation"

type ProfileRequest struct {
	ID uint `json:"id"`
}

type ProfileResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func (req ProfileRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.ID, validation.Required),
	)
}
