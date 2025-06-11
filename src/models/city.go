package models

type City struct {
	CityID   uint   `gorm:"primaryKey"`
	CityName string `gorm:"not null"`
	StateID  uint   `gorm:"foreignKey:StateID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	} 