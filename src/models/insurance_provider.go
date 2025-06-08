package models

type InsuranceProvider struct {
    ProviderID   uint   `gorm:"primaryKey"`
    ProviderName string
}
