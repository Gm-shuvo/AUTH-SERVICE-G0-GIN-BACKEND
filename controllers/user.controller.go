package controllers

import (
	"net/http"
	"strconv"


	"github.com/gin-gonic/gin"

	"github.com/gmshuvo/go-gin-postgres/models"
	"github.com/gmshuvo/go-gin-postgres/utils"
)

type UserController struct {
	UserService models.UserService
}

// func NewUserController(us models.UserService, timeout time.Duration) *userController {
// 	return &userController{
// 		userService: us,

// 	}
// }


func (uc *UserController) FindAll(c *gin.Context) {
	users, err := uc.UserService.FindAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// func (uc *UserController) FindById(c *gin.Context) {
// 	// id := c.Param("id")
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	user, err := uc.UserService.FindUserById(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

func (uc *UserController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// check if user exists
	
	// _, err = uc.UserService.FindUserById(id)

	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	// 	return
	// }


	// check login status
	if !utils.IsLoggedIn(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var updatedUserData models.User
	if err := c.ShouldBindJSON(&updatedUserData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUserData.ID = uint(id) // Assuming ID is of type uint in your User model


	

	updatedUser, err := uc.UserService.UpdateUser(c, &updatedUserData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// Delete user by id
func (uc *UserController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = uc.UserService.DeleteUserById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
