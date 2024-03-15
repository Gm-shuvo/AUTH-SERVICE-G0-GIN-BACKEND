package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gmshuvo/go-gin-postgres/controllers"
	"github.com/gmshuvo/go-gin-postgres/repositories"
	"github.com/gmshuvo/go-gin-postgres/services"
	"github.com/gmshuvo/go-gin-postgres/middlewares"
	"gorm.io/gorm"
)

func NewUserRouters(router *gin.RouterGroup, db *gorm.DB, timeout time.Duration) {
	ur := repositories.NewUserRepository(db)
	uc := &controllers.UserController{
		UserService: services.NewUserService(ur, timeout),
	}
	router.GET("/users", uc.FindAll)
	// router.GET("/user/:id", uc.FindById)
	// make group for private routes
	private := router.Group("")
	private.Use(middleware.RequireAuth)
	private.PATCH("/user/:id" , uc.Update)
	private.DELETE("/user/:id", uc.Delete)
	
}
	
