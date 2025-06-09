package models

type Medicine struct {
    MedicineID    uint   `gorm:"primaryKey"`
    MedicineName  string
    Description   string
    Manufacturer  string
    DosageForm    string
}
