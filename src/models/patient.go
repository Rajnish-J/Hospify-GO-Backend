package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	PatientName     string `gorm:"not null"`
	PatientAge      int
	PatientGender   string
	PatientEmail    string `gorm:"unique"`
	PatientPhone    string
	PatientAddress  string
}
