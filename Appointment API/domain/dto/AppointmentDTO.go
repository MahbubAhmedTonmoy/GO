package dto

import "time"

type AppointmentDTO struct {
	DoctorId          int       `json:"doctorId"`
	DateOfAppointment time.Time `json:"dateOfAppointment"`
	PatientName       string    `json:"patientName"`
	PatientEmail      string    `json:"patientEmail"`
	PatientPhone      string    `json:"patientPhone"`
}
