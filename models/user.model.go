package models

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name" form:"first_name" `
	LastName  string `json:"last_name" form:"last_name" `
	Email     string `gorm:"unique ; not null " json:"email" form:"email" `
	Password  string `json:"password" form:"password" gorm:"not null"`
}

type UserRepository interface {
	FindUserByEmail( email string) (*User, error)
	FindAllUsers() ([]User, error)
	UpdateUser( u *User) (*User, error)
	DeleteUserById( id int) error
}

type UserService interface {
	FindUserByEmail(c context.Context, email string) (*User, error)
	FindAllUsers(c context.Context) ([]User, error)
	UpdateUser(c context.Context, u *User) (*User, error)
	DeleteUserById(c context.Context, id int) error
}
