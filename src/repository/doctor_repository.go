package repository

import (
	"hospify/src/models"
	"gorm.io/gorm"
)

type DoctorRepository interface {
	Create(doctor *models.Doctor) error
	GetByID(id uint) (*models.Doctor, error)
	GetAll() ([]*models.Doctor, error)
	Update(doctor *models.Doctor) error
	Delete(id uint) error
}

type doctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) DoctorRepository {
	return &doctorRepository{db: db}
}

func (r *doctorRepository) Create(doctor *models.Doctor) error {
	return r.db.Create(doctor).Error
}

func (r *doctorRepository) GetByID(id uint) (*models.Doctor, error) {
	var doctor models.Doctor
	if err := r.db.First(&doctor, id).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *doctorRepository) GetAll() ([]*models.Doctor, error) {
	var doctors []*models.Doctor
	if err := r.db.Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}

func (r *doctorRepository) Update(doctor *models.Doctor) error {
	return r.db.Save(doctor).Error
}

func (r *doctorRepository) Delete(id uint) error {
	return r.db.Delete(&models.Doctor{}, id).Error
}
