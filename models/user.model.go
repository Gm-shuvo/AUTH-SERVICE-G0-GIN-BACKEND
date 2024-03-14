package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name" form:"first_name" `
	LastName  string `json:"last_name" form:"last_name" `
	Email     string `gorm:"unique ; not null " json:"email" form:"email" `
	Password  string `json:"password" form:"password" gorm:"not null"`
}
