package service

import (
	"errors"
	"hospify/src/dto"
	"hospify/src/repository"
)

type DoctorService struct {
	repo repository.DoctorRepository
}

func NewDoctorService(repo repository.DoctorRepository) *DoctorService {
	return &DoctorService{repo: repo}
}

func (s *DoctorService) CreateDoctor(d dto.DoctorDTO) error {
	doctor, err := d.ToModel()
	if err != nil {
		return err
	}
	return s.repo.Create(doctor)
}

func (s *DoctorService) GetDoctorByID(id uint) (*dto.DoctorDTO, error) {
	doctor, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return dto.FromDoctorModel(doctor), nil
}

func (s *DoctorService) GetAllDoctors() ([]*dto.DoctorDTO, error) {
	doctors, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	result := make([]*dto.DoctorDTO, 0, len(doctors))
	for _, doc := range doctors {
		result = append(result, dto.FromDoctorModel(doc))
	}
	return result, nil
}

func (s *DoctorService) UpdateDoctor(d dto.DoctorDTO) error {
	if d.ID == 0 {
		return errors.New("doctor ID is required")
	}
	doctor, err := d.ToModel()
	if err != nil {
		return err
	}
	return s.repo.Update(doctor)
}

func (s *DoctorService) DeleteDoctor(id uint) error {
	return s.repo.Delete(id)
}
