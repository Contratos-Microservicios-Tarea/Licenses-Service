package models

import (
	domain "license-service/internal/domain/model"
	"time"
)

type LicenseEntity struct {
	ID        uint      `gorm:"primarykey"`
	Folio     string    `gorm:"uniqueIndex;not null;size:50"`
	PatientID string    `gorm:"not null;size:50;index:idx_licenses_patient_id;column:patient_id"`
	DoctorID  string    `gorm:"not null;size:50;column:doctor_id"`
	Diagnosis string    `gorm:"not null;type:text"`
	StartDate time.Time `gorm:"not null;type:date;column:start_date"`
	Days      int       `gorm:"not null;check:days > 0"`
	Status    string    `gorm:"not null;default:'issued';size:20"`
	CreatedAt time.Time `gorm:"default:now()"`
}

func (LicenseEntity) TableName() string {
	return "licenses"
}

func (e *LicenseEntity) ToDomain() *domain.License {
	return &domain.License{
		Folio:     e.Folio,
		PatientID: e.PatientID,
		DoctorID:  e.DoctorID,
		Diagnosis: e.Diagnosis,
		StartDate: e.StartDate,
		Days:      uint8(e.Days),
		Status:    e.Status,
	}
}

func FromDomain(license *domain.License) *LicenseEntity {
	return &LicenseEntity{
		Folio:     license.Folio,
		PatientID: license.PatientID,
		DoctorID:  license.DoctorID,
		Diagnosis: license.Diagnosis,
		StartDate: license.StartDate,
		Days:      int(license.Days),
		Status:    string(license.Status),
		CreatedAt: time.Now(),
	}
}

func (e *LicenseEntity) UpdateFromDomain(license *domain.License) {
	e.Folio = license.Folio
	e.PatientID = license.PatientID
	e.DoctorID = license.DoctorID
	e.Diagnosis = license.Diagnosis
	e.StartDate = license.StartDate
	e.Days = int(license.Days)
	e.Status = string(license.Status)
}
