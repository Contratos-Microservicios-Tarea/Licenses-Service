package models

import (
	"time"

	"gorm.io/gorm"
)

type LicenseModel struct {
	ID        uint           `gorm:"primarykey"`
	Folio     string         `gorm:"uniqueIndex;not null;size:50"`
	PatientID string         `gorm:"not null;size:50;index:idx_licenses_patient_id;column:patient_id"`
	DoctorID  string         `gorm:"not null;size:50;column:doctor_id"`
	Diagnosis string         `gorm:"not null;type:text"`
	StartDate time.Time      `gorm:"not null;type:date;column:start_date"`
	Days      int            `gorm:"not null;check:days > 0"`
	Status    string         `gorm:"not null;default:'issued';size:20"`
	CreatedAt time.Time      `gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (LicenseModel) TableName() string {
	return "licenses"
}
