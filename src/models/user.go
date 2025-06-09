package models

type User struct {
	UserID          uint   `gorm:"primaryKey"`
	UserName        string `gorm:"not null"`
	UserEmail       string `gorm:"not null;unique"`
	UserPassword    string `gorm:"not null"`
	RoleId          uint   `gorm:"not null"`
	UserPhoneNumber string `gorm:"not null"`
	IsActive        bool   `gorm:"default:true"`
}