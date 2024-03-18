package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gmshuvo/go-gin-postgres/config"
	// "github.com/gmshuvo/go-gin-postgres/middlewares"
	"github.com/gmshuvo/go-gin-postgres/routes"
	"github.com/gmshuvo/go-gin-postgres/utils"
)

func main() {
	// Load the environment variables
	utils.LoadEnv()
	
	// Initialize the database
	db := config.InitDB()
	// Migrate the schema
	config.MigrateDB()
	// Close the connection
	defer config.CloseDB()

	// Create routers
	gin := gin.Default()
	// gin.Use(middleware.ErrorHandlingMiddleware())
	
	gin.SetTrustedProxies([]string{"192.168.0.1"})
	// make group for  routes
	gRouter := gin.Group("/api/v1")

	routes.SetUp(gRouter, db,  2)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not specified
	}
	gin.Run(":" + port)
}
