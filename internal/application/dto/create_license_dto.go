package dto

import (
	"encoding/json"
	"time"
)

type CreateLicenseDTO struct {
	PatientID string     `json:"patientId" validate:"required"`
	DoctorID  string     `json:"doctorId" validate:"required"`
	Diagnosis string     `json:"diagnosis" validate:"required"`
	StartDate CustomDate `json:"startDate" validate:"required"`
	Days      uint8      `json:"days" validate:"required,gt=0"`
}

type CustomDate struct {
	time.Time
}

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1]

	formats := []string{
		"2006-01-02",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05Z07:00",
		time.RFC3339,
	}

	for _, format := range formats {
		if t, err := time.Parse(format, s); err == nil {
			cd.Time = t
			return nil
		}
	}

	return &time.ParseError{Layout: "2006-01-02", Value: s}
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(cd.Time.Format("2006-01-02"))
}
