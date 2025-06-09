package models

import "time"

type MedicalHistory struct {
	HistoryID       uint      `gorm:"primaryKey"`
	Diagnosis       string    `gorm:"not null"`
	Treatment       string    `gorm:"not null"`
	DateOfDiagnosis time.Time `gorm:"not null"`
}
