package models

type Insurance struct {
    InsuranceID     uint   `gorm:"primaryKey"`
    ProviderID      *uint
    PolicyNumber    string
    CoverageDetails string
    Provider        InsuranceProvider `gorm:"foreignKey:ProviderID"`
}
