package models

type City struct {

	CityID   uint   `gorm:"primaryKey"`
	CityName string `gorm:"not null"`
	StateID  uint   `gorm:"not null"`       // Foreign key to State
	State    State  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Optional: safer foreign key behavior

}

 