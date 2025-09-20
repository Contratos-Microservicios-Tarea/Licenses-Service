package database

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	envConfig "license-service/pkg/env"
	appError "license-service/pkg/log/error"
	appLogger "license-service/pkg/log/logger"
)

func NewConnection() (*gorm.DB, error) {
	config := envConfig.Load()

	var dsn string

	if config.Database.Url != "" {
		dsn = strings.Replace(config.Database.Url, "db:5432", "localhost:5432", 1)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Santiago",
			"localhost",
			config.Database.Username,
			config.Database.Password,
			config.Database.Name,
			"5432")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		dbError := appError.NewAppError(
			appError.ErrDBConnection,
			"Database",
			"NewConnection",
			"Failed to connect to database")
		appLogger.Error("Database", "NewConnection", dbError, "dsn", dsn, "error", err.Error())
		return nil, dbError
	}

	if err := configureConnectionPool(db); err != nil {
		poolError := appError.NewAppError(
			appError.ErrDBConnection,
			"Database",
			"configureConnectionPool",
			"Failed to configure connection pool")
		appLogger.Error("Database", "configureConnectionPool", poolError, "error", err.Error())
		return nil, poolError
	}

	appLogger.NewLogger().Info("Successfully | Database connection established")
	return db, nil
}

func configureConnectionPool(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		poolError := appError.NewAppError(
			appError.ErrDBConnection,
			"Database",
			"configureConnectionPool",
			"Failed to get underlying sql.DB")
		appLogger.Error("Database", "configureConnectionPool", poolError, "error", err.Error())
		return poolError
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	appLogger.NewLogger().Info("Successfully | Connection pool configured")
	return nil
}
