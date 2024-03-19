package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gmshuvo/go-gin-postgres/config"
	"github.com/gmshuvo/go-gin-postgres/models"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Type:    "Unauthorized",
				Message: "No authorization token provided",
				Details: []string{err.Error()},
			})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Type:    "Unauthorized",
				Message: "Invalid authorization token",
				Details: []string{err.Error()},
			})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			// refresh token exits and valid
			_, err := c.Cookie("RefreshToken")

			if err != nil {
			c.JSON(http.StatusUnauthorized,
				models.ErrorResponse{
					Code:    http.StatusUnauthorized,
					Type:    "Unauthorized",
					Message: "Token has expired",
				})
			c.Abort()
			return
			}
			// TODO: Generate new access token and refresh token
		}

		userID := claims["sub"]
		var user models.User
		result := config.GetDB().First(&user, userID)
		if result.Error != nil || user.ID == 0 {
			c.JSON(http.StatusUnauthorized,
				models.ErrorResponse{
					Code:    http.StatusUnauthorized,
					Type:    "Unauthorized",
					Message: "User not found",
					Details: []string{result.Error.Error()},
				})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized,
			models.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Type:    "Unauthorized",
				Message: "Invalid authorization token",
			})
		c.Abort()
	}
}
