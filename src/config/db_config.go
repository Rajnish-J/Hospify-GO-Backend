package config

import (
	"fmt"
	"log"
	"os"

	"hospify/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL using GORM")

	if err := DB.AutoMigrate(&models.Patient{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
}
