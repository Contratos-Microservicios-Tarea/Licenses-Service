package implementations

import (
	"context"
	"license-service/internal/application/dto"
	"license-service/internal/application/usecase/contrats"
	model "license-service/internal/domain/model"
	valueobject "license-service/internal/domain/valueobject"
	errorInfo "license-service/pkg/log/error"
	logger "license-service/pkg/log/logger"
)

type IssueLicenseUseCase struct {
}

func NewIssueLicenseUseCase() contrats.LicenseIssuer {
	return &IssueLicenseUseCase{}
}

func (usecase *IssueLicenseUseCase) Execute(ctx context.Context, createLicenseDTO dto.CreateLicenseDTO) (*dto.LicenseDTO, error) {
	patientID, err := valueobject.NewRut(createLicenseDTO.PatientID)
	if err != nil {
		AppErr := errorInfo.NewAppError(
			errorInfo.ErrMissingRequiredField,
			"IssueLicenseUseCase",
			"Execute",
			"requiere PatientID field",
		)
		logger.Error("IssueLicenseUseCase", "Execute", AppErr, "requiere PatientID field")
		return nil, AppErr
	}

	doctorID, err := valueobject.NewDoctorID(createLicenseDTO.DoctorID)
	if err != nil {
		AppErr := errorInfo.NewAppError(
			errorInfo.ErrMissingRequiredField,
			"IssueLicenseUseCase",
			"Execute",
			"requiere doctorID field",
		)
		logger.Error("IssueLicenseUseCase", "Execute", AppErr, "requiere doctorID field")
		return nil, AppErr
	}

	diagnosis, err := valueobject.NewDiagnosis(createLicenseDTO.Diagnosis)
	if err != nil {
		AppErr := errorInfo.NewAppError(
			errorInfo.ErrMissingRequiredField,
			"IssueLicenseUseCase",
			"Execute",
			"requiere diagnosis field",
		)
		logger.Error("IssueLicenseUseCase", "Execute", AppErr, "requiere diagnosis field")
		return nil, AppErr
	}

	license := model.NewLicense(
		model.License{
			PatientID: patientID.Value(),
			DoctorID:  doctorID.Value(),
			Diagnosis: diagnosis.Value(),
			Status:    model.StatusIssued,
			StartDate: createLicenseDTO.StartDate,
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
		logger.Error("IssueLicenseUseCase", "Execute", AppErr, err.Error())
		return nil, AppErr
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

	return responseDTO, nil
}
