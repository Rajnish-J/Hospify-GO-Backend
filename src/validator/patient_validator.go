package validator

import (
	"errors"
	"hospify/src/models"
	"gorm.io/gorm"
)

func CreatePatient(db *gorm.DB, patient *models.Patient) (*models.Patient, error) {
	if err := db.Create(patient).Error; err != nil {
		return nil, errors.New("failed to create patient")
	}
	return patient, nil
}

func GetPatientByID(db *gorm.DB, id uint) (*models.Patient, error) {
	var patient models.Patient
	if err := db.First(&patient, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("patient not found")
		}
		return nil, err
	}
	return &patient, nil
}

func GetAllPatients(db *gorm.DB) ([]models.Patient, error) {
	var patients []models.Patient
	if err := db.Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}

func UpdatePatient(db *gorm.DB, patient *models.Patient) error {
	return db.Save(patient).Error
}

func DeletePatient(db *gorm.DB, id uint) error {
	return db.Delete(&models.Patient{}, id).Error
}
