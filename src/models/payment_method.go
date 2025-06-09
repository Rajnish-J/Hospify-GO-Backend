package models

type PaymentMethod struct {
	PaymentMethodID uint   `gorm:"primarykey"`
	MethodName      string `gorm:"not null"`
}