package main

import (
	"ginGo/src/models"
	"ginGo/src/routes"
	"ginGo/src/utils"
)

func main() {
	utils.LoadEnv()
	models.InitDB()
	models.MigrateDB()
	router := routes.SetupRouter()
	router.Run(":8080")
}
