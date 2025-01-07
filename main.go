package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100;unique;not null"`
	Password string `gorm:"size:255;not null"`
}

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dsn := os.Getenv("NEON_DATABASE_URL")

	if dsn == "" {
		log.Fatalf("DATABASE URL is not set in the environment")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Neon PostgreSQL: %v", err)
	}
	fmt.Println("Connected to Neon PostgreSQL with GORM!")

	err = db.AutoMigrate(&User{})

	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	fmt.Println("Database schema migrated!")

	newUser := User{
		Name:     "Arpan Das",
		Email:    "arpandas@arpan.tech",
		Password: "123456",
	}
	result := db.Create(&newUser)
	if result.Error != nil {
		log.Fatalf("Failed to insert user: %v", result.Error)
	}
	fmt.Printf("Inserted user: %+v\n", newUser)

	var users []User
	result = db.Find(&users)
	if result.Error != nil {
		log.Fatalf("Failed to fetch users: %v", result.Error)
	}
	fmt.Println("Fetched users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}
