package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	envConfig "license-service/pkg/env"
	appError "license-service/pkg/log/error"
	appLogger "license-service/pkg/log/logger"
)

func NewConnection() (*gorm.DB, error) {
	config := envConfig.Load()
	log := appLogger.NewLogger()

	var dsn string

	// Priorizar URL completa si está disponible
	if config.Database.Url != "" {
		dsn = config.Database.Url
		log.Info("Database", "NewConnection", "Using POSTGRES_URL from environment")
	} else {
		// Construir DSN desde componentes individuales
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			config.Database.Host,
			config.Database.Username,
			config.Database.Password,
			config.Database.Name,
			config.Database.Port,
			config.Database.SSLMode,
			config.Database.TimeZone)
		log.Info("Database", "NewConnection", "Building DSN from individual components")
	}

	// Configurar nivel de log de GORM
	gormLogLevel := getGormLogLevel(config.Database.GormLogLevel)

	dbLogger := logger.New(
		appLogger.NewLogger(),
		logger.Config{
			SlowThreshold: config.Database.SlowThreshold,
			LogLevel:      gormLogLevel,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		dbError := appError.NewAppError(
			appError.ErrDBConnection,
			"Database",
			"NewConnection",
			"Failed to connect to database")
		log.Error("Database", "NewConnection", dbError, "error: "+err.Error())
		return nil, dbError
	}

	// Configurar pool de conexiones
	if err := configureConnectionPool(db, config); err != nil {
		return nil, err
	}

	log.Info("Database", "NewConnection", "Database connection established successfully")
	return db, nil
}

func configureConnectionPool(db *gorm.DB, config *envConfig.Config) error {
	log := appLogger.NewLogger()

	sqlDB, err := db.DB()
	if err != nil {
		poolError := appError.NewAppError(
			appError.ErrDBConnection,
			"Database",
			"configureConnectionPool",
			"Failed to get underlying sql.DB")
		log.Error("Database", "configureConnectionPool", poolError, "error: "+err.Error())
		return poolError
	}

	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.Database.ConnMaxLifetime)

	log.Info("Database", "configureConnectionPool", fmt.Sprintf("Connection pool configured - MaxIdle: %d, MaxOpen: %d, MaxLifetime: %v",
		config.Database.MaxIdleConns, config.Database.MaxOpenConns, config.Database.ConnMaxLifetime))

	return nil
}

func getGormLogLevel(level string) logger.LogLevel {
	switch level {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info
	}
}
