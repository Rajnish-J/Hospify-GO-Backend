package dto

import (
	"hospify/src/models"
)

type ClinicDTO struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (c *ClinicDTO) ToModel() *models.Clinic {
	return &models.Clinic{
		ClinicID:   c.ID,
		ClinicName: c.Name,
		Address:    c.Address,
		Phone:      c.Phone,
	}
}

func FromClinicModel(m *models.Clinic) *ClinicDTO {
	return &ClinicDTO{
		ID:      m.ClinicID,
		Name:    m.ClinicName,
		Address: m.Address,
		Phone:   m.Phone,
	}
}
