package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gmshuvo/go-gin-postgres/controllers"
)

func UserRouters(router *gin.RouterGroup) {
	userController := controllers.UserController{}
	router.GET("/users", userController.FindAll)
	router.GET("/user/:id", userController.FindById)
	router.POST("/user", userController.Create)
	router.PUT("/user/:id", userController.Update)
	router.DELETE("/user/:id", userController.Delete)
}
