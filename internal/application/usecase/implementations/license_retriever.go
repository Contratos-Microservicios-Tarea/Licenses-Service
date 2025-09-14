package implementations

import (
	"context"
	"license-service/internal/application/dto"
	"license-service/internal/application/usecase/contrats"
)

type LicenseRetrieverUseCase struct {
}

func NewLicenseRetrieverUseCase() contrats.LicenseRetriever {
	return &LicenseRetrieverUseCase{}
}

// Execute implements contrats.LicenseRetriever.
func (l *LicenseRetrieverUseCase) Execute(ctx context.Context, folio string) (dto.LicenseDTO, error) {
	panic("unimplemented")
}
