package dto

import (
	"hospify/src/models"
	"time"
)

type DoctorDTO struct {
	ID               uint    `json:"id"`
	Name             string  `json:"name"`
	Email            string  `json:"email"`
	Phone            string  `json:"phone"`
	Password         string  `json:"password"`
	DOB              string  `json:"dob"` // "YYYY-MM-DD"
	SpecializationID uint    `json:"specialization_id"`
	ClinicID         uint    `json:"clinic_id"`
	FeesPerSession   float64 `json:"fees_per_session"`
}

func (d *DoctorDTO) ToModel() (*models.Doctor, error) {
	dob, err := time.Parse("2006-01-02", d.DOB)
	if err != nil {
		return nil, err
	}

	return &models.Doctor{
		DoctorID:         d.ID,
		DoctorName:       d.Name,
		DoctorEmail:      d.Email,
		DoctorPhone:      d.Phone,
		DoctorPassword:   d.Password,
		DoctorDOB:        dob,
		SpecializationID: d.SpecializationID,
		ClinicID:         d.ClinicID,
		FeesPerSession:   d.FeesPerSession,
	}, nil
}

func FromDoctorModel(m *models.Doctor) *DoctorDTO {
	return &DoctorDTO{
		ID:               m.DoctorID,
		Name:             m.DoctorName,
		Email:            m.DoctorEmail,
		Phone:            m.DoctorPhone,
		Password:         m.DoctorPassword,
		DOB:              m.DoctorDOB.Format("2006-01-02"),
		SpecializationID: m.SpecializationID,
		ClinicID:         m.ClinicID,
		FeesPerSession:   m.FeesPerSession,
	}
}
