package repositories

import (
	"github.com/gmshuvo/go-gin-postgres/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB, ) *userRepository {
	return &userRepository{db}
}


// Find user by email
func (ur *userRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}


// func FindUserById(id int) (*m.User, error) {
// 	var user m.User
// 	err := c.GetDB().Where("id = ?", id).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
	
// 	return &user, nil
// }

// Find all users
func (ur *userRepository) FindAllUsers() ([]models.User, error) {
	var users []models.User
	err := ur.db.Find(&users).Error
	return users, err
}

// Create user
func (ur *userRepository) CreateUser(u *models.User) (*models.User, error) {
	err := ur.db.Model(&u).Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}


// Update user
func (ur *userRepository) UpdateUser(u *models.User) (*models.User, error) {

	err := ur.db.Model(&u).Updates(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Delete user by id
func (ur *userRepository) DeleteUserById(id int) error {
	err := ur.db.Where("id = ?", id).Delete(&models.User{}).Error
	return err
}

