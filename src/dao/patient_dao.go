package dao

import (
	"hospify/src/models"
	"gorm.io/gorm"
)

type PatientDAO struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Age            int    `json:"age"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
}

// Convert DAO to Model
func (p *PatientDAO) ToModel() *models.Patient {
	return &models.Patient{
		Model:          gorm.Model{ID: p.ID},
		PatientName:    p.Name,
		PatientAge:     p.Age,
		PatientGender:  p.Gender,
		PatientEmail:   p.Email,
		PatientPhone:   p.Phone,
		PatientAddress: p.Address,
	}
}

// Convert Model to DAO
func FromModel(m *models.Patient) *PatientDAO {
	return &PatientDAO{
		ID:      m.ID,
		Name:    m.PatientName,
		Age:     m.PatientAge,
		Gender:  m.PatientGender,
		Email:   m.PatientEmail,
		Phone:   m.PatientPhone,
		Address: m.PatientAddress,
	}
}
