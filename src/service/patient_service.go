package service

import (
	"hospify/src/config"
	"hospify/src/dto"
	"hospify/src/validator"
)

func CreatePatient(p *dto.PatientDTO) (*dto.PatientDTO, error) {
	model := p.ToModel()
	created, err := validator.CreatePatient(config.DB, model)
	if err != nil {
		return nil, err
	}
	return dto.FromModel(created), nil
}

func GetAllPatients() ([]dto.PatientDTO, error) {
	models, err := validator.GetAllPatients(config.DB)
	if err != nil {
		return nil, err
	}

	var daos []dto.PatientDTO
	for _, m := range models {
		daos = append(daos, *dto.FromModel(&m))
	}
	return daos, nil
}

func GetPatientByID(id uint) (*dto.PatientDTO, error) {
	model, err := validator.GetPatientByID(config.DB, id)
	if err != nil {
		return nil, err
	}
	return dto.FromModel(model), nil
}

func UpdatePatient(p *dto.PatientDTO) (*dto.PatientDTO, error) {
	model := p.ToModel()
	err := validator.UpdatePatient(config.DB, model)
	if err != nil {
		return nil, err
	}
	return dto.FromModel(model), nil
}

func DeletePatient(id uint) error {
	return validator.DeletePatient(config.DB, id)
}
