package repositories

import (
	"context"
	models "license-service/internal/persistence/entities"
)

type LicenseRepository interface {
	Save(ctx context.Context, license *models.LicenseModel) error
	FindByFolio(ctx context.Context, folio string) (*models.LicenseModel, error)
	FindByPatientID(ctx context.Context, patientID string) ([]*models.LicenseModel, error)
	ExistsByFolioAndStatus(ctx context.Context, folio string, status string) (bool, error)
}
