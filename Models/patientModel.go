package Models

import "time"

type PatientModel struct {
	PatientId 			int			`json:"patient_id"`
	PatientName     	string  	`json:"patient_name"`
	PatientEmail		string  	`json:"patient_email"`
	PatientPassword 	string  	`json:"patient_password"`
	PatientPhone    	int64   	`json:"patient_phone"`
	PatientAddressId	int     	`json:"patient_address_id"`
	PatientDOB          time.Time   `json:"patient_dob"`
	PatientAnnualIncome int64		`json:"patient_annual_income"`
}