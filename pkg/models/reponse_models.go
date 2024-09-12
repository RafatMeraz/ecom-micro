package models

type Response interface {
	ToJson() map[string]interface{}
}

type SuccessResponse struct {
	Data interface{}
}

func (model SuccessResponse) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   model.Data,
	}
}

type SuccessMessage struct {
	Message string
}

func (model SuccessMessage) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": model.Message,
	}
}

type ErrorResponseModel struct {
	Message string
}

func (model ErrorResponseModel) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"status":  "failed",
		"message": model.Message,
	}
}
