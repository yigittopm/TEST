package dtos

import "github.com/invopop/validation"

type DeleteRoleRequest struct {
	ID uint `json:"id"`
}

type DeleteRoleResponse struct {
	ID uint `json:"id"`
}

func (req DeleteRoleRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.ID, validation.Required))
}
