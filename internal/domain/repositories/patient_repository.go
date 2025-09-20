package repositories

import (
	"context"
	model "license-service/internal/domain/model"
)

type PatientRepository interface {
	ExistsByPatientID(ctx context.Context, patientID string) (bool, error)
	SavePatient(ctx context.Context, patient model.Patient) error
}
