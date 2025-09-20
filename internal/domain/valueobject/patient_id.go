package valueobject

import (
	"regexp"
	"strings"

	errors "license-service/pkg/log/error"
	logger "license-service/pkg/log/logger"
)

type Rut struct {
	value string
}

func NewRut(value string) (*Rut, error) {
	if err := validateRutFormat(value); err != nil {
		return nil, err
	}
	return &Rut{value: value}, nil
}

func validateRutFormat(rut string) error {
	rutRegex := regexp.MustCompile(`^\d{7,8}-[\dkK]$`)

	if strings.TrimSpace(rut) == "" {
		appError := errors.NewAppError(
			errors.ErrValidationFailed,
			"Rut",
			"validateRutFormat",
			"RUT cannot be empty")
		logger.Error("Rut", "validateRutFormat", appError, "rut", rut, "message", "RUT validation failed: empty or whitespace-only RUT provided")
		return appError
	}

	if !rutRegex.MatchString(rut) {
		appError := errors.NewAppError(
			errors.ErrValidationFailed,
			"Rut",
			"validateRutFormat",
			"RUT format must be XXXXXXX-X")
		logger.Error("Rut", "validateRutFormat", appError, "rut", rut, "expected_format", "XXXXXXX-X or XXXXXXXX-X", "message", "RUT validation failed: format does not match Chilean RUT pattern")
		return appError
	}
	return nil
}

func (r Rut) Value() string {
	return r.value
}

func (r Rut) String() string {
	return r.value
}
