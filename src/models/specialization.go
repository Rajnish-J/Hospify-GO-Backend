package models

type Specialization struct {
	SpecializationID   uint   `gorm:"primaryKey"`
	SpecializationName string `gorm:"not null"`
}
