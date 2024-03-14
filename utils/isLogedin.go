package utils

import "github.com/gin-gonic/gin"

func IsLoggedIn(c *gin.Context) bool {
	// Check cookie is expired or not 
	// if expired return false
	// else return true
	// c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	cookie, err := c.Cookie("Authorization")
	if err != nil {
		return false
	}
	// if cookie is not expired
	if cookie == "" {
		return false
	}

	return true

}