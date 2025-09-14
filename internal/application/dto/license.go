package dto

import "time"

type CreateLicenseDTO struct {
	PatientID string
	DoctorID  string
	Diagnosis string
	StartDate time.Time
	Days      uint8
}

type LicenseDTO struct {
	Folio     string
	PatientID string
	DoctorID  string
	Diagnosis string
	StartDate string
	Days      uint8
	Status    string
}
