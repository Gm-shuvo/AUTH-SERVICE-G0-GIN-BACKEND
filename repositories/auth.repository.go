package repositories

import (
	c "github.com/gmshuvo/go-gin-postgres/config"
	m "github.com/gmshuvo/go-gin-postgres/models"
)

// Register Repository
func Register(u *m.User) (*m.User, error) {
	err := c.GetDB().Model(&u).Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Login Repository
func Login(u *m.User) (*m.User, error) {
	var user m.User
	err := c.GetDB().Where("email = ? AND password = ?", u.Email, u.Password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Logout Repository
func Logout(u *m.User) (*m.User, error) {
	return u, nil
}
