package models

import "time"

type Doctor struct {
	DoctorID         uint `gorm:"primaryKey"`
	DoctorName       string
	DoctorDOB        time.Time
	DoctorPassword   string
	DoctorEmail      string
	DoctorPhone      string
	SpecializationID uint		`gorm:"foreignKey:SpecializationID"`
	ClinicID         uint		`gorm:"foreignKey:ClinicID"`
	FeesPerSession   float64      
}
