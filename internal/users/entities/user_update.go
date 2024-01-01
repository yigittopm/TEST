package entities

import (
	"time"

	"github.com/yigittopm/test/internal/users/dtos"
)

type UpdateUser struct {
	UserID      int64
	Email       string
	Fullname    string
	Password    string
	PhoneNumber string
	UserType    string
	UpdatedAt   time.Time
	UpdatedBy   string
}

func NewUpdateUsers(data dtos.UpdateUserRequest) UpdateUser {
	return UpdateUser{
		UserID:      data.UserID,
		Fullname:    data.Fullname,
		PhoneNumber: data.PhoneNumber,
		UserType:    data.UserType,
		UpdatedAt:   time.Now(),
		UpdatedBy:   data.UpdateBy,
	}
}
