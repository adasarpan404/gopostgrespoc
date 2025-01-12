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
	ID       uint    `gorm:"primaryKey"`
	Name     string  `gorm:"size:100;not null"`
	Email    string  `gorm:"size:100;unique;not null"`
	Password string  `gorm:"size:255;not null"`
	Profile  Profile `gorm:"constraint:OnDelete:CASCADE;"`
}

type Profile struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   uint   `gorm:"unique;not null"`
	Bio      string `gorm:"size:255"`
	Location string `gorm:"size:100"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("NEON_DATABASE_URL")
	if dsn == "" {
		log.Fatal("NEON_DATABASE_URL is not set in the environment")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Neon PostgreSQL: %v", err)
	}
	fmt.Println("Connected to Neon PostgreSQL with GORM!")

	err = db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
	fmt.Println("Database schema migrated!")

	newUser := User{
		Name:     "John Doe",
		Email:    "john.doe@neon.tech",
		Password: "securepassword123",
		Profile: Profile{
			Bio:      "Software Developer at Neon",
			Location: "San Francisco, CA",
		},
	}
	result := db.Create(&newUser)
	if result.Error != nil {
		log.Fatalf("Failed to insert user with profile: %v", result.Error)
	}
	fmt.Printf("Inserted user with profile: %+v\n", newUser)

	var users []User
	result = db.Preload("Profile").Find(&users)
	if result.Error != nil {
		log.Fatalf("Failed to fetch users with profiles: %v", result.Error)
	}
	fmt.Println("Fetched users with profiles:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Bio: %s, Location: %s\n",
			user.ID, user.Name, user.Email, user.Profile.Bio, user.Profile.Location)
	}
}
