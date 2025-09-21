package implementations

import (
	"context"
	"license-service/internal/application/dto"
	"license-service/internal/application/usecase/contrats"
	"license-service/internal/domain/repositories"
	errorInfo "license-service/pkg/log/error"
	logger "license-service/pkg/log/logger"
)

type LicensesByPatientRetrieverUseCase struct {
	licenseRepository repositories.LicenseRepository
	logger            logger.Logger
}

func NewLicensesByPatientRetrieverUseCase(licenseRepository repositories.LicenseRepository) contrats.LicensesByPatientRetriever {
	return &LicensesByPatientRetrieverUseCase{
		licenseRepository: licenseRepository,
		logger:            *logger.NewLogger(),
	}
}

func (usecase *LicensesByPatientRetrieverUseCase) Execute(ctx context.Context, patientID string) ([]*dto.LicenseDTO, error) {
	usecase.logger.Info("LicensesByPatientRetrieverUseCase", "Execute", "retrieving licenses for patient: "+patientID)

	if patientID == "" {
		appErr := errorInfo.NewAppError(
			errorInfo.ErrMissingRequiredField,
			"LicensesByPatientRetrieverUseCase",
			"Execute",
			"patientID is required",
		)
		usecase.logger.Error("LicensesByPatientRetrieverUseCase", "Execute", appErr, "empty patientID provided")
		return nil, appErr
	}

	licenses, err := usecase.licenseRepository.FindByPatientID(ctx, patientID)
	if err != nil {
		usecase.logger.Error("LicensesByPatientRetrieverUseCase", "Execute", err, "failed to retrieve licenses from repository")
		return nil, err
	}

	licenseDTOs := make([]*dto.LicenseDTO, 0, len(licenses))
	for _, license := range licenses {
		licenseDTO := &dto.LicenseDTO{
			Folio:     license.Folio,
			PatientID: license.PatientID,
			DoctorID:  license.DoctorID,
			Diagnosis: license.Diagnosis,
			StartDate: license.StartDate.Format("2006-01-02"),
			Days:      license.Days,
			Status:    license.Status,
		}
		licenseDTOs = append(licenseDTOs, licenseDTO)
	}

	usecase.logger.Info("LicensesByPatientRetrieverUseCase", "Execute", "retrieved licenses successfully for patient: "+patientID, "count", len(licenseDTOs))
	return licenseDTOs, nil
}
