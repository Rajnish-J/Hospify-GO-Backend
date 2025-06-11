package models

type Address struct {
	AddressID  uint   `gorm:"primaryKey"`
	DoorNo     string
	StreetName string
	CityID     uint   `gorm:"foreignKey:CityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StateID    uint   `gorm:"foreignKey:StateID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ZipCode    string 
}