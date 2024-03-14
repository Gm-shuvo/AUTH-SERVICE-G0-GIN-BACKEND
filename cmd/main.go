package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/gmshuvo/go-gin-postgres/config"
	r "github.com/gmshuvo/go-gin-postgres/routes"
	u "github.com/gmshuvo/go-gin-postgres/utils"
	"os"
)

func main() {
	// Load the environment variables
	u.LoadEnv()
	// Initialize the database
	c.InitDB()
	// Migrate the schema
	c.MigrateDB()
	// Close the connection
	defer c.CloseDB()

	// Create routers
	router := gin.Default()
	
	router.SetTrustedProxies([]string{"192.168.0.1"})
	// make group for  routes
	Routes := router.Group("/api/v1")

	r.AuthRouters(Routes)
	r.UserRouters(Routes)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not specified
	}
	router.Run(":" + port)
}
