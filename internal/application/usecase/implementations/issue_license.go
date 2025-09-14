package implementations

import (
	"context"
	"license-service/internal/application/dto"
	"license-service/internal/application/usecase/contrats"
)

type IssueLicenseUseCase struct {
}

func NewIssueLicenseUseCase() contrats.IssueLicenseUseCase {
	return &IssueLicenseUseCase{}
}

// Execute implements contrats.IssueLicenseUseCase.
func (i *IssueLicenseUseCase) Execute(ctx context.Context, fermenterDTO dto.CreateLicenseDTO) (dto.LicenseDTO, error) {
	panic("unimplemented")
}
