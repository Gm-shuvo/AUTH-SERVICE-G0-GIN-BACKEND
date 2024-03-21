package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gmshuvo/go-gin-postgres/controllers"
	"github.com/gmshuvo/go-gin-postgres/repositories"
	"github.com/gmshuvo/go-gin-postgres/services"
	"gorm.io/gorm"
)

func AuthRouters(router *gin.RouterGroup, db *gorm.DB, timeout time.Duration) {
	ar := repositories.NewAuthRepository(db)
	ac := &controllers.AuthController{
		AuthService: services.NewAuthService(ar, timeout),
		
	}
	router.POST("/login", ac.Login)
	router.POST("/register", ac.Register)
	router.POST("/logout", ac.Logout)
}
