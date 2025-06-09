package models

import "time"

type Appointment struct {
<<<<<<< HEAD
	AppointmentID     uint      `gorm:"primaryKey"`
	DoctorID          uint      `gorm:"not null"`
	PatientID         uint      `gorm:"not null"`
	DateOfAppointment time.Time `gorm:"not null"`
	TimeOfAppointment time.Time    `gorm:"not null"`
=======
	AppointmentID     uint      	  `gorm:"primaryKey"`
	DoctorID          uint      	  `gorm:"not null"`
	PatientID         uint      	  `gorm:"not null"`
	DateOfAppointment time.Time 	  `gorm:"not null"`
	TimeOfAppointment string    	  `gorm:"not null"`
>>>>>>> cd4e49e0dc061abf54be13ed20e0a52daf7ac933
	DiagnosisID       *uint
	StatusID          uint
	CreatedAt         time.Time
	UpdatedAt         *time.Time
	Doctor            Doctor          `gorm:"foreignKey:DoctorID"`
	Patient           Patient         `gorm:"foreignKey:PatientID"`
	Diagnosis         *MedicalHistory `gorm:"foreignKey:DiagnosisID"`
	Status            Status          `gorm:"foreignKey:StatusID"`
}
