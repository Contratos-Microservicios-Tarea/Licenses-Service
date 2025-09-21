package main

import (
	"fmt"
	"license-service/internal/application/usecase/implementations"
	database "license-service/internal/persistence/configuration"
	persistenceRepo "license-service/internal/persistence/repositories"
	env "license-service/pkg/env"
	logs "license-service/pkg/log/logger"
	"net/http"
	"time"

	"license-service/internal/presentation/router"

	"gorm.io/gorm"
)

func main() {
	config := env.Load()
	logger := logs.NewLogger()

	db, err := connectWithRetry(logger)
	if err != nil {
		logger.Error("Main", "main", err, "Failed to connect to database")
		panic(err)
	}

	licenseRepo := persistenceRepo.NewLicenseRepositoryImpl(db.(*gorm.DB))

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

func connectWithRetry(logger *logs.Logger) (interface{}, error) {
	maxRetries := 3
	retryDelay := 3 * time.Second

	for i := 0; i < maxRetries; i++ {
		logger.Info("Main", "connectWithRetry", fmt.Sprintf("Database connection attempt %d/%d", i+1, maxRetries))

		db, err := database.NewConnection()
		if err == nil {
			logger.Info("Main", "connectWithRetry", "Database connection successful")
			return db, nil
		}

		logger.Error("Main", "connectWithRetry", err, fmt.Sprintf("Connection attempt %d failed", i+1))

		if i < maxRetries-1 {
			logger.Info("Main", "connectWithRetry", fmt.Sprintf("Retrying in %v...", retryDelay))
			time.Sleep(retryDelay)
		}
	}

	return nil, fmt.Errorf("failed to connect to database after %d attempts", maxRetries)
}
