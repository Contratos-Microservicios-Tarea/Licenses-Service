package errors

import (
	"fmt"
	"time"
)

type ErrorCode string

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

const (
	ErrDBConnection  ErrorCode = "DB_CONNECTION_FAILED"
	ErrDBTimeout     ErrorCode = "DB_TIMEOUT"
	ErrDBNotFound    ErrorCode = "DB_NOT_FOUND"
	ErrDBQueryFailed ErrorCode = "DB_QUERY_FAILED"
	ErrDBConflict    ErrorCode = "DB_CONFLICT"
	ErrDBDeadlock    ErrorCode = "DB_DEADLOCK"
	ErrDBMigration   ErrorCode = "DB_MIGRATION_FAILED"
	ErrDBTransaction ErrorCode = "DB_TRANSACTION_FAILED"

	ErrDeviceNotFound  ErrorCode = "DEVICE_NOT_FOUND"
	ErrMeasureNotFound ErrorCode = "MEASURE_NOT_FOUND"
	ErrInvalidData     ErrorCode = "INVALID_DATA"
	ErrAlreadyExists   ErrorCode = "ALREADY_EXISTS"
	ErrResourceLocked  ErrorCode = "RESOURCE_LOCKED"

	ErrValidationFailed     ErrorCode = "VALIDATION_FAILED"
	ErrMissingRequiredField ErrorCode = "MISSING_REQUIRED_FIELD"
	ErrInvalidFormat        ErrorCode = "INVALID_FORMAT"
	ErrValueOutOfRange      ErrorCode = "VALUE_OUT_OF_RANGE"

	ErrServiceConfig      ErrorCode = "SERVICE_CONFIG_ERROR"
	ErrServiceInit        ErrorCode = "SERVICE_INIT_FAILED"
	ErrServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"

	ErrExternalTimeout            ErrorCode = "EXTERNAL_SERVICE_TIMEOUT"
	ErrExternalError              ErrorCode = "EXTERNAL_SERVICE_ERROR"
	ErrExternalBadRequest         ErrorCode = "EXTERNAL_SERVICE_BAD_REQUEST"
	ErrExternalAuth               ErrorCode = "EXTERNAL_SERVICE_AUTH_ERROR"
	ErrExternalNotFound           ErrorCode = "EXTERNAL_SERVICE_NOT_FOUND"
	ErrExternalServiceUnavailable ErrorCode = "EXTERNAL_SERVICE_UNAVAILABLE"
	ErrExternalServiceBadResponse ErrorCode = "EXTERNAL_SERVICE_BAD_RESPONSE"

	ErrFileUploadFailed    ErrorCode = "FILE_UPLOAD_FAILED"
	ErrFileTooLarge        ErrorCode = "FILE_TOO_LARGE"
	ErrUnsupportedFileType ErrorCode = "UNSUPPORTED_FILE_TYPE"

	ErrTooManyRequests ErrorCode = "TOO_MANY_REQUESTS"

	ErrOptimisticLockFailed ErrorCode = "OPTIMISTIC_LOCK_FAILED"

	ErrNetworkError ErrorCode = "NETWORK_ERROR"

	ErrBadRequest            ErrorCode = "BAD_REQUEST"
	ErrRequestTimeout        ErrorCode = "REQUEST_TIMEOUT"
	ErrMethodNotAllowed      ErrorCode = "METHOD_NOT_ALLOWED"
	ErrNotAcceptable         ErrorCode = "NOT_ACCEPTABLE"
	ErrUnsupportedMediaType  ErrorCode = "UNSUPPORTED_MEDIA_TYPE"
	ErrUnprocessableEntity   ErrorCode = "UNPROCESSABLE_ENTITY"
	ErrGatewayTimeout        ErrorCode = "GATEWAY_TIMEOUT"
	ErrRequestEntityTooLarge ErrorCode = "REQUEST_ENTITY_TOO_LARGE"
	ErrRequestURITooLong     ErrorCode = "REQUEST_URI_TOO_LONG"
	ErrPreconditionFailed    ErrorCode = "PRECONDITION_FAILED"
	ErrExpectationFailed     ErrorCode = "EXPECTATION_FAILED"
	ErrUpgradeRequired       ErrorCode = "UPGRADE_REQUIRED"
	ErrTooManyRedirects      ErrorCode = "TOO_MANY_REDIRECTS"
	ErrMalformedRequest      ErrorCode = "MALFORMED_REQUEST"
	ErrInvalidContentType    ErrorCode = "INVALID_CONTENT_TYPE"
	ErrInvalidAcceptHeader   ErrorCode = "INVALID_ACCEPT_HEADER"

	ErrCacheMiss             ErrorCode = "CACHE_MISS"
	ErrCacheConnectionFailed ErrorCode = "CACHE_CONNECTION_FAILED"
	ErrConfigLoadFailed      ErrorCode = "CONFIG_LOAD_FAILED"
	ErrDependencyFailure     ErrorCode = "DEPENDENCY_FAILURE"
	ErrShutdownInProgress    ErrorCode = "SHUTDOWN_IN_PROGRESS"
	ErrMemoryExhausted       ErrorCode = "MEMORY_EXHAUSTED"
	ErrDiskSpaceExhausted    ErrorCode = "DISK_SPACE_EXHAUSTED"
	ErrResourceExhausted     ErrorCode = "RESOURCE_EXHAUSTED"
	ErrCircuitBreakerOpen    ErrorCode = "CIRCUIT_BREAKER_OPEN"
	ErrHealthCheckFailed     ErrorCode = "HEALTH_CHECK_FAILED"
	ErrGracefulShutdown      ErrorCode = "GRACEFUL_SHUTDOWN_FAILED"

	ErrInternalError ErrorCode = "INTERNAL_ERROR"
	ErrUnknownError  ErrorCode = "UNKNOWN_ERROR"
	ErrNotFound      ErrorCode = "NOT_FOUND"
	ErrConflict      ErrorCode = "ERROR_CONFLICT"

	ErrJSONToDBConversion ErrorCode = "JSON_TO_DB_CONVERSION_ERROR"
	ErrDBToJSONConversion ErrorCode = "DB_TO_JSON_CONVERSION_ERROR"
)

