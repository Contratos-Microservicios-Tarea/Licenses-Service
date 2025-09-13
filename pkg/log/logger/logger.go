package logs

import (
	"os"

	"github.com/sirupsen/logrus"

	errors "license-service/pkg/log/error"
)

type Logger struct {
	*logrus.Logger
}

var globalLogger *Logger

func init() {
	globalLogger = NewLogger()
}

func NewLogger() *Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
		ForceColors:     true,
		DisableQuote:    true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "msg",
		},
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)

	return &Logger{log}
}

func (logs *Logger) LogError(component, operation string, err error, fields ...any) {
	logFields := logs.buildFields(component, operation, fields...)
	logFields["error"] = err.Error()
	logs.WithFields(logFields).Error("❌ ERROR")
}

func (logs *Logger) LogInfo(component, operation string, fields ...any) {
	logFields := logs.buildFields(component, operation, fields...)
	logs.WithFields(logFields).Info("✅ SUCCESS")
}

func (logs *Logger) LogWarn(component, operation string, message string, fields ...any) {
	logFields := logs.buildFields(component, operation, fields...)
	logs.WithFields(logFields).Warn("⚠️  " + message)
}

func (logs *Logger) LogDebug(component, operation string, message string, fields ...any) {
	logFields := logs.buildFields(component, operation, fields...)
	logs.WithFields(logFields).Debug("🔍 " + message)
}

func (logs *Logger) buildFields(component, operation string, fields ...any) logrus.Fields {
	logFields := logrus.Fields{
		"component": component,
		"operation": operation,
	}

	for i := 0; i < len(fields)-1; i += 2 {
		if key, ok := fields[i].(string); ok {
			logFields[key] = fields[i+1]
		}
	}

	return logFields
}

func Error(component, operation string, err error, fields ...any) {
	globalLogger.LogError(component, operation, err, fields...)
}

func Info(component, operation string, fields ...any) {
	globalLogger.LogInfo(component, operation, fields...)
}

func Warn(component, operation string, message string, fields ...any) {
	globalLogger.LogWarn(component, operation, message, fields...)
}

func Debug(component, operation string, message string, fields ...any) {
	globalLogger.LogDebug(component, operation, message, fields...)
}

func (logs *Logger) LogAppError(component, operation string, err *errors.AppError, fields ...any) {
	logFields := logs.buildFields(component, operation, fields...)
	logFields["error_code"] = string(err.Code)
	logFields["error"] = err.Error()

	if err.Details != "" {
		logFields["details"] = err.Details
	}

	logs.WithFields(logFields).Error("❌ ERROR")
}
