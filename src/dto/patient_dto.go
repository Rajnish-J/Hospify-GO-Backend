package dto

import (
	"hospify/src/models"
)

// PatientDTO defines the structure for transferring patient data
type PatientDTO struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Password         string `json:"password"`
	AddressID        uint   `json:"address_id"`
	DOB              string `json:"dob"` // format: YYYY-MM-DD (optional: use time.Time if needed)
	EmploymentStatus string `json:"employment_status,omitempty"`
	AnnualIncome     float64 `json:"annual_income,omitempty"`
	InsuranceID      *uint   `json:"insurance_id,omitempty"`
}

// ToModel converts a DTO to a GORM Patient model
func (p *PatientDTO) ToModel() *models.Patient {
	return &models.Patient{
		PatientID:        p.ID,
		PatientName:      p.Name,
		PatientEmail:     p.Email,
		PatientPhone:     p.Phone,
		AddressID:        p.AddressID,
		EmploymentStatus: p.EmploymentStatus,
		AnnualIncome:     p.AnnualIncome,
		InsuranceID:      p.InsuranceID,
		// Note: Convert DOB string to time.Time in service layer if needed
	}
}

// FromModel converts a GORM model to PatientDTO
func FromModel(m models.Patient) *PatientDTO {
	return &PatientDTO{
		ID:               m.PatientID,
		Name:             m.PatientName,
		Email:            m.PatientEmail,
		Phone:            m.PatientPhone,
		AddressID:        m.AddressID,
		EmploymentStatus: m.EmploymentStatus,
		AnnualIncome:     m.AnnualIncome,
		InsuranceID:      m.InsuranceID,
		// Note: Format m.PatientDOB as string in service/controller if needed
	}
}
