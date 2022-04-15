package repository

import (
	"AppointmentApi/domain"
	"gorm.io/gorm"
)

var DoctorRepo *DoctorRepository

type DoctorRepository struct {
	*repository[domain.Doctor]
}

func (r *DoctorRepository) Initialize(db *gorm.DB) *DoctorRepository {
	DoctorRepo = &DoctorRepository{
		repository: &repository[domain.Doctor]{db: db},
	}
	return DoctorRepo
}
