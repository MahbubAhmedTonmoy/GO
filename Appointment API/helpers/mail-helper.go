package helpers

import (
	"AppointmentApi/domain/dto"
	"AppointmentApi/pkg"
	"fmt"
	"time"
)

func GetMailBodyForDoctor(apnt dto.AppointmentDTO) string {

	dateString := pkg.GetReadableDateStringFromDate(apnt.DateOfAppointment)
	mailBody := fmt.Sprintf(`<body> 
		<p> Hello Doctor, <p> 
		<p> A new patient has made an appointment to you on %s. 
		Please be there on time. </p> 
		
		<h4> Patient Details </h4> 
		<ul> 
			<li>Patient's Name: %s</li> 
			<li>Email: %s</li> 
			<li>Phone: %s</li> 
		</ul> 
		
		<i> Sincerely, <br> 
		<strong> Good Health Clinic </strong> <br> 
		Mirpur Branch Co. Ltd. </i> <br> <br> 
	</body>`, dateString, apnt.PatientName, apnt.PatientEmail, apnt.PatientPhone)

	return mailBody
}

func GetMailBodyForPatient(apnt dto.AppointmentDTO, doctorName string) string {

	dateString := pkg.GetReadableDateStringFromDate(apnt.DateOfAppointment)
	mailBody := fmt.Sprintf(`
		<body> <p> Dear Patient, <p> 
		<p> Your request for an appointment with %s 
		has been successfully granted! We hope good health for you. 
		Please be there on time. </p> 
		
		<h4> Submitted Info: </h4> 
		<ul> 
			<li>Patient's Name: %s</li> 
			<li>Email: %s</li> 
			<li>Phone: %s</li> 
			<li> Date Of Appointment: %s </li> 
		</ul> 
		
		<i> Sincerely, <br> 
		<strong> Good Health Clinic </strong> <br> 
		Mirpur Branch Co. Ltd. </i> <br> <br> 
	</body>`, doctorName, apnt.PatientName, apnt.PatientEmail, apnt.PatientPhone, dateString)

	return mailBody
}

func GetMailSubjectForPatient(doctorName string) string {
	return fmt.Sprintf("Your request for appointment with %s has been successfully granted", doctorName)
}

func GetMailSubjectForDoctor(date time.Time) string {
	return fmt.Sprintf("You have a new patient on %s", pkg.GetReadableDateStringFromDate(date))
}
