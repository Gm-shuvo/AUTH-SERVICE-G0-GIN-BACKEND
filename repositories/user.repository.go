package repositories

import (
	c "github.com/gmshuvo/go-gin-postgres/config"
	m "github.com/gmshuvo/go-gin-postgres/models"
)


// // Create user 
// func CreateUser(u *m.User) (*m.User, error) {
// 	err := c.GetDB().Model(&u).Create(u).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return u, nil
// }

// Find user by email
func FindUserByEmail(email string) (*m.User, error) {
	var user m.User
	err := c.GetDB().Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserById(id int) (*m.User, error) {
	var user m.User
	err := c.GetDB().Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}