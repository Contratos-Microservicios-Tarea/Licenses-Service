package repositories

import (
	"context"
	models "license-service/internal/domain/model"
)

type LicenseRepository interface {
	Save(ctx context.Context, license *models.License) error
	FindByFolio(ctx context.Context, folio string) (*models.License, error)
	FindByPatientID(ctx context.Context, patientID string) ([]*models.License, error)
	ExistsByFolioAndStatus(ctx context.Context, folio string, status string) (bool, error)
}
