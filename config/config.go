package config

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()
}
