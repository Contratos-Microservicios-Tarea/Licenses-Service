package valueobject

import (
	errors "license-service/pkg/log/error"
	logger "license-service/pkg/log/logger"
)

type Diagnosis struct {
	value string
}

func NewDiagnosis(value string) (*Diagnosis, error) {
	err := validateDiagnosis(value)
	if err != nil {
		return nil, err
	}
	return &Diagnosis{value: value}, nil
}

func validateDiagnosis(value string) error {
	if value == "" {
		appError := errors.NewAppError(
			errors.ErrMissingRequiredField,
			"Diagnosis",
			"validateDiagnosis",
			"Diagnosis cannot be empty")
		logger.Error("Diagnosis", "validateDiagnosis", appError, "diagnosis", value, "message", "Diagnosis validation failed: Is empty")
		return appError
	}
	return nil
}

func (diagnosis Diagnosis) Value() string {
	return diagnosis.value
}

func (diagnosis Diagnosis) String() string {
	return diagnosis.value
}
