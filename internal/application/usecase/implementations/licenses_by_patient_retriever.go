// LicensesByPatientRetriever

package implementations

import (
	"context"
	"license-service/internal/application/dto"
	"license-service/internal/application/usecase/contrats"
)

type LicensesByPatientRetrieverUseCase struct {
}

func NewLicensesByPatientRetrieverUseCase() contrats.LicensesByPatientRetriever {
	return &LicensesByPatientRetrieverUseCase{}
}

// Execute implements contrats.LicensesByPatientRetriever.
func (usecase *LicensesByPatientRetrieverUseCase) Execute(ctx context.Context, patientID string) ([]*dto.LicenseDTO, error) {
	panic("unimplemented")
}
