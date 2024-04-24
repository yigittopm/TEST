package entities

import (
	"github.com/yigittopm/wl-auth/internal/authentication/entities/role"
	"github.com/yigittopm/wl-auth/internal/users/dtos"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string      `gorm:"index;unique;not null"`
	Password string      `gorm:"not null"`
	Detail   UserDetail  `gorm:"foreignKey:UserID"`
	Roles    []role.Role `gorm:"many2many:user_roles;"`
	IsActive bool        `gorm:"default:true"`
}

func New(data dtos.RegisterRequest) User {
	return User{
		Email:    data.Email,
		Password: data.Password,
	}
}
