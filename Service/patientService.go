package service

import (
	"goBackend/dao"
	"goBackend/Models"
)

// GetAllPatients calls the DAO to fetch all patients
func GetAllPatients() ([]Models.Patient, error) {
	return dao.GetAllPatients()
}

// CreatePatient calls the DAO to create a new patient
func CreatePatient(p Models.Patient) error {
	return dao.CreatePatient(p)
}

// UpdatePatient calls the DAO to update an existing patient
func UpdatePatient(p Models.Patient) error {
	return dao.UpdatePatient(p)
}

// DeletePatient calls the DAO to delete a patient by ID
func DeletePatient(id int) error {
	return dao.DeletePatient(id)
}
