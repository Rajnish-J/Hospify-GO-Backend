package models

type Status struct {
	StatusID   uint   `gorm:"primaryKey"`
	EntityType string `gorm:"not null"`
	StatusName string `gorm:"not null"`
}
