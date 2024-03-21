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
	r := gin.Default()

	// Initialize the logger
	logger := utils.Initialize()
	defer logger.Sync()

	

	// make group for  routes
	gRouter := r.Group("/api/v1")

	routes.SetUp(gRouter, db, 2)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not specified
	}
	r.Run(":" + port)
}
