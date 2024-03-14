package services

import (
	m "github.com/gmshuvo/go-gin-postgres/models"
	repo "github.com/gmshuvo/go-gin-postgres/repositories"
	"github.com/gmshuvo/go-gin-postgres/utils"
)

// Register user
func Register(u *m.User) (*m.User, error) {
	// hash password
	u.Password = utils.HashPassword(u.Password)
	// create user
	user, err := repo.Register(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Login user
func Login(u *m.User) (*m.User, error) {
	user, err := repo.Login(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

