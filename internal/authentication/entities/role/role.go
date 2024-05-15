package role

import (
	dtos "github.com/yigittopm/wl-auth/internal/authentication/dtos/role"
	"github.com/yigittopm/wl-auth/internal/authentication/entities/privilege"
	"github.com/yigittopm/wl-auth/pkg/model"
)

type Role struct {
	model.Base
	Name       string                `json:"name" gorm:"unique;not null"`
	Privileges []privilege.Privilege `json:"privileges" gorm:"many2many:role_privileges;"`
}

func New(data dtos.CreateRoleRequest) Role {
	return Role{
		Name:       data.Name,
		Privileges: data.Privileges,
	}
}
