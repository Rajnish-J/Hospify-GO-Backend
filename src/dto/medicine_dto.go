package dto

import (
	"hospify/src/models"
)

type MedicineDTO struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Manufacturer string `json:"manufacturer"`
	DosageForm   string `json:"dosage_form"`
}

func (m *MedicineDTO) ToModel() *models.Medicine {
	return &models.Medicine{
		MedicineID:   m.ID,
		MedicineName: m.Name,
		Description:  m.Description,
		Manufacturer: m.Manufacturer,
		DosageForm:   m.DosageForm,
	}
}

func FromMedicineModel(model *models.Medicine) *MedicineDTO {
	return &MedicineDTO{
		ID:           model.MedicineID,
		Name:         model.MedicineName,
		Description:  model.Description,
		Manufacturer: model.Manufacturer,
		DosageForm:   model.DosageForm,
	}
}
