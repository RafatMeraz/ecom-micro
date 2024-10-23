package errors

import (
	"errors"
	"github.com/RafatMeraz/ecom-micro/pkg/models"
	"net/http"
)

func GetErrorResponse(err error) (int, models.Response) {
	var statusCode int
	var response models.Response

	switch {
	case errors.Is(err, ErrUnAuthorization):
		statusCode = http.StatusUnauthorized
		response = models.ErrorResponseModel{Message: err.Error()}
		break
	case errors.Is(err, ErrNoContent):
		statusCode = http.StatusOK
		response = models.SuccessResponse{Data: []string{}}
		break
	case errors.Is(err, ErrEmailExists):
		statusCode = http.StatusBadRequest
		response = models.ErrorResponseModel{Message: err.Error()}
		break
	case errors.Is(err, ErrRequiredFieldsMissing):
		statusCode = http.StatusBadRequest
		response = models.ErrorResponseModel{Message: err.Error()}
		break
	case errors.Is(err, ErrInvalidData):
		statusCode = http.StatusBadRequest
		response = models.ErrorResponseModel{Message: err.Error()}
		break
	case errors.Is(err, ErrDatabaseOperation):
		statusCode = http.StatusInternalServerError
		response = models.ErrorResponseModel{Message: "internal server error"}
	case errors.Is(err, ErrRecordNotFound):
		statusCode = http.StatusNotFound
		response = models.ErrorResponseModel{Message: err.Error()}
	case errors.Is(err, ErrNoDataFound):
		statusCode = http.StatusBadRequest
		response = models.ErrorResponseModel{Message: err.Error()}
		break
	case errors.Is(err, ErrRequiredParamsMissing):
		statusCode = http.StatusBadRequest
		response = models.ErrorResponseModel{Message: err.Error()}
		break
	case errors.Is(err, ErrProductCodeExists):
		statusCode = http.StatusBadRequest
		response = models.ErrorResponseModel{Message: err.Error()}
		break
	case errors.Is(err, ErrStockUnavailable):
		statusCode = http.StatusBadRequest
		response = models.ErrorResponseModel{Message: err.Error()}
		break
	case errors.Is(err, ErrInvalidToken):
		statusCode = http.StatusBadRequest
		response = models.ErrorResponseModel{Message: err.Error()}
		break
	}

	return statusCode, response
}
