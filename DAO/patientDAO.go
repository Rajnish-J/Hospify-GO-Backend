package DAO

import (
	config "goBackend/Config"
	"goBackend/Models"
)

func GetAllPatients() ([]Models.Patient, error) {
	rows, err := config.DB.Query("SELECT patient_id, patient_name, patient_email, patient_password, patient_phone, patient_address_id, patient_dob, patient_annual_income FROM Patient")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []Models.Patient
	for rows.Next() {
		var p Models.Patient
		err = rows.Scan(
			&p.PatientId,
			&p.PatientName,
			&p.PatientEmail,
			&p.PatientPassword,
			&p.PatientPhone,
			&p.PatientAddressId,
			&p.PatientDOB,
			&p.PatientAnnualIncome,
		)
		if err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}
	return patients, nil
}

func CreatePatient(p Models.Patient) error {
	_, err := config.DB.Exec(`
        INSERT INTO Patient
        (patient_name, patient_email, patient_password, patient_phone, patient_address_id, patient_dob, patient_annual_income)
        VALUES (?, ?, ?, ?, ?, ?, ?)`,
		p.PatientName,
		p.PatientEmail,
		p.PatientPassword,
		p.PatientPhone,
		p.PatientAddressId,
		p.PatientDOB,
		p.PatientAnnualIncome,
	)
	return err
}

func UpdatePatient(p Models.Patient) error {
	_, err := config.DB.Exec(`
        UPDATE Patient
        SET patient_name=?, patient_email=?, patient_password=?, patient_phone=?, patient_address_id=?, patient_dob=?, patient_annual_income=?
        WHERE patient_id=?`,
		p.PatientName,
		p.PatientEmail,
		p.PatientPassword,
		p.PatientPhone,
		p.PatientAddressId,
		p.PatientDOB,
		p.PatientAnnualIncome,
		p.PatientId,
	)
	return err
}

func DeletePatient(id int) error {
	_, err := config.DB.Exec("DELETE FROM Patient WHERE patient_id=?", id)
	return err
}
