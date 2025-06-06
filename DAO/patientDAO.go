package dao

import (
	"goBackend/config"
	"goBackend/Models"
)

// GetAllPatients retrieves all patients using GORM
func GetAllPatients() ([]Models.Patient, error) {
	var patients []Models.Patient
	result := config.DB.Find(&patients)
	return patients, result.Error
}

// CreatePatient inserts a new patient using GORM
func CreatePatient(p Models.Patient) error {
	result := config.DB.Create(&p)
	return result.Error
}

// UpdatePatient updates an existing patient using GORM
func UpdatePatient(p Models.Patient) error {
	result := config.DB.Save(&p) // GORM updates based on primary key
	return result.Error
}

// DeletePatient deletes a patient by ID using GORM
func DeletePatient(id int) error {
	result := config.DB.Delete(&Models.Patient{}, id)
	return result.Error
}
