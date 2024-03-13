package main

import (
	"os"
	"github.com/gin-gonic/gin"
	c "github.com/gmshuvo/go-gin-postgres/config"
	repo "github.com/gmshuvo/go-gin-postgres/repositories"
	r "github.com/gmshuvo/go-gin-postgres/routes"
	u "github.com/gmshuvo/go-gin-postgres/utils"
	
)

func main() {
	u.LoadEnv()
	c.InitDB()
	defer c.CloseDB()
	router := gin.Default()
	// make group for  routes
	Routes := router.Group("/api/v1")
	r.UserRouters(Routes)

	// Migrate the schema
	repo.MigrateDB()
	router.Run(":" + os.Getenv("PORT"))
}
