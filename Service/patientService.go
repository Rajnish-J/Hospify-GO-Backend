package Service
 
import (
    "goBackend/DAO"
    "goBackend/Models"
)
 
func GetAllPatients() ([]Models.Patient, error) {
    return DAO.GetAllPatients()
}
 
func CreatePatient(p Models.Patient) error {
    return DAO.CreatePatient(p)
}
 
func UpdatePatient(p Models.Patient) error {
    return DAO.UpdatePatient(p)
}
 
func DeletePatient(id int) error {
    return DAO.DeletePatient(id)
}