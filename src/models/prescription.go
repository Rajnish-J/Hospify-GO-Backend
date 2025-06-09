package models

type Prescription struct {
	PrescriptionID uint `gorm:"primaryKey"`
	HistoryID      uint `gorm:"not null"`
	MedicineID     uint `gorm:"not null"`
	Dosage         string
	Duration       string
	History        MedicalHistory `gorm:"foreignKey:HistoryID"`
	Medicine       Medicine       `gorm:"foreignKey:MedicineID"`
}
