package entities

import (
	"github.com/yigittopm/wl-auth/internal/authentication/entities/role"
	"github.com/yigittopm/wl-auth/internal/users/dtos"
	"github.com/yigittopm/wl-auth/pkg/model"
)

type User struct {
	model.Base
	Email    string      `json:"email" gorm:"index;unique;not null"`
	Password string      `json:"-" gorm:"not null"`
	Detail   UserDetail  `json:"detail" gorm:"foreignKey:UserID"`
	Roles    []role.Role `json:"roles" gorm:"many2many:user_roles;"`
	IsActive bool        `json:"isActive" gorm:"default:true"`
}

func New(data dtos.RegisterRequest) User {
	return User{
		Email:    data.Email,
		Password: data.Password,
	}
}
