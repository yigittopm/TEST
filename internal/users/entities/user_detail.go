package entities

import (
	"github.com/yigittopm/wl-auth/pkg/model"
)

type UserDetail struct {
	model.Base
	UserID    uint   `json:"userID" gorm:"primaryKey"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Address   string `json:"address"`
}
