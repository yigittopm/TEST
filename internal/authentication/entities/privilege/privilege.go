package privilege

import (
	dtos "github.com/yigittopm/wl-auth/internal/authentication/dtos/privilege"
	"gorm.io/gorm"
)

type Privilege struct {
	gorm.Model
	Description string `gorm:"unique;not null"`
	Key         string `gorm:"unique;not null"`
}

func New(data dtos.CreatePrivilegeRequest) Privilege {
	return Privilege{
		Description: data.Description,
		Key:         data.Key,
	}
}
