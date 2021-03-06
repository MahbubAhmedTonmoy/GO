package service

import (
	"AppointmentApi/constants"
	"AppointmentApi/domain"
	"AppointmentApi/domain/dto"
	"AppointmentApi/helpers"
	"AppointmentApi/pkg"
	"AppointmentApi/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var DoctorService *doctorService

type doctorService struct {
	*base
	repo *repository.DoctorRepository
}

func (s *doctorService) Initialize(b *base, r *repository.DoctorRepository) {
	DoctorService = &doctorService{
		base: b,
		repo: r,
	}
}

func (s *doctorService) GetAllDoctors() ([]domain.Doctor, int, string) {
	var doctors []domain.Doctor

	val, isRetrieved, err := pkg.RetrieveValueFromRedisClient(s.rdb, s.ctx, "doctors")
	if err != nil {
		return nil, http.StatusInternalServerError, ""
	}

	if isRetrieved { // cached data got from redis server
		fmt.Println("redis server: data retrieved successfully")
		_ = json.Unmarshal([]byte(val), &doctors)
	} else { // data not found in cache
		s.repo.FindAll(&doctors, "")
	}

	code := http.StatusOK
	message := ""
	if len(doctors) == 0 {
		code = http.StatusNotFound
		message = "doctors list empty"
	}

	_ = pkg.InsertKeyValuePairInRedisClient(s.rdb, s.ctx, "doctors", doctors)

	return doctors, code, message

}

func (s *doctorService) GetDoctorById(id int) (domain.Doctor, int, string) {
	var doctor domain.Doctor

	key := "doctor/" + string(id)
	val, isRetrieved, _ := pkg.RetrieveValueFromRedisClient(s.rdb, s.ctx, key)

	if isRetrieved {
		_ = json.Unmarshal([]byte(val), &doctor)
	} else {
		s.repo.FindById(&doctor, id)
	}

	if doctor.ID > 0 { // means valid
		_ = pkg.InsertKeyValuePairInRedisClient(s.rdb, s.ctx, key, doctor)
		return doctor, http.StatusOK, ""
	} else {
		return doctor, http.StatusNotFound, fmt.Sprintf(constants.ApplicationMessage.ItemNotFoundByTheGivenPrimaryKeyOfItem, "Doctor", "doctor")
	}

}

func (s *doctorService) GetDoctorsByAvailability(fromDate time.Time, toDate time.Time) ([]dto.SearchAvailabilityDTO, int, string) {
	var schedules []dto.SearchAvailabilityDTO

	err := s.repo.ExecuteRawSqlAndScan(&schedules, "EXEC SearchAvailability ?, ?", fromDate, toDate)
	if err != nil {
		return nil, http.StatusBadRequest, err.Error()
	}

	return schedules, http.StatusOK, ""
}

func (s *doctorService) RequestAppointmentToDoctor(apnt dto.AppointmentDTO) (int, string) {
	var message string = "command executed successfully"
	err := s.repo.ExecuteRawSqlAndScan(&message, "EXEC CreateAppointmentToDoctor ?, ?, ?, ?, ?", apnt.DoctorId, apnt.DateOfAppointment, apnt.PatientName, apnt.PatientEmail, apnt.PatientPhone)
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	var doctor domain.Doctor
	s.repo.FindById(&doctor, apnt.DoctorId)

	// sending mails in a new go routine...
	go s.sendMailToDoctorAndPatient(apnt, doctor)

	return http.StatusAccepted, "Accepted"
}

func (h *doctorService) sendMailToDoctorAndPatient(apnt dto.AppointmentDTO, doctor domain.Doctor) {
	// sending mail to patient
	mailToPatient := domain.MailData{
		To:      apnt.PatientEmail,
		Subject: helpers.GetMailSubjectForPatient(doctor.Name),
		Content: helpers.GetMailBodyForPatient(apnt, doctor.Name),
	}
	h.mailChan <- mailToPatient

	// sending mail to doctor
	mailToDoctor := domain.MailData{
		To:      doctor.Email,
		Subject: helpers.GetMailSubjectForDoctor(apnt.DateOfAppointment),
		Content: helpers.GetMailBodyForDoctor(apnt),
	}
	h.mailChan <- mailToDoctor
}
