package models

import "time"

type MedicalHistory struct {
	HistoryID       uint `gorm:"primaryKey"`
	Diagnosis       string
	Treatment       string
	DateOfDiagnosis time.Time
}