var ErrorCodeMap = map[ErrorCode]struct {
	HTTPStatus int
	Message    string
}{

	ErrDBConnection:  {500, "Error de conexión a la base de datos"},
	ErrDBTimeout:     {504, "Tiempo de espera agotado en la base de datos"},
	ErrDBNotFound:    {404, "Recurso no encontrado en la base de datos"},
	ErrDBQueryFailed: {500, "Error en la consulta a la base de datos"},
	ErrDBConflict:    {409, "Conflicto de datos en la base de datos"},
	ErrDBDeadlock:    {409, "Bloqueo mutuo detectado en la base de datos"},
	ErrDBMigration:   {500, "Error en la migración de la base de datos"},
	ErrDBTransaction: {500, "Error en la transacción de la base de datos"},

	ErrDeviceNotFound:  {404, "Dispositivo no encontrado"},
	ErrMeasureNotFound: {404, "Medida no encontrada"},
	ErrInvalidData:     {400, "Datos inválidos"},
	ErrAlreadyExists:   {409, "El recurso ya existe"},
	ErrResourceLocked:  {423, "El recurso está bloqueado"},

	ErrValidationFailed:     {400, "Fallo de validación de datos"},
	ErrMissingRequiredField: {400, "Campo requerido faltante"},
	ErrInvalidFormat:        {400, "Formato de datos inválido"},
	ErrValueOutOfRange:      {400, "Valor fuera de rango permitido"},

	ErrServiceConfig:      {500, "Error de configuración del servicio"},
	ErrServiceInit:        {500, "Error al inicializar el servicio"},
	ErrServiceUnavailable: {503, "Servicio no disponible temporalmente"},

	ErrBadRequest:            {400, "Solicitud incorrecta"},
	ErrRequestTimeout:        {408, "Tiempo de espera de solicitud agotado"},
	ErrMethodNotAllowed:      {405, "Método no permitido"},
	ErrNotAcceptable:         {406, "No aceptable"},
	ErrUnsupportedMediaType:  {415, "Tipo de medio no soportado"},
	ErrUnprocessableEntity:   {422, "Entidad no procesable"},
	ErrGatewayTimeout:        {504, "Tiempo de espera de gateway agotado"},
	ErrRequestEntityTooLarge: {413, "Entidad de solicitud demasiado grande"},
	ErrRequestURITooLong:     {414, "URI de solicitud demasiado largo"},
	ErrPreconditionFailed:    {412, "Precondición fallida"},
	ErrExpectationFailed:     {417, "Expectativa fallida"},
	ErrUpgradeRequired:       {426, "Actualización requerida"},
	ErrTooManyRedirects:      {310, "Demasiadas redirecciones"},
	ErrMalformedRequest:      {400, "Solicitud malformada"},
	ErrInvalidContentType:    {400, "Tipo de contenido inválido"},
	ErrInvalidAcceptHeader:   {400, "Cabecera Accept inválida"},

	ErrCacheMiss:             {404, "Elemento no encontrado en caché"},
	ErrCacheConnectionFailed: {500, "Error de conexión al caché"},
	ErrConfigLoadFailed:      {500, "Error al cargar configuración"},
	ErrDependencyFailure:     {503, "Fallo de dependencia"},
	ErrShutdownInProgress:    {503, "Apagado en progreso"},
	ErrMemoryExhausted:       {507, "Memoria agotada"},
	ErrDiskSpaceExhausted:    {507, "Espacio en disco agotado"},
	ErrResourceExhausted:     {429, "Recursos agotados"},
	ErrCircuitBreakerOpen:    {503, "Circuit breaker abierto"},
	ErrHealthCheckFailed:     {503, "Verificación de salud fallida"},
	ErrGracefulShutdown:      {500, "Error en apagado gradual"},

	ErrExternalTimeout:            {504, "Tiempo de espera agotado con servicio externo"},
	ErrExternalError:              {502, "Error en servicio externo"},
	ErrExternalBadRequest:         {400, "Solicitud inválida a servicio externo"},
	ErrExternalAuth:               {401, "Error de autenticación con servicio externo"},
	ErrExternalNotFound:           {404, "Recurso no encontrado en servicio externo"},
	ErrExternalServiceUnavailable: {503, "Servicio externo no disponible"},
	ErrExternalServiceBadResponse: {502, "Respuesta inesperada de servicio externo"},

	ErrFileUploadFailed:    {500, "Fallo al subir archivo"},
	ErrFileTooLarge:        {413, "Archivo demasiado grande"},
	ErrUnsupportedFileType: {415, "Tipo de archivo no soportado"},

	ErrTooManyRequests: {429, "Demasiadas solicitudes"},

	ErrOptimisticLockFailed: {409, "Fallo de bloqueo optimista"},

	ErrNetworkError: {500, "Error de red interno"},

	ErrInternalError: {500, "Error interno del servidor"},
	ErrUnknownError:  {500, "Error desconocido"},
	ErrNotFound:      {404, "Recurso no encontrado"},
	ErrConflict:      {409, "Conflicto de recursos"},

	ErrJSONToDBConversion: {500, "Error de conversión JSON a DB"},
	ErrDBToJSONConversion: {500, "Error de conversión DB a JSON"},
}

