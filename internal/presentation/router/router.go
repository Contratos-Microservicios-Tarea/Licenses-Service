package router

import (
	"github.com/gorilla/mux"

	contrats "license-service/internal/application/usecase/contrats"
	"license-service/internal/presentation/controller"
	logs "license-service/pkg/log/logger"
)

func SetupRoutes(
	licenseIssuer contrats.LicenseIssuer,
	logger *logs.Logger,
) *mux.Router {
	licenseController := controller.NewLicenseController(licenseIssuer)

	router := mux.NewRouter()
	router.HandleFunc("/licenses", licenseController.CreateLicense).Methods("POST")

	return router
}
