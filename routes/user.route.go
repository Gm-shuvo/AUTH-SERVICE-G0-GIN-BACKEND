package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gmshuvo/go-gin-postgres/controllers"
	"github.com/gmshuvo/go-gin-postgres/repositories"
	"github.com/gmshuvo/go-gin-postgres/services"
	"gorm.io/gorm"
)

func NewUserRouters(router *gin.RouterGroup, db *gorm.DB, timeout time.Duration) {
	ur := repositories.NewUserRepository(db)
	uc := &controllers.UserController{
		UserService: services.NewUserService(ur, timeout),
	}
	router.GET("/users", uc.FindAll)
	// router.GET("/user/:id", uc.FindById)
	router.PUT("/user/:id", uc.Update)
	router.DELETE("/user/:id", uc.Delete)
	
}
	
