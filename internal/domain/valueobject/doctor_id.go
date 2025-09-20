package valueobject

import (
	errors "license-service/pkg/log/error"
	logger "license-service/pkg/log/logger"
)

type DoctorID struct {
	value string
}

func NewDoctorID(value string) (*DoctorID, error) {
	err := validateDoctorID(value)
	if err != nil {
		return nil, err
	}
	return &DoctorID{}, nil
}

func validateDoctorID(value string) error {
	if value == "" {
		appError := errors.NewAppError(
			errors.ErrMissingRequiredField,
			"Diagnosis",
			"validateDiagnosis",
			"Diagnosis cannot be empty")
		logger.Error("DoctorID", "validateDoctorID0", appError, "DoctorID", value, "message", "DoctorID validation failed: Is empty")
		return appError
	}
	return nil
}

func (id DoctorID) Value() string {
	return id.value
}

func (id DoctorID) String() string {
	return id.value
}
