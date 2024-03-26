package role

import (
	dtos "github.com/yigittopm/wl-auth/internal/authentication/dtos/role"
	"github.com/yigittopm/wl-auth/internal/authentication/entities/privilege"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name       string                `gorm:"unique;not null"`
	Privileges []privilege.Privilege `gorm:"many2many:role_privileges;"`
}

func New(data dtos.CreateRoleRequest) Role {
	return Role{
		Name:       data.Name,
		Privileges: data.Privileges,
	}
}
