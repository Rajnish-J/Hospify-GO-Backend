package service

import (
	"hospify/src/config"
	"hospify/src/dto"
	"hospify/src/validator"
)

// CreatePatient creates a new patient record.
func CreatePatient(p *dto.PatientDTO) (*dto.PatientDTO, error) {
	model := p.ToModel()
	created, err := validator.CreatePatient(config.DB, model)
	if err != nil {
		return nil, err
	}
	return dto.FromModel(*created), nil
}

// GetAllPatients retrieves all patient records.
func GetAllPatients() ([]dto.PatientDTO, error) {
	models, err := validator.GetAllPatients(config.DB)
	if err != nil {
		return nil, err
	}

	dtos := make([]dto.PatientDTO, 0, len(models))
	for _, m := range models {
		dtos = append(dtos, *dto.FromModel(m))
	}
	return dtos, nil
}

// GetPatientByID retrieves a patient by their ID.
func GetPatientByID(id uint) (*dto.PatientDTO, error) {
	model, err := validator.GetPatientByID(config.DB, id)
	if err != nil {
		return nil, err
	}
	return dto.FromModel(*model), nil
}

// UpdatePatient updates an existing patient record.
func UpdatePatient(p *dto.PatientDTO) (*dto.PatientDTO, error) {
	model := p.ToModel()
	err := validator.UpdatePatient(config.DB, model)
	if err != nil {
		return nil, err
	}
	return dto.FromModel(*model), nil
}

// DeletePatient deletes a patient record by ID.
func DeletePatient(id uint) error {
	return validator.DeletePatient(config.DB, id)
}
