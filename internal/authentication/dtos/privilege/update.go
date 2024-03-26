package privilege

import "github.com/invopop/validation"

type UpdatePrivilegeRequest struct {
	Description string `json:"description"`
	Key         string `json:"key"`
}

type UpdatePrivilegeResponse struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Key         string `json:"key"`
}

func (req UpdatePrivilegeRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Description, validation.Required, validation.Length(3, 255)),
		validation.Field(&req.Key, validation.Required, validation.Length(3, 255)),
	)
}
