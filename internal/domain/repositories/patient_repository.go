package repositories

import (
	"context"
)

type PatientRepository interface {
	ExistsByPatientID(ctx context.Context, patientID string) (bool, error)
}
