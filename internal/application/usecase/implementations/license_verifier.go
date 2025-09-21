package implementations

import (
	"context"
	"license-service/internal/application/usecase/contrats"
	"license-service/internal/domain/repositories"
	errorInfo "license-service/pkg/log/error"
	logger "license-service/pkg/log/logger"
)

type LicenseVerifierUseCase struct {
	licenseRepository repositories.LicenseRepository
	logger            logger.Logger
}

func NewLicenseVerifierUseCase(licenseRepository repositories.LicenseRepository) contrats.LicenseVerifier {
	return &LicenseVerifierUseCase{
		licenseRepository: licenseRepository,
		logger:            *logger.NewLogger(),
	}
}

func (usecase *LicenseVerifierUseCase) Execute(ctx context.Context, folio string) (bool, error) {
	usecase.logger.Info(
		"LicenseVerifierUseCase",
		"Execute",
		"starting license verification process",
	)

	if err := usecase.validateFolio(folio); err != nil {
		usecase.logger.Error(
			"LicenseVerifierUseCase",
			"Execute",
			err,
			"validation failed",
		)
		return false, err
	}

	license, err := usecase.licenseRepository.FindByFolio(ctx, folio)
	if err != nil {
		usecase.logger.Error(
			"LicenseVerifierUseCase",
			"Execute",
			err,
			"failed to retrieve license from repository",
		)
		return false, err
	}

	if license == nil {
		usecase.logger.Info(
			"LicenseVerifierUseCase",
			"Execute",
			"license not found",
		)
		return false, nil
	}

	isIssued := license.IsIssued()

	if isIssued {
		usecase.logger.Info(
			"LicenseVerifierUseCase",
			"Execute",
			"license verification successful - valid license",
		)
	} else {
		usecase.logger.Info(
			"LicenseVerifierUseCase",
			"Execute",
			"license found but not issued",
		)
	}
	return isIssued, nil
}

func (usecase *LicenseVerifierUseCase) validateFolio(folio string) error {
	if folio == "" {
		return errorInfo.NewAppError(
			errorInfo.ErrMissingRequiredField,
			"LicenseVerifierUseCase",
			"validateFolio",
			"folio is required",
		)
	}

	return nil
}
