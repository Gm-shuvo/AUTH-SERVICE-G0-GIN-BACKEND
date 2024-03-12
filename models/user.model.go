package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name" form:"first_name" binding:"required" gorm:"not null" `
	LastName  string `json:"last_name" form:"last_name" binding:"required" gorm:"not null"`
	Email     string `gorm:"unique ; not null " json:"email" form:"email" binding:"required" `
	Password  string  `json:"password" form:"password" binding:"required" gorm:"not null"`  
}


