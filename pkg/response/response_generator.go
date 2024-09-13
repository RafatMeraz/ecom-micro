package response

import (
	"github.com/RafatMeraz/ecom-micro/pkg/errors"
	"github.com/RafatMeraz/ecom-micro/pkg/models"
	"net/http"
)

func GenerateResponse(data interface{}, err error, isMessage bool) (int, map[string]interface{}) {
	if err != nil {
		c, r := errors.GetErrorResponse(err)
		return c, r.ToJson()
	} else if isMessage {
		return http.StatusOK, models.SuccessMessage{Message: data.(string)}.ToJson()
	} else {
		return http.StatusOK, models.SuccessResponse{Data: data}.ToJson()
	}
}
