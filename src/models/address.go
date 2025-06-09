package models

type Address struct {
	AddressID  uint `gorm:"primaryKey"`
	DoorNo     string
	StreetName string
	CityID     uint
	StateID    uint
	ZipCode    string
	City       City  `gorm:"foreignKey:CityID"`
	State      State `gorm:"foreignKey:StateID"`
}
