package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gmshuvo/go-gin-postgres/models"
	"github.com/gmshuvo/go-gin-postgres/utils"
	"golang.org/x/crypto/bcrypt"

	"net/http"
)

type AuthController struct {
	AuthService models.AuthService
}

// Register new user
func (ac *AuthController) Register(c *gin.Context) {
	// validate body
	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid input",
			Details: []string{err.Error()},
		})
		return
	}
	// auth service
	user, err := ac.AuthService.Register(c, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Details: []string{err.Error()},
		})
		return
	}

	c.JSON(http.StatusCreated,models.SucecssResponse{
		Code:    http.StatusCreated,
		Message: "User created successfully",
		Data:    user,
	})
}

// Login user
func (ac *AuthController) Login(c *gin.Context) {
	var body models.User
	// var _body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Invalid input",
				Details: []string{err.Error()},
			})
		return
	}

	// err := config.GetDB().Where("email = ?", body.Email).First(&_body).Error

	user, err := ac.AuthService.Login(c, &body)

	if err != nil {
		c.JSON(http.StatusNotFound,
			models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Invalid credentials",
				Details: []string{err.Error()},
			})
		return
	}

	println(user.Password, body.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusNotFound, 
			models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Invalid credentials",
				Details: []string{err.Error()},
			})
		return
	}

	// generate token with user id and email
	access_token, refresh_token, err := utils.GenerateNewToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, 
			models.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error generating token",
				Details: []string{err.Error()},
			})
		return
	}

	// set token in cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", access_token, 3600, "", "", false, true)
	c.SetCookie("RefreshToken", refresh_token, 3600, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token":      access_token,
		"refresh_token":     refresh_token,
		"expires_in": 3600,
	})
}

// Logout user
func (u *AuthController) Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.SetCookie("RefreshToken", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "User logged out"})
}
