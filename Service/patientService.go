package Service
 
import (
    "goBackend/DAO"
    "goBackend/Models"
)
 
func GetAllPatients() ([]Models.PatientModel, error) {
    return DAO.GetAllPatients()
}
 
func CreatePatient(p Models.PatientModel) error {
    return DAO.CreatePatient(p)
}
 
func UpdatePatient(p Models.PatientModel) error {
    return DAO.UpdatePatient(p)
}
 
func DeletePatient(id int) error {
    return DAO.DeletePatient(id)
}