package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"license-service/internal/application/dto"
	"license-service/internal/application/usecase/contrats"
	handler "license-service/pkg/handler"
	errors "license-service/pkg/log/error"
	logs "license-service/pkg/log/logger"

	"github.com/gorilla/mux"
)

type LicenseController struct {
	issueLicenseUseCase               contrats.LicenseIssuer
	retrieveLicenseUseCase            contrats.LicenseRetriever
	licenseVerifierUseCase            contrats.LicenseVerifier
	licensesByPatientRetrieverUseCase contrats.LicensesByPatientRetriever
	logger                            logs.Logger
}

func NewLicenseController(
	issueLicenseUseCase contrats.LicenseIssuer,
	retrieveLicenseUseCase contrats.LicenseRetriever,
	licenseVerifierUseCase contrats.LicenseVerifier,
	licensesByPatientRetrieverUseCase contrats.LicensesByPatientRetriever,
) *LicenseController {
	return &LicenseController{
		issueLicenseUseCase:               issueLicenseUseCase,
		retrieveLicenseUseCase:            retrieveLicenseUseCase,
		licenseVerifierUseCase:            licenseVerifierUseCase,
		licensesByPatientRetrieverUseCase: licensesByPatientRetrieverUseCase,
		logger:                            *logs.NewLogger(),
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

	ctx := r.Context()
	license, err := lc.issueLicenseUseCase.Execute(ctx, req)
	if err != nil {
		lc.logger.Error("LicenseController", "CreateLicense", err, "use case execution failed")

		lc.logger.Info("LicenseController", "CreateLicense", "Error message: "+err.Error())

		switch {
		case strings.Contains(err.Error(), "Days must be greater than 0"):
			lc.logger.Info("LicenseController", "CreateLicense", "Matched invalid days case")
			handler.WriteErrorResponse(w, http.StatusBadRequest, "INVALID_DAYS")
			return
		case strings.Contains(err.Error(), "invalid date"):
			handler.WriteErrorResponse(w, http.StatusBadRequest, "INVALID_DATE")
			return
		case strings.Contains(err.Error(), "PatientID is required"):
			handler.WriteErrorResponse(w, http.StatusBadRequest, "INVALID_PATIENT")
			return
		case strings.Contains(err.Error(), "DoctorID is required"):
			handler.WriteErrorResponse(w, http.StatusBadRequest, "INVALID_DOCTOR")
			return
		default:
			lc.logger.Info("LicenseController", "CreateLicense", "No specific error match, using default handler")
			handler.HandleUseCaseError(w, err)
			return
		}
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

func (lc *LicenseController) GetLicense(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		lc.logger.Error("LicenseController", "GetLicense", nil, "method not allowed: "+r.Method)
		handler.WriteErrorResponse(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED")
		return
	}

	vars := mux.Vars(r)
	folio := vars["folio"]

	if folio == "" {
		lc.logger.Error("LicenseController", "GetLicense", nil, "folio parameter is missing")
		handler.WriteErrorResponse(w, http.StatusBadRequest, "NOT_FOUND")
		return
	}

	lc.logger.Info("LicenseController", "GetLicense", "retrieving license with folio: "+folio)

	ctx := r.Context()
	license, err := lc.retrieveLicenseUseCase.Execute(ctx, folio)
	if err != nil {
		lc.logger.Error("LicenseController", "GetLicense", err, "use case execution failed")
		handler.HandleUseCaseError(w, err)
		return
	}

	lc.logger.Info("LicenseController", "GetLicense", "license retrieved successfully")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(license); err != nil {
		AppErr := errors.NewAppError(
			errors.ErrInternalError,
			"LicenseController",
			"GetLicense",
			"failed to encode response",
		)
		lc.logger.Error("LicenseController", "GetLicense", AppErr, "response encoding failed")
		handler.WriteErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR")
	}
}

func (controller *LicenseController) VerifyLicense(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	folio := vars["folio"]

	if folio == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Folio is required",
		})
		return
	}

	isValid, err := controller.licenseVerifierUseCase.Execute(r.Context(), folio)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Internal server error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if isValid {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"valid": true})
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]bool{"valid": false})
	}
}

func (lc *LicenseController) GetLicensesByPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		lc.logger.Error("LicenseController", "GetLicensesByPatient", nil, "method not allowed: "+r.Method)
		handler.WriteErrorResponse(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED")
		return
	}

	patientID := r.URL.Query().Get("patientId")
	if patientID == "" {
		lc.logger.Error("LicenseController", "GetLicensesByPatient", nil, "patientId parameter is missing")
		handler.WriteErrorResponse(w, http.StatusBadRequest, "MISSING_REQUIRED_FIELD")
		return
	}

	lc.logger.Info("LicenseController", "GetLicensesByPatient", "retrieving licenses for patient: "+patientID)

	ctx := r.Context()
	licenses, err := lc.licensesByPatientRetrieverUseCase.Execute(ctx, patientID)
	if err != nil {
		lc.logger.Error("LicenseController", "GetLicensesByPatient", err, "use case execution failed")
		handler.HandleUseCaseError(w, err)
		return
	}

	lc.logger.Info("LicenseController", "GetLicensesByPatient", "licenses retrieved successfully", "count", len(licenses))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(licenses); err != nil {
		AppErr := errors.NewAppError(
			errors.ErrInternalError,
			"LicenseController",
			"GetLicensesByPatient",
			"failed to encode response",
		)
		lc.logger.Error("LicenseController", "GetLicensesByPatient", AppErr, "response encoding failed")
		handler.WriteErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR")
	}
}
