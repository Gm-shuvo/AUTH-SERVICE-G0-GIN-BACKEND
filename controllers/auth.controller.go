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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// auth service
	user, err := ac.AuthService.Register(c, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User created",
		"user":    user,
	})
}

// Login user
func (ac *AuthController) Login(c *gin.Context) {
	var body models.User
	// var _body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	// err := config.GetDB().Where("email = ?", body.Email).First(&_body).Error

	user, err := ac.AuthService.Login(c, &body)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	println(user.Password, body.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid User and Password"})
		return
	}

	// generate token with user id and email
	token, err := utils.GenerateToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating the token"})
		return
	}

	// set token in cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"expires_in": 3600 * 24 * 30,
	})
}

// Logout user
func (u *AuthController) Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "User logged out"})
}
