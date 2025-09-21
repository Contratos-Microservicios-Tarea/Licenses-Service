package implementations

import (
	"context"
	"license-service/internal/application/dto"
	"license-service/internal/application/usecase/contrats"
	model "license-service/internal/domain/model"
	"license-service/internal/domain/repositories"
	valueobject "license-service/internal/domain/valueobject"
	errorInfo "license-service/pkg/log/error"
	logger "license-service/pkg/log/logger"
)

type IssueLicenseUseCase struct {
	licenseRepository repositories.LicenseRepository
	logger            logger.Logger
}

func NewIssueLicenseUseCase(licenseRepository repositories.LicenseRepository) contrats.LicenseIssuer {
	return &IssueLicenseUseCase{
		licenseRepository: licenseRepository,
		logger:            *logger.NewLogger(),
	}
}

func (usecase *IssueLicenseUseCase) Execute(ctx context.Context, createLicenseDTO dto.CreateLicenseDTO) (*dto.LicenseDTO, error) {

	if err := usecase.validateRequiredFields(createLicenseDTO); err != nil {
		usecase.logger.Error("IssueLicenseUseCase", "Execute", err, "validation failed")
		return nil, err
	}

	patientID, err := valueobject.NewRut(createLicenseDTO.PatientID)
	if err != nil {
		AppErr := errorInfo.NewAppError(
			errorInfo.ErrInvalidData,
			"IssueLicenseUseCase",
			"Execute",
			"invalid PatientID format",
		)
		usecase.logger.Error("IssueLicenseUseCase", "Execute", AppErr, "invalid PatientID")
		return nil, AppErr
	}

	doctorID, err := valueobject.NewDoctorID(createLicenseDTO.DoctorID)
	if err != nil {
		AppErr := errorInfo.NewAppError(
			errorInfo.ErrInvalidData,
			"IssueLicenseUseCase",
			"Execute",
			"invalid DoctorID format",
		)
		usecase.logger.Error("IssueLicenseUseCase", "Execute", AppErr, "invalid DoctorID")
		return nil, AppErr
	}

	diagnosis, err := valueobject.NewDiagnosis(createLicenseDTO.Diagnosis)
	if err != nil {
		AppErr := errorInfo.NewAppError(
			errorInfo.ErrInvalidData,
			"IssueLicenseUseCase",
			"Execute",
			"invalid Diagnosis format",
		)
		usecase.logger.Error("IssueLicenseUseCase", "Execute", AppErr, "invalid Diagnosis")
		return nil, AppErr
	}

	license := model.NewLicense(
		model.License{
			PatientID: patientID.Value(),
			DoctorID:  doctorID.Value(),
			Diagnosis: diagnosis.Value(),
			Status:    model.StatusIssued,
			StartDate: createLicenseDTO.StartDate.Time,
			Days:      createLicenseDTO.Days,
		},
	)
	license.GenerateFolio()

	if err := license.IsValid(); err != nil {
		AppErr := errorInfo.NewAppError(
			errorInfo.ErrInvalidData,
			"IssueLicenseUseCase",
			"Execute",
			"license is not valid",
		)
		usecase.logger.Error("IssueLicenseUseCase", "Execute", AppErr, err.Error())
		return nil, AppErr
	}

	if err := usecase.licenseRepository.Save(ctx, license); err != nil {
		usecase.logger.Error("IssueLicenseUseCase", "Execute", err, "failed to save license")
		return nil, err
	}

	responseDTO := &dto.LicenseDTO{
		Folio:     license.Folio,
		PatientID: license.PatientID,
		DoctorID:  license.DoctorID,
		Diagnosis: license.Diagnosis,
		StartDate: license.StartDate.String(),
		Days:      license.Days,
		Status:    license.Status,
	}

	usecase.logger.Info("IssueLicenseUseCase", "Execute", "license created successfully")
	return responseDTO, nil
}

func (usecase *IssueLicenseUseCase) validateRequiredFields(dto dto.CreateLicenseDTO) error {
	if dto.PatientID == "" {
		return errorInfo.NewAppError(
			errorInfo.ErrMissingRequiredField,
			"IssueLicenseUseCase",
			"validateRequiredFields",
			"PatientID is required",
		)
	}

	if dto.DoctorID == "" {
		return errorInfo.NewAppError(
			errorInfo.ErrMissingRequiredField,
			"IssueLicenseUseCase",
			"validateRequiredFields",
			"DoctorID is required",
		)
	}

	if dto.Diagnosis == "" {
		return errorInfo.NewAppError(
			errorInfo.ErrMissingRequiredField,
			"IssueLicenseUseCase",
			"validateRequiredFields",
			"Diagnosis is required",
		)
	}

	if dto.Days <= 0 {
		return errorInfo.NewAppError(
			errorInfo.ErrInvalidData,
			"IssueLicenseUseCase",
			"validateRequiredFields",
			"Days must be greater than 0",
		)
	}

	return nil
}
