package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Token string `gorm:"uniqueIndex"`
}

func init() {
	if os.Getenv("ENVIRONMENT") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("No .env file found or error loading .env file")
		}
	}
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tokens.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	err = db.AutoMigrate(&Token{})

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}

func SeedToken(db *gorm.DB) {
	token := os.Getenv("TOKEN")

	if token == "" {
		log.Fatal("No token found in environment")
	}

	result := db.FirstOrCreate(&Token{Token: token})

	if result.Error != nil {
		log.Fatal("Failed to seed token:", result.Error)
	}
}
