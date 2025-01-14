package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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

}
