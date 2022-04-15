package repository

import (
	"gorm.io/gorm"

	"AppointmentApi/domain"
)

var AccountRepo *AccountRepository

type AccountRepository struct {
	*repository[domain.User]
}

func (r *AccountRepository) Initialize(db *gorm.DB) *AccountRepository {
	AccountRepo = &AccountRepository{
		repository: &repository[domain.User]{db: db},
	}
	return AccountRepo
}
