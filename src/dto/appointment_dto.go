package dto

import (
	"hospify/src/models"
	"time"
)

type AppointmentDTO struct {
	ID                 uint      	`json:"id"`
	DoctorID           uint      	`json:"doctor_id"`
	PatientID          uint      	`json:"patient_id"`
	DateOfAppointment  time.Time 	`json:"date_of_appointment"`  // use RFC3339 format
	TimeOfAppointment  time.Time 	`json:"time_of_appointment"`  // "15:04" format
	DiagnosisID        *uint     	`json:"diagnosis_id,omitempty"`
	StatusID           uint      	`json:"status_id"`
}

func (a *AppointmentDTO) ToModel() *models.Appointment {

	return &models.Appointment{
		AppointmentID:     a.ID,
		DoctorID:          a.DoctorID,
		PatientID:         a.PatientID,
		DateOfAppointment: a.DateOfAppointment,
		TimeOfAppointment: a.TimeOfAppointment.Format("15:04"),
		DiagnosisID:       a.DiagnosisID,
		StatusID:          a.StatusID,
	}
}

func FromAppointmentModel(m *models.Appointment) *AppointmentDTO {
	var timeOfAppointment time.Time
	if m.TimeOfAppointment != "" {
		parsedTime, err := time.Parse("15:04", m.TimeOfAppointment)
		if err == nil {
			timeOfAppointment = parsedTime
		}
	}
	return &AppointmentDTO{
		ID:                m.AppointmentID,
		DoctorID:          m.DoctorID,
		PatientID:         m.PatientID,
		DateOfAppointment: m.DateOfAppointment,
		TimeOfAppointment: timeOfAppointment,
		DiagnosisID:       m.DiagnosisID,
		StatusID:          m.StatusID,
	}
}
