package models

import (
	"time"

)

type Patient struct {
    PatientID        uint    `gorm:"primaryKey"`
    PatientName      string  `gorm:"not null"`
    PatientEmail     string  `gorm:"not null"`
    PatientPhone     string  `gorm:"not null"`
    AddressID       *uint   `gorm:"column:address_id"`  // pointer to allow NULL
    PatientDOB       time.Time
    EmploymentStatus string
    AnnualIncome     float64
    InsuranceID      *uint
    Address          Address  `gorm:"foreignKey:AddressID"`
    Insurance        Insurance `gorm:"foreignKey:InsuranceID"`
}

