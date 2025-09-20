package models

import (
	"time"

	"gorm.io/gorm"
)

type PatientModel struct {
	ID        uint           `gorm:"primarykey"`
	PatientID string         `gorm:"uniqueIndex;not null;size:50;column:patient_id"`
	FirstName string         `gorm:"size:100;column:first_name"`
	LastName  string         `gorm:"size:100;column:last_name"`
	CreatedAt time.Time      `gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (PatientModel) TableName() string {
	return "patients"
}
