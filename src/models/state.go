package models

type State struct {
	StateID   uint   `gorm:"primaryKey"`
	StateName string `gorm:"not null"`
}
