package models

import "time"

type Appointment struct {
	AppointmentID     uint      `gorm:"primaryKey"`
	DoctorID          uint      `gorm:"not null"`
	PatientID         uint      `gorm:"not null"`
	DateOfAppointment time.Time `gorm:"not null"`
	TimeOfAppointment time.Time    `gorm:"not null"`
	DiagnosisID       *uint
	StatusID          uint
	CreatedAt         time.Time
	UpdatedAt         *time.Time
	Doctor            Doctor          `gorm:"foreignKey:DoctorID"`
	Patient           Patient         `gorm:"foreignKey:PatientID"`
	Diagnosis         *MedicalHistory `gorm:"foreignKey:DiagnosisID"`
	Status            Status          `gorm:"foreignKey:StatusID"`
}
