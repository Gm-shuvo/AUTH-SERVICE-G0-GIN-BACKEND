package models

import "context"

// Login struct
type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register struct

type Register struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse struct
type LoginResponse struct {
	Login bool `json:"login"`
	Redirect string `json:"redirect"`
	Status int `json:"status"`
}

// RegisterResponse struct
type RegisterResponse struct {
	Registered bool `json:"registered"`
	Status int `json:"status"`
	Redirect string `json:"redirect"`
}

type AuthService interface {
	Register(c context.Context,u *User) (*User, error)
	Login(c context.Context, u *User) (*User, error)
	Logout(c context.Context, u *User) (*User, error)
}

type AuthRepository interface {
	Register(u *User) (*User, error)
	Login(u *User) (*User, error)
	Logout(u *User) (*User, error)
}