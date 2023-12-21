package shared

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func BuildResponse(status string, data interface{}) Response {
	res := Response{
		Status: status,
		Data:   data,
	}
	return res
}

func BuildErrorResponse(status string, message string) ErrorResponse {
	res := ErrorResponse{
		Status:  status,
		Message: message,
	}
	return res
}
