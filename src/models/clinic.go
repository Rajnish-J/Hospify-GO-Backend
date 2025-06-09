package models

type Clinic struct {
	ClinicID   uint   `gorm:"primaryKey"`
	ClinicName string `gorm:"not null"`
	Address    string `gorm:"not null"`
	Phone      string `gorm:"not null"`
}
