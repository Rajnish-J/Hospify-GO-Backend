package models

type Clinic struct {
    ClinicID   uint   `gorm:"primaryKey"`
    ClinicName string
    Address    string
    Phone      string
}
