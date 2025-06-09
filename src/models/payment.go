package models

import "time"

type Payment struct {
	PaymentID       uint    	`gorm:"primaryKey"`
	AppointmentID   uint    	`gorm:"not null"`
	PatientID       uint    	`gorm:"not null"`
	PaymentMethodID uint    	`gorm:"not null"`
	StatusID        string  	`gorm:"not null"`
	Amount          float64 	`gorm:"not null"`
	DateOfPayment   time.Time 	`gorm:"not null"`
	Notes 			string 		`gorm:"type:text"`
	Balance 		float64 	`gorm:"not null"`
}