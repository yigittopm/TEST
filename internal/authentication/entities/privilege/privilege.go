package privilege

import (
	dtos "github.com/yigittopm/wl-auth/internal/authentication/dtos/privilege"
	"github.com/yigittopm/wl-auth/pkg/model"
)

type Privilege struct {
	model.Base
	Description string `json:"description" gorm:"unique;not null"`
	Key         string `json:"key" gorm:"unique;not null"`
}

func New(data dtos.CreatePrivilegeRequest) Privilege {
	return Privilege{
		Description: data.Description,
		Key:         data.Key,
	}
}
