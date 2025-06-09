package models

type Medicine struct {
	MedicineID   uint   `gorm:"primaryKey"`
	MedicineName string `gorm:"not null"`
	Description  string `gorm:"not null"`
	Manufacturer string `gorm:"not null"`
	DosageForm   string `gorm:"not null"`
}