type AppError struct {
	Code      ErrorCode `json:"code"`
	Message   string    `json:"message"`
	Details   string    `json:"details,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Component string    `json:"component"`
	Operation string    `json:"operation"`
	Cause     error     `json:"-"`
}

func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s: %s (caused by: %v)", e.Code, e.Message, e.Details, e.Cause)
	}
	return fmt.Sprintf("[%s] %s: %s", e.Code, e.Message, e.Details)
}

func (e *AppError) Unwrap() error {
	return e.Cause
}

func NewAppError(code ErrorCode, component, operation, message string) *AppError {
	if message == "" {
		if errorInfo, exists := ErrorCodeMap[code]; exists {
			message = errorInfo.Message
		}
	}

	return &AppError{
		Code:      code,
		Message:   message,
		Component: component,
		Operation: operation,
		Timestamp: time.Now(),
	}
}

func WrapError(code ErrorCode, component, operation, message string, cause error) *AppError {
	if message == "" {
		if errorInfo, exists := ErrorCodeMap[code]; exists {
			message = errorInfo.Message
		}
	}

	return &AppError{
		Code:      code,
		Message:   message,
		Component: component,
		Operation: operation,
		Timestamp: time.Now(),
		Cause:     cause,
	}
}

func (e *AppError) WithDetails(details string) *AppError {
	e.Details = details
	return e
}

func IsNotFoundError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == ErrNotFound || appErr.Code == ErrDBNotFound || appErr.Code == ErrDeviceNotFound || appErr.Code == ErrMeasureNotFound || appErr.Code == ErrExternalNotFound
	}
	return false
}

func IsAppErrorCode(err error, code string) bool {
	if appErr, ok := err.(*AppError); ok {
		return string(appErr.Code) == code
	}
	return false
}

func (e *AppError) HTTPStatus() int {
	if errorInfo, exists := ErrorCodeMap[e.Code]; exists {
		return errorInfo.HTTPStatus
	}
	return 500
}

func (e *AppError) ToErrorDTO() Error {
	return Error{
		Code:    string(e.Code),
		Message: e.Message,
		Details: e.Details,
	}
}
