package models

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (u *User) Authenticate(password string) bool {
	return u.Password == password
}

// Create user
func CreateUser(u *User) (*User, error) {
	err := GetDB().Model(&u).Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil

}

// Find user by email
func FindUserByEmail(email string) (*User, error) {
	var user User
	err := GetDB().Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Find user by id
func FindUserById(id int) (*User, error) {
	var user User
	err := GetDB().Where("id = ?", id).First(&user).Error
	return &user, err
}

// Find all users
func FindAllUsers() ([]User, error) {
	var users []User
	err := GetDB().Find(&users).Error
	return users, err
}

// Update user
func UpdateUser(u *User) (*User, error) {
	err := GetDB().Model(&u).Updates(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Delete user by id
func DeleteUserById(id int) error {
	err := GetDB().Where("id = ?", id).Delete(&User{}).Error
	return err
}
