package entities

import (
	"time"

	"github.com/yigittopm/test/internal/users/dtos"
	"github.com/yigittopm/test/pkg/constant"
)

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	UserType  string
	IsActive  bool
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

func New(data dtos.CreateUserRequest) User {
	return User{
		Username:  data.Username,
		Email:     data.Email,
		Password:  data.Password, // TODO: encrypt this field
		UserType:  constant.USER_TYPE,
		IsActive:  true,
		CreatedAt: time.Now(),
		CreatedBy: "SYSTEM",
		UpdatedAt: time.Now(),
		UpdatedBy: "SYSTEM",
	}
}
