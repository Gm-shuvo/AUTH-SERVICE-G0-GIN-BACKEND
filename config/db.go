package config

import (
	"fmt"
	"os"
	// m "github.com/gmshuvo/go-gin-postgres/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"), // Optionally set sslmode via environment variable
	)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
}

func GetDB() *gorm.DB {
	return db
}


func CloseDB() {
	psqlDB, err := db.DB()
	if err != nil {
		panic("failed to close database connection")
	}
	psqlDB.Close()
}
