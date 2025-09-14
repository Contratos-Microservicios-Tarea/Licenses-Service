package implementations

import (
	"context"
	"license-service/internal/application/usecase/contrats"
)

type LicenseVerifierUseCase struct {
}

func NewLicenseVerifierUseCase() contrats.LicenseVerifier {
	return &LicenseVerifierUseCase{}
}

// Execute implements contrats.LicenseVerifier.
func (usecase *LicenseVerifierUseCase) Execute(ctx context.Context, folio string) (bool, error) {
	panic("unimplemented")
}
