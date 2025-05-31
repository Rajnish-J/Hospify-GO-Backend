package DAO
 
import (
    "goBackend/config"
    "goBackend/Models"
)
 
func GetAllPatients() ([]Models.PatientModel, error) {
    rows, err := config.DB.Query("SELECT * FROM Patiet_Details")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
 
    var patients []Models.PatientModel
    for rows.Next() {
        var p Models.PatientModel
        err = rows.Scan(&p.PatientId, &p.PatientName, &p.PatientPassword, &p.PatientEmail, &p.PatientPhone, &p.PatientAddressId, &p.PatientDOB)
        if err != nil {
            return nil, err
        }
        patients = append(patients, p)
    }
    return patients, nil
}
 
func CreatePatient(p Models.PatientModel) error {
    _, err := config.DB.Exec(`
        INSERT INTO Patiet_Details
        (Patiet_ID, Patiet_Name, Password, Patiet_Email, Patiet_phone, AddressID, Patiet_DOB, Employment_Status, Annual_Income)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
        p.PatientId, p.PatientName, p.PatientPassword, p.PatientEmail, p.PatientPhone, p.PatientAddressId, p.PatientDOB)
    return err
}
 
func UpdatePatient(p Models.PatientModel) error {
    _, err := config.DB.Exec(`
        UPDATE Patiet_Details
        SET Password=?, Patiet_Email=?, Patiet_phone=?, AddressID=?, Patiet_DOB=?, Employment_Status=?, Annual_Income=?
        WHERE Patiet_ID=?`, p.PatientPassword, p.PatientEmail, p.PatientPhone, p.PatientAddressId, p.PatientDOB, p.PatientId)
    return err
}
 
func DeletePatient(id int) error {
    _, err := config.DB.Exec("DELETE FROM Patiet_Details WHERE Patiet_ID=?", id)
    return err
}