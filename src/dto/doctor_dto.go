package dto

import (
	"hospify/src/models"
)

type DoctorDTO struct {
	ID               uint    `json:"id"`
	Name             string  `json:"name"`
	Email            string  `json:"email"`
	Phone            string  `json:"phone"`
	Password         string  `json:"password"`
	DOB              string  `json:"dob"`
	SpecializationID uint    `json:"specialization_id"`
	ClinicID         uint    `json:"clinic_id"`
	FeesPerSession   float64 `json:"fees_per_session"`
}

func (d *DoctorDTO) ToModel() *models.Doctor {
	return &models.Doctor{
		DoctorID:         d.ID,
		DoctorName:       d.Name,
		DoctorEmail:      d.Email,
		DoctorPhone:      d.Phone,
		DoctorPassword:   d.Password,
		SpecializationID: d.SpecializationID,
		ClinicID:         d.ClinicID,
		FeesPerSession:   d.FeesPerSession,
	}
}

func FromDoctorModel(m *models.Doctor) *DoctorDTO {
	return &DoctorDTO{
		ID:               m.DoctorID,
		Name:             m.DoctorName,
		Email:            m.DoctorEmail,
		Phone:            m.DoctorPhone,
		Password:         m.DoctorPassword,
		SpecializationID: m.SpecializationID,
		ClinicID:         m.ClinicID,
		FeesPerSession:   m.FeesPerSession,
	}
}
