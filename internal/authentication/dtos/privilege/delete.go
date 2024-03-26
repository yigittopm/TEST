package privilege

import "github.com/invopop/validation"

type DeletePrivilegeRequest struct {
	ID uint `json:"id"`
}

type DeletePrivilegeResponse struct {
	ID uint `json:"id"`
}

func (req DeletePrivilegeRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.ID, validation.Required),
	)
}
