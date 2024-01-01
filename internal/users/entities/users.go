package entities

import (
	"time"

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

func New() User {
	return User{
		Username:  "",
		Email:     "",
		Password:  "",
		UserType:  constant.USER_TYPE,
		IsActive:  true,
		CreatedAt: time.Now(),
		CreatedBy: "SYSTEM",
	}
}
