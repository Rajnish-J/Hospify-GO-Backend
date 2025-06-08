package models

import "time"

type Doctor struct {
	DoctorID         uint `gorm:"primaryKey"`
	DoctorName       string
	DoctorDOB        time.Time
	DoctorPassword   string
	DoctorEmail      string
	DoctorPhone      string
	SpecializationID uint
	ClinicID         uint
	FeesPerSession   float64
	Clinic           Clinic         `gorm:"foreignKey:ClinicID"`
	Specialization   Specialization `gorm:"foreignKey:SpecializationID"`
}
