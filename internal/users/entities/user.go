package entities

import (
	"github.com/yigittopm/wl-auth/internal/users/dtos"
	"github.com/yigittopm/wl-auth/pkg/constant"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"index;unique;not null"`
	Password string `gorm:"not null"`
	UserType string
	IsActive bool
}

func New(data dtos.RegisterRequest) User {
	return User{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		UserType: constant.USER_TYPE,
		IsActive: true,
	}
}
