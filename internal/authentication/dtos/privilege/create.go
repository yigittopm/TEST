package privilege

import "github.com/invopop/validation"

type CreatePrivilegeRequest struct {
	Description string `json:"description"`
	Key         string `json:"key"`
}

type CreatePrivilegeResponse struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Key         string `json:"key"`
}

func (req CreatePrivilegeRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Description, validation.Required, validation.Length(3, 255)),
		validation.Field(&req.Key, validation.Required, validation.Length(3, 255)),
	)
}
