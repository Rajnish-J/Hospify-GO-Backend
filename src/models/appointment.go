package models

import "time"

type Appointment struct {
	AppointmentID     uint      	  `gorm:"primaryKey"`
	DoctorID          uint      	  `gorm:"foreignKey:DoctorID"`
	PatientID         uint      	  `gorm:"foreignKey:PatientID"`
	DateOfAppointment time.Time 	  `gorm:"not null"`
	TimeOfAppointment string    	  `gorm:"not null"`
	DiagnosisID       *uint			  `gorm:"foreignKey:DiagnosisID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`	
	StatusID          uint				`gorm:"foreignKey:StatusID"`
	CreatedAt         time.Time
	UpdatedAt         *time.Time
}
