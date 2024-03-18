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

func (uc *UserController) FindAll(c *gin.Context) {
	users, err := uc.UserService.FindAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Details: []string{err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format",
			Details: []string{err.Error()},
		})
		return
	}

	if !utils.IsLoggedIn(c) {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
			Details: []string{"You must be logged in to perform this action"},
		})
		return
	}

	var updatedUserData models.User
	if err := c.ShouldBindJSON(&updatedUserData); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error binding JSON",
			Details: []string{err.Error()},
		})
		return
	}

	updatedUserData.ID = uint(id) // Assuming ID is of type uint in your User model
	updatedUser, err := uc.UserService.UpdateUser(c, &updatedUserData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Details: []string{err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (uc *UserController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format",
		})
		return
	}

	if !utils.IsLoggedIn(c) {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	err = uc.UserService.DeleteUserById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Details: []string{err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
