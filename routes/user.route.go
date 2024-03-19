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
	
	// make group for private routes
	private := router.Group("")
	private.Use(middleware.RequireAuth)
	private.PATCH("/user/:id" , uc.Update)
	private.DELETE("/user/:id", uc.Delete)
	private.GET("/user/:id", uc.FindUserById)
	// private.GET("/user", uc.FindUserByEmail)
	
}
	
