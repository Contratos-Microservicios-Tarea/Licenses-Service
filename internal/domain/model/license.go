package domain

import (
	"time"
)

type License struct {
	Folio     string
	Rut       string
	LastName  string
	DoctorID  string
	Diagnosis string
	StartDate time.Time
	Status    string
	Days      uint8
}

func NewLicense(license License) (*License, error) {
	return &License{
		Folio:     license.Folio,
		Rut:       license.Rut,
		LastName:  license.LastName,
		DoctorID:  license.DoctorID,
		Diagnosis: license.Diagnosis,
		StartDate: license.StartDate,
		Status:    license.Status,
		Days:      license.Days,
	}, nil

}
