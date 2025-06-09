package dto

import "hospify/src/models"

type PrescriptionDTO struct {
	ID         uint   `json:"id"`
	HistoryID  uint   `json:"history_id"`
	MedicineID uint   `json:"medicine_id"`
	Dosage     string `json:"dosage"`
	Duration   string `json:"duration"`
}
 
func (p *PrescriptionDTO) ToModel() *models.Prescription {
	return &models.Prescription{
		PrescriptionID: p.ID,
		HistoryID:      p.HistoryID,
		MedicineID:     p.MedicineID,
		Dosage:         p.Dosage,
		Duration:       p.Duration,
	}
}

func FromPrescriptionModel(m *models.Prescription) *PrescriptionDTO {
	return &PrescriptionDTO{
		ID:         m.PrescriptionID,
		HistoryID:  m.HistoryID,
		MedicineID: m.MedicineID,
		Dosage:     m.Dosage,
		Duration:   m.Duration,
	}
}
