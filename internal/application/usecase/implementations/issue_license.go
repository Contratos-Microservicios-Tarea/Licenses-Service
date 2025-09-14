package implementations

import (
	"context"
	"license-service/internal/application/dto"
	"license-service/internal/application/usecase/contrats"
)

type IssueLicenseUseCase struct {
}

func NewIssueLicenseUseCase() contrats.LicenseIssuer {
	return &IssueLicenseUseCase{}
}

// Execute implements contrats.IssueLicenseUseCase.
func (usecase *IssueLicenseUseCase) Execute(ctx context.Context, fermenterDTO dto.CreateLicenseDTO) (dto.LicenseDTO, error) {
	panic("unimplemented")
}
