package repositories

import (
	"github.com/gmshuvo/go-gin-postgres/models"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

// Register Repository
func (ar *authRepository) Register(u *models.User) (*models.User, error) {
	err := ar.db.Model(&u).Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Login Repository
func (ar *authRepository) Login(u *models.User) (*models.User, error) {
	var user models.User
	err := ar.db.Where("email = ?", u.Email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Logout Repository
func (ar *authRepository) Logout(u *models.User) (*models.User, error) {
	return u, nil
}
