package dto

import (
	"hospify/src/models"
	"time"
)

type AppointmentDTO struct {
	ID                 uint      `json:"id"`
	DoctorID           uint      `json:"doctor_id"`
	PatientID          uint      `json:"patient_id"`
	DateOfAppointment  string    `json:"date_of_appointment"`  // use RFC3339 format
	TimeOfAppointment  string    `json:"time_of_appointment"`  // "15:04" format
	DiagnosisID        *uint     `json:"diagnosis_id,omitempty"`
	StatusID           uint      `json:"status_id"`
}

func (a *AppointmentDTO) ToModel() *models.Appointment {
	date, _ := time.Parse("2006-01-02", a.DateOfAppointment)
	timeapp, _ := time.Parse("15:04", a.TimeOfAppointment)

	return &models.Appointment{
		AppointmentID:     a.ID,
		DoctorID:          a.DoctorID,
		PatientID:         a.PatientID,
		DateOfAppointment: date,
		TimeOfAppointment: timeapp,
		DiagnosisID:       a.DiagnosisID,
		StatusID:          a.StatusID,
	}
}

func FromAppointmentModel(m *models.Appointment) *AppointmentDTO {
	return &AppointmentDTO{
		ID:                m.AppointmentID,
		DoctorID:          m.DoctorID,
		PatientID:         m.PatientID,
		DateOfAppointment: m.DateOfAppointment.Format("2006-01-02"),
		TimeOfAppointment: m.TimeOfAppointment.Format("15:04"),
		DiagnosisID:       m.DiagnosisID,
		StatusID:          m.StatusID,
	}
}
