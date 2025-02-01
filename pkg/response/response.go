package response

type APIResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code,omitempty"`
}

func SuccessResponse(data interface{}) APIResponse {
	return APIResponse{
		Status: "success",
		Data:   data,
	}
}

func FailResponse(data interface{}) APIResponse {
	return APIResponse{
		Status: "fail",
		Data:   data,
	}
}

func ErrorResponse(message string, code int, data interface{}) APIResponse {
	return APIResponse{
		Status:  "error",
		Message: message,
		Code:    code,
		Data:    data,
	}
}
