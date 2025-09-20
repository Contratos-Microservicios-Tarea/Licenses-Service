package contrats

import (
	"context"
	dto "license-service/internal/application/dto"
)

// Para POST /licenses
type LicenseIssuer interface {
	Execute(ctx context.Context, createLicenseDTO dto.CreateLicenseDTO) (*dto.LicenseDTO, error)
}

// Para GET /licenses/{folio}
type LicenseRetriever interface {
	Execute(ctx context.Context, folio string) (*dto.LicenseDTO, error)
}

// Para GET /licenses?patientId={id}
type LicensesByPatientRetriever interface {
	Execute(ctx context.Context, patientID string) ([]*dto.LicenseDTO, error)
}

// Para GET /licenses/{folio}/verify
type LicenseVerifier interface {
	Execute(ctx context.Context, folio string) (bool, error)
}
