package routes

import (
	"github.com/gin-gonic/gin"
	"ginGo/src/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userController := new(controllers.UserController)

	r.GET("/users", userController.FindAll)
	r.GET("/users/:id", userController.FindById)
	r.POST("/users", userController.Create)
	r.PUT("/users/:id", userController.Update)
	r.DELETE("/users/:id", userController.Delete)

	return r
}

// Path: src/routes/user.go

