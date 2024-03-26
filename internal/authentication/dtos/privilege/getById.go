package privilege

import "github.com/invopop/validation"

type GetPrivilegeByIdRequest struct {
	ID uint `json:"id"`
}

type GetPrivilegeByIdResponse struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Key         string `json:"key"`
}

func (req GetPrivilegeByIdRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.ID, validation.Required),
	)
}
