package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gmshuvo/go-gin-postgres/controllers"
)

func UserRouters(router *gin.RouterGroup) {
	userController := controllers.UserController{}
	router.GET("/users", userController.FindAll)
	router.GET("/users/:id", userController.FindById)
	router.POST("/users", userController.Create)
	router.PUT("/users/:id", userController.Update)
	router.DELETE("/users/:id", userController.Delete)
}
