package router

import (
	"license-service/internal/application/usecase/contrats"
	"license-service/internal/presentation/controller"
	logs "license-service/pkg/log/logger"

	"github.com/gorilla/mux"
)

func SetupRoutes(
	licenseIssuer contrats.LicenseIssuer,
	licenseRetriever contrats.LicenseRetriever,
	licenseVerifier contrats.LicenseVerifier,
	logger logs.Logger,
) *mux.Router {
	router := mux.NewRouter()

	licenseController := controller.NewLicenseController(
		licenseIssuer,
		licenseRetriever,
		licenseVerifier,
	)

	router.HandleFunc("/licenses", licenseController.CreateLicense).Methods("POST")
	router.HandleFunc("/licenses/{folio}", licenseController.GetLicense).Methods("GET")
	router.HandleFunc("/licenses/{folio}/verify", licenseController.VerifyLicense).Methods("GET")

	return router
}
