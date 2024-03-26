package dtos

import "github.com/invopop/validation"

type UpdateRoleRequest struct {
	Name       string   `json:"name"`
	Privileges []string `json:"privileges"`
}

type UpdateRoleResponse struct {
	ID         uint     `json:"id"`
	Name       string   `json:"name"`
	Privileges []string `json:"privileges"`
}

func (req UpdateRoleRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Privileges, validation.Required, validation.Length(1, 10)))
}
