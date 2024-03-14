package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gmshuvo/go-gin-postgres/controllers"
)

func AuthRouters(router *gin.RouterGroup) {
	userController := controllers.AuthController{}
	router.POST("/login", userController.Login)
	router.POST("/register", userController.Register)
	router.POST("/logout", userController.Logout)
}
