package dtos

import "github.com/invopop/validation"

type GetRoleByIdRequest struct {
	ID uint `json:"id"`
}

type GetdRoleByIdResponse struct {
	ID         uint     `json:"id"`
	Name       string   `json:"name"`
	Privileges []string `json:"privileges"`
}

func (req GetRoleByIdRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.ID, validation.Required))
}
