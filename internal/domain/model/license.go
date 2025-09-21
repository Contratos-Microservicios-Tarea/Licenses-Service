package domain

import (
	"fmt"
	err "license-service/pkg/log/error"
	logger "license-service/pkg/log/logger"
	"time"
)

const (
	StatusIssued  = "issued"
	StatusExpired = "expired"
	StatusRevoked = "revoked"
)

type License struct {
	Folio     string
	PatientID string
	DoctorID  string
	Diagnosis string
	StartDate time.Time
	Status    string
	Days      uint8
}

func NewLicense(license License) *License {
	return &License{
		Folio:     license.Folio,
		PatientID: license.PatientID,
		DoctorID:  license.DoctorID,
		Diagnosis: license.Diagnosis,
		StartDate: license.StartDate,
		Status:    license.Status,
		Days:      license.Days,
	}
}

func ValidateDays(days uint8) error {
	if days <= 0 {
		AppError := err.NewAppError(err.ErrInternalError, "license model", "validateDays", "DAy invalide")
		logger.Error("IssueLicenseUseCase", "Execute", AppError)
		return AppError
	}
	return nil
}

func (license *License) GenerateFolio() {
	license.Folio = fmt.Sprintf("L-%d", time.Now().Unix())
}

func (license *License) SetDefaultStatus(status string) {
	license.Status = StatusIssued
}

func (license *License) IsValid() error {
	if license.Days <= 0 {
		AppError := err.NewAppError(err.ErrInternalError, "license model", "validateDays", "DAy invalide")
		logger.Error("IssueLicenseUseCase", "Execute", AppError)
		return AppError
	}
	if license.Status != StatusIssued {
		AppError := err.NewAppError(err.ErrInternalError, "license model", "validateDays", "INVALIDED STATUS")
		logger.Error("IssueLicenseUseCase", "Execute", AppError)
		return AppError
	}
	return nil
}

func (license *License) IsIssued() bool {
	return license.Status == "issued"
}
