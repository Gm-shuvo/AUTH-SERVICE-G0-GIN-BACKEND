package services

import (
	"context"
	"time"

	// "github.com/gmshuvo/go-gin-postgres/config"
	"github.com/gmshuvo/go-gin-postgres/models"
	// "github.com/gmshuvo/go-gin-postgres/repositories"
)

type userService struct {
	userRepository models.UserRepository
	timeout        time.Duration
}

func NewUserService(ur models.UserRepository, timeout time.Duration) *userService {
	return &userService{
		userRepository: ur,
		timeout:        timeout,
	}
}

// Find user by email
func (us *userService) FindUserByEmail(c context.Context, email string) (*models.User, error) {
	user, err := us.userRepository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Find user by id
// func (us *userService) FindUserById(id int) (*models.User, error) {
// 	user, err := us.userRepository.FindUserById(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// Find all users
func (us *userService) FindAllUsers(c context.Context) ([]models.User, error) {
	_, cancel := context.WithTimeout(c, us.timeout)
	defer cancel()


	users, err := us.userRepository.FindAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Update user
func (us *userService) UpdateUser(c context.Context, u *models.User) (*models.User, error) {
	_, cancel := context.WithTimeout(c, us.timeout)
	defer cancel()
	user, err := us.userRepository.UpdateUser(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Delete user by id
func (us *userService) DeleteUserById(c context.Context, id int) error {
	_, cancel := context.WithTimeout(c, us.timeout)
	defer cancel()
	err := us.userRepository.DeleteUserById(id)
	return err
}
