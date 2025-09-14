package contrats

import (
	"context"
	dto "license-service/internal/application/dto"
)

type IssueLicenseUseCase interface {
	Execute(ctx context.Context, fermenterDTO dto.CreateLicenseDTO) (dto.LicenseDTO, error)
}
