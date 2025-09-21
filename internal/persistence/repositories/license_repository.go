package repositories

import (
	"context"
	"errors"
	"fmt"

	domain "license-service/internal/domain/model"
	"license-service/internal/domain/repositories"
	entities "license-service/internal/persistence/entities"
	errorInfo "license-service/pkg/log/error"
	logger "license-service/pkg/log/logger"

	"gorm.io/gorm"
)

type licenseRepositoryImpl struct {
	db     *gorm.DB
	logger logger.Logger
}

func NewLicenseRepositoryImpl(db *gorm.DB) repositories.LicenseRepository {
	return &licenseRepositoryImpl{
		db:     db,
		logger: *logger.NewLogger(),
	}
}

func (r *licenseRepositoryImpl) Save(ctx context.Context, license *domain.License) error {
	r.logger.Info("LicenseRepository", "Save", "attempting to save license with folio: "+license.Folio)

	entity := entities.FromDomain(license)

	result := r.db.WithContext(ctx).Create(entity)
	if result.Error != nil {
		var appErr *errorInfo.AppError

		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			appErr = errorInfo.NewAppError(
				errorInfo.ErrInvalidData,
				"LicenseRepository",
				"Save",
				"license with this folio already exists",
			)
			r.logger.Error("LicenseRepository", "Save", appErr, "duplicate folio: "+license.Folio)
		} else {
			appErr = errorInfo.NewAppError(
				errorInfo.ErrInternalError,
				"LicenseRepository",
				"Save",
				fmt.Sprintf("failed to save license: %v", result.Error),
			)
			r.logger.Error("LicenseRepository", "Save", appErr, "database insert failed")
		}
		return appErr
	}

	r.logger.Info("LicenseRepository", "Save", fmt.Sprintf("license saved successfully with ID: %d", entity.ID))
	return nil
}

func (r *licenseRepositoryImpl) FindByFolio(ctx context.Context, folio string) (*domain.License, error) {
	r.logger.Info("LicenseRepository", "FindByFolio", "searching for folio: "+folio)

	var entity entities.LicenseEntity
	result := r.db.WithContext(ctx).Where("folio = ?", folio).First(&entity)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.logger.Info("LicenseRepository", "FindByFolio", "license not found for folio: "+folio)
			return nil, nil
		}

		appErr := errorInfo.NewAppError(
			errorInfo.ErrInternalError,
			"LicenseRepository",
			"FindByFolio",
			fmt.Sprintf("failed to find license: %v", result.Error),
		)
		r.logger.Error("LicenseRepository", "FindByFolio", appErr, "database query failed")
		return nil, appErr
	}

	r.logger.Info("LicenseRepository", "FindByFolio", "license found for folio: "+folio)
	return entity.ToDomain(), nil
}

func (r *licenseRepositoryImpl) FindByPatientID(ctx context.Context, patientID string) ([]*domain.License, error) {
	r.logger.Info("LicenseRepository", "FindByPatientID", "searching licenses for patient: "+patientID)

	var entities []entities.LicenseEntity
	result := r.db.WithContext(ctx).
		Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&entities)

	if result.Error != nil {
		appErr := errorInfo.NewAppError(
			errorInfo.ErrInternalError,
			"LicenseRepository",
			"FindByPatientID",
			fmt.Sprintf("failed to query licenses: %v", result.Error),
		)
		r.logger.Error("LicenseRepository", "FindByPatientID", appErr, "database query failed")
		return nil, appErr
	}

	licenses := make([]*domain.License, 0, len(entities))
	for _, entity := range entities {
		licenses = append(licenses, entity.ToDomain())
	}

	r.logger.Info("LicenseRepository", "FindByPatientID", fmt.Sprintf("found %d licenses for patient: %s", len(licenses), patientID))
	return licenses, nil
}

func (r *licenseRepositoryImpl) ExistsByFolioAndStatus(ctx context.Context, folio string, status string) (bool, error) {
	r.logger.Info("LicenseRepository", "ExistsByFolioAndStatus", fmt.Sprintf("checking existence for folio: %s, status: %s", folio, status))

	var count int64
	result := r.db.WithContext(ctx).
		Model(&entities.LicenseEntity{}).
		Where("folio = ? AND status = ?", folio, status).
		Count(&count)

	if result.Error != nil {
		appErr := errorInfo.NewAppError(
			errorInfo.ErrInternalError,
			"LicenseRepository",
			"ExistsByFolioAndStatus",
			fmt.Sprintf("failed to check existence: %v", result.Error),
		)
		r.logger.Error("LicenseRepository", "ExistsByFolioAndStatus", appErr, "database query failed")
		return false, appErr
	}

	exists := count > 0
	r.logger.Info("LicenseRepository", "ExistsByFolioAndStatus", fmt.Sprintf("license exists: %t", exists))
	return exists, nil
}
