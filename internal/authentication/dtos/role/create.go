package dtos

import (
	"github.com/invopop/validation"
	"github.com/yigittopm/wl-auth/internal/authentication/entities/privilege"
)

type CreateRoleRequest struct {
	Name       string                `json:"name"`
	Privileges []privilege.Privilege `json:"privileges"`
}

type CreateRoleResponse struct {
	ID         uint                  `json:"id"`
	Name       string                `json:"name"`
	Privileges []privilege.Privilege `json:"privileges"`
}

func (req CreateRoleRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Privileges, validation.Required, validation.Length(1, 10)))
}
