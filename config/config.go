package config

import (
	"log"
	"os"

	"github.com/adasarpan404/gopostgrespoc/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error While Loading .env file: %v", err)
	}

	dsn := os.Getenv("NEON_DATABASE_URL")

	if dsn == "" {
		log.Fatalf("There is no Database URL set in the enviroment")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Profile{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	return db
}
