package entities

import "gorm.io/gorm"

type UserDetail struct {
	gorm.Model
	UserID    uint `gorm:"primaryKey"`
	Firstname string
	Lastname  string
	Phone     string
	Country   string
	City      string
	Address   string
}
