package dto

import "time"

type SearchAvailabilityDTO struct {
	ID                          int       `json:"id"`
	Date                        time.Time `json:"date"`
	DoctorId                    int       `json:"doctorId"`
	DoctorName                  string    `json:"doctorName"`
	AppointmentsToCapacityRatio string    `json:"appointmentsToCapacityRatio"`
	IsAvailable                 bool      `json:"isAvailable"`
}
