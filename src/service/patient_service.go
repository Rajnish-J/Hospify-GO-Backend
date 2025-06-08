package service

import (
	"hospify/src/config"
	"hospify/src/dao"
	"hospify/src/validator"
)

func CreatePatient(p *dao.PatientDAO) (*dao.PatientDAO, error) {
	model := p.ToModel()
	created, err := validator.CreatePatient(config.DB, model)
	if err != nil {
		return nil, err
	}
	return dao.FromModel(created), nil
}

func GetAllPatients() ([]dao.PatientDAO, error) {
	models, err := validator.GetAllPatients(config.DB)
	if err != nil {
		return nil, err
	}

	var daos []dao.PatientDAO
	for _, m := range models {
		daos = append(daos, *dao.FromModel(&m))
	}
	return daos, nil
}

func GetPatientByID(id uint) (*dao.PatientDAO, error) {
	model, err := validator.GetPatientByID(config.DB, id)
	if err != nil {
		return nil, err
	}
	return dao.FromModel(model), nil
}

func UpdatePatient(p *dao.PatientDAO) (*dao.PatientDAO, error) {
	model := p.ToModel()
	err := validator.UpdatePatient(config.DB, model)
	if err != nil {
		return nil, err
	}
	return dao.FromModel(model), nil
}

func DeletePatient(id uint) error {
	return validator.DeletePatient(config.DB, id)
}
