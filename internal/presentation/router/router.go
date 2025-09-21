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
	logger logs.Logger,
) *mux.Router {
	r := mux.NewRouter()

	licenseController := controller.NewLicenseController(
		licenseIssuer,
		licenseRetriever,
	)

	r.HandleFunc("/licenses", licenseController.CreateLicense).Methods("POST")
	r.HandleFunc("/licenses/{folio}", licenseController.GetLicense).Methods("GET")

	return r
}
