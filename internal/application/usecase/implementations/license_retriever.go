package implementations

import (
	"context"
	"license-service/internal/application/dto"
	"license-service/internal/application/usecase/contrats"
	"license-service/internal/domain/repositories"
	errorInfo "license-service/pkg/log/error"
	logger "license-service/pkg/log/logger"
)

type LicenseRetrieverUseCase struct {
	licenseRepository repositories.LicenseRepository
	logger            logger.Logger
}

func NewLicenseRetrieverUseCase(licenseRepository repositories.LicenseRepository) contrats.LicenseRetriever {
	return &LicenseRetrieverUseCase{
		licenseRepository: licenseRepository,
		logger:            *logger.NewLogger(),
	}
}

func (usecase *LicenseRetrieverUseCase) Execute(ctx context.Context, folio string) (*dto.LicenseDTO, error) {
	usecase.logger.Info("LicenseRetrieverUseCase", "Execute", "retrieving license with folio: "+folio)

	if folio == "" {
		appErr := errorInfo.NewAppError(
			errorInfo.ErrMissingRequiredField,
			"LicenseRetrieverUseCase",
			"Execute",
			"folio is required",
		)
		usecase.logger.Error("LicenseRetrieverUseCase", "Execute", appErr, "empty folio provided")
		return nil, appErr
	}

	license, err := usecase.licenseRepository.FindByFolio(ctx, folio)
	if err != nil {
		usecase.logger.Error("LicenseRetrieverUseCase", "Execute", err, "failed to retrieve license from repository")
		return nil, err
	}

	if license == nil {
		appErr := errorInfo.NewAppError(
			errorInfo.ErrNotFound,
			"LicenseRetrieverUseCase",
			"Execute",
			"license not found",
		)
		usecase.logger.Info("LicenseRetrieverUseCase", "Execute", "license not found for folio: "+folio)
		return nil, appErr
	}

	responseDTO := &dto.LicenseDTO{
		Folio:     license.Folio,
		PatientID: license.PatientID,
		DoctorID:  license.DoctorID,
		Diagnosis: license.Diagnosis,
		StartDate: license.StartDate.Format("2006-01-02"),
		Days:      license.Days,
		Status:    license.Status,
	}

	usecase.logger.Info("LicenseRetrieverUseCase", "Execute", "license retrieved successfully for folio: "+folio)
	return responseDTO, nil
}
