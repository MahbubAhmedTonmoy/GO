package repository

import (
	"log"

	"AppointmentApi/domain"
	"AppointmentApi/domain/dto"

	"gorm.io/gorm"
)

func KeepAutoMigrationUpAndRunning(db *gorm.DB) {
	// the entities from vm package are only used to represent stored proc result sets, these tables are empty
	err := db.AutoMigrate(
		&domain.Doctor{},
		&domain.Appointment{},
		&domain.AppointmentDetails{},
		&dto.SearchAvailabilityDTO{},
		&domain.User{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
