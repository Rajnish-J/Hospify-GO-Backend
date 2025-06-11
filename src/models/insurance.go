package models

type Insurance struct {
	InsuranceID     uint `gorm:"primaryKey"`
	ProviderID      *uint `gorm:"foreignKey:ProviderID"`
	PolicyNumber    string
	CoverageDetails string
}
