package dto

import "time"

type CreateLicenseDTO struct {
	PatientID string
	DoctorID  string
	Diagnosis string
	StartDate time.Time
	Days      uint8
}
