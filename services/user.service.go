package services

import (
	c "github.com/gmshuvo/go-gin-postgres/config"
	m "github.com/gmshuvo/go-gin-postgres/models"
	repo "github.com/gmshuvo/go-gin-postgres/repositories"
)



// Create user
// func CreateUser(u *m.User) (*m.User, error) {
// 	user, err := repo.CreateUser(u)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil

// }

// Find user by email
func FindUserByEmail(email string) (*m.User, error) {
	user, err := repo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Find user by id
func FindUserById(id int) (*m.User, error) {
	user, err := repo.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Find all users
func FindAllUsers() ([]m.User, error) {
	var users []m.User
	err := c.GetDB().Find(&users).Error
	return users, err
}

// Update user
func UpdateUser(u *m.User) (*m.User, error) {
	
	err := c.GetDB().Model(&u).Updates(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Delete user by id
func DeleteUserById(id int) error {
	err := c.GetDB().Where("id = ?", id).Delete(&m.User{}).Error
	return err
}

// Delete all users
func DeleteAllUsers() error {
	err := c.GetDB().Delete(&m.User{}).Error
	return err
}

// called db migration
