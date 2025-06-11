package models

type Prescription struct {
	PrescriptionID uint `gorm:"primaryKey"`
	HistoryID      uint `gorm:"foreignKey:HistoryID"`
	MedicineID     uint `gorm:"foreignKey:MedicineID"`
	Dosage         string
	Duration       string   
}
