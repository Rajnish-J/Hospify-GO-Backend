package models

type City struct {
    CityID   uint   `gorm:"primaryKey"`
    CityName string
    StateID  uint
    State    State `gorm:"foreignKey:StateID"`
}
