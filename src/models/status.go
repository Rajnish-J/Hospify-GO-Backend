package models

type Status struct {
    StatusID   uint   `gorm:"primaryKey"`
    EntityType string
    StatusName string
}
