package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUp(gRouter *gin.RouterGroup, db *gorm.DB, timeout time.Duration) {
	NewUserRouters(gRouter, db, timeout)
	AuthRouters(gRouter, db, timeout)
}