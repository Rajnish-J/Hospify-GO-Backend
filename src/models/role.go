package models

type Role struct {
	RoleID   uint   `gorm:"primaryKey"`
	RoleName string `gorm:"unique;not null"`
}
