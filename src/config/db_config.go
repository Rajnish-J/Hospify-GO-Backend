package config

import (
	"fmt"
	"log"
	"os"

	"hospify/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Example DSN: "username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to MySQL: %v", err)
	}

	fmt.Println("✅ Connected to MySQL using GORM")

	// Auto-migrate your models
	if err := DB.AutoMigrate(

		// Independent lookup tables first
		&models.Role{},
		&models.State{},
		&models.City{},
		&models.Status{},
		&models.Specialization{},
		&models.PaymentMethod{},

		// Tables with dependencies on lookup tables
		&models.User{},
		&models.Address{},
		&models.Clinic{},
		&models.Medicine{},
		&models.MedicalHistory{},
		&models.InsuranceProvider{},

		// Tables depending on above
		&models.Insurance{},
		&models.Patient{},
		&models.Doctor{},
		&models.Appointment{},
		&models.Prescription{},
		&models.Payment{},
	); err != nil {

		log.Fatalf("❌ AutoMigrate failed: %v", err)

	}

}
