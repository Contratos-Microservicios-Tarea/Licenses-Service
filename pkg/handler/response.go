package handler

import (
	"encoding/json"
	errors "license-service/pkg/log/error"
	logger "license-service/pkg/log/logger"
	"net/http"
	"time"
)

func WriteErrorResponse(w http.ResponseWriter, statusCode int, errorCode string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"error":     errorCode,
		"timestamp": time.Now().Format(time.RFC3339),
		"status":    statusCode,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("LicenseController", "writeErrorResponse", err, "failed to encode error response")
		w.Write([]byte(`{"error":"ENCODING_ERROR"}`))
	}
}

func WriteDetailedErrorResponse(w http.ResponseWriter, statusCode int, errorCode string, details string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"error":     errorCode,
		"details":   details,
		"timestamp": time.Now().Format(time.RFC3339),
		"status":    statusCode,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("LicenseController", "writeErrorResponse", err, "failed to encode error response")
		w.Write([]byte(`{"error":"ENCODING_ERROR"}`))
	}
}

func HandleUseCaseError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*errors.AppError); ok {
		switch appErr.Code {
		case errors.ErrMissingRequiredField:
			WriteDetailedErrorResponse(w, http.StatusBadRequest, "MISSING_REQUIRED_FIELD", appErr.Message)
		case errors.ErrInvalidData:
			WriteDetailedErrorResponse(w, http.StatusBadRequest, "INVALID_DATA", appErr.Message)
		default:
			WriteDetailedErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "An unexpected error occurred")
		}
	} else {
		WriteErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR")
	}
}
