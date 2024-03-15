package services

import (
	"context"
	"time"

	"github.com/gmshuvo/go-gin-postgres/models"
	"github.com/gmshuvo/go-gin-postgres/utils"
)

type authService struct {
	authRepository models.AuthRepository
	timeout        time.Duration
}

func NewAuthService(ar models.AuthRepository, timeout time.Duration) *authService {
	return &authService{
		authRepository: ar,
		timeout:        timeout,
	}
}

// Register user
func (as *authService) Register(c context.Context, u *models.User) (*models.User, error) {
	_, cancel := context.WithTimeout(c, as.timeout)
	defer cancel()
	// hash password
	u.Password = utils.HashPassword(u.Password)
	// create user
	user, err := as.authRepository.Register(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Login user
func (as *authService) Login(c context.Context, u *models.User) (*models.User, error) {
	_, cancel := context.WithTimeout(c, as.timeout)
	defer cancel()
	user, err := as.authRepository.Login(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Logout user
func (as *authService) Logout(c context.Context, u *models.User) (*models.User, error) {
	_, cancel := context.WithTimeout(c, as.timeout)
	defer cancel()
	user, err := as.authRepository.Logout(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}
