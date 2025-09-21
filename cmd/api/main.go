package main

import (
	"license-service/internal/application/usecase/implementations"
	database "license-service/internal/persistence/configuration"
	persistenceRepo "license-service/internal/persistence/repositories"
	env "license-service/pkg/env"
	logs "license-service/pkg/log/logger"
	"net/http"

	"license-service/internal/presentation/router"
)

func main() {
	config := env.Load()
	logger := logs.NewLogger()

	db, err := database.NewConnection()
	if err != nil {
		logger.Error("Main", "main", err, "Failed to connect to database")
		panic(err)
	}

	licenseRepo := persistenceRepo.NewLicenseRepositoryImpl(db)

	licenseIssuer := implementations.NewIssueLicenseUseCase(licenseRepo)
	licenseRetriever := implementations.NewLicenseRetrieverUseCase(licenseRepo)
	licenseVerifier := implementations.NewLicenseVerifierUseCase(licenseRepo)
	licensesByPatientRetriever := implementations.NewLicensesByPatientRetrieverUseCase(licenseRepo)

	router := router.SetupRoutes(licenseIssuer, licenseRetriever, licenseVerifier, licensesByPatientRetriever, *logger)

	port := ":" + config.Server.Port
	logger.Info("Main", "main", "Server starting on port "+config.Server.Port)

	if err := http.ListenAndServe(port, router); err != nil {
		logs.Error("Main", "Failed to start server:", err)
	}

}
