package utils

import (
	"os"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gmshuvo/go-gin-postgres/models"
)

// GenerateToken is used to generate token
func GenerateToken(u *models.User) (string, error) {
	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
