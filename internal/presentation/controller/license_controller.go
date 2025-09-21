package controller

import (
	"encoding/json"
	"net/http"

	"license-service/internal/application/dto"
	"license-service/internal/application/usecase/contrats"
	handler "license-service/pkg/handler"
	errors "license-service/pkg/log/error"
	logs "license-service/pkg/log/logger"
)

type LicenseController struct {
	issueLicenseUseCase contrats.LicenseIssuer
	logger              logs.Logger
}

func NewLicenseController(issueLicenseUseCase contrats.LicenseIssuer) *LicenseController {
	return &LicenseController{
		issueLicenseUseCase: issueLicenseUseCase,
		logger:              *logs.NewLogger(),
	}
}

func (lc *LicenseController) CreateLicense(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		lc.logger.Error("LicenseController", "CreateLicense", nil, "method not allowed: "+r.Method)
		handler.WriteErrorResponse(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED")
		return
	}

	var req dto.CreateLicenseDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		AppErr := errors.NewAppError(
			errors.ErrInvalidData,
			"LicenseController",
			"CreateLicense",
			"failed to decode request body",
		)
		lc.logger.Error("LicenseController", "CreateLicense", AppErr, "invalid JSON in request body")
		handler.WriteErrorResponse(w, http.StatusBadRequest, "INVALID_REQUEST")
		return
	}

	lc.logger.Info("LicenseController", "CreateLicense", "request decoded successfully, executing use case")

	ctx := r.Context()
	license, err := lc.issueLicenseUseCase.Execute(ctx, req)
	if err != nil {
		lc.logger.Error("LicenseController", "CreateLicense", err, "use case execution failed")
		handler.HandleUseCaseError(w, err)
		return
	}

	lc.logger.Info("LicenseController", "CreateLicense", "license created successfully")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(license); err != nil {
		AppErr := errors.NewAppError(
			errors.ErrInternalError,
			"LicenseController",
			"CreateLicense",
			"failed to encode response",
		)
		lc.logger.Error("LicenseController", "CreateLicense", AppErr, "response encoding failed")
		handler.WriteErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR")
	}
}
