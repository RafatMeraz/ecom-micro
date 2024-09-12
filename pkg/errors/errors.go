package errors

import "errors"

var (
	ErrFailedToLoadConfigs      = errors.New("failed to load configurations")
	ErrRequiredFieldsMissing    = errors.New("required fields are missing")
	ErrRequiredParamsMissing    = errors.New("required params are missing")
	ErrInvalidData              = errors.New("required fields data invalid")
	ErrDatabaseConnectionFailed = errors.New("database connection failed")
	ErrNoDataFound              = errors.New("no data found")
	ErrNoContent                = errors.New("no content")
	ErrDatabaseOperation        = errors.New("failed database operation")
	ErrRecordNotFound           = errors.New("record not found")
	ErrUsernameExists           = errors.New("username exists")
	ErrUnAuthorization          = errors.New("unauthorized")
	ErrFileUploadFailed         = errors.New("failed file upload")
	ErrProductCodeExists        = errors.New("product code exists")
	ErrInternalServer           = errors.New("internal server error")
	ErrStockUnavailable         = errors.New("stock not available")
	ErrInvalidToken             = errors.New("invalid token")
)
