package config

import (
	
	"github.com/gmshuvo/go-gin-postgres/models"
)

// MigrateDB is used to migrate the schema
func MigrateDB() {
	GetDB().AutoMigrate(&models.User{})
}