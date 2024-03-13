package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/gmshuvo/go-gin-postgres/config"
	repo "github.com/gmshuvo/go-gin-postgres/repositories"
	r "github.com/gmshuvo/go-gin-postgres/routes"
	u "github.com/gmshuvo/go-gin-postgres/utils"
	"os"
)

func main() {
	u.LoadEnv()
	c.InitDB()
	defer c.CloseDB()
	router := gin.Default()
	// make group for  routes
	Routes := router.Group("/api/v1")
	router.Use(gin.Logger())
	router.SetTrustedProxies([]string{"192.168.0.1"})
	router.Use(gin.Recovery())
	// Migrate the schema
	repo.MigrateDB()

	r.UserRouters(Routes)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not specified
	}
	router.Run(":" + port)
}
