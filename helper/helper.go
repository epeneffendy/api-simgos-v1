package helper

type Response struct {
	Message string
	Code    int
	Status  string
	Data    interface{}
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	response := Response{
		Message: message,
		Code:    code,
		Status:  status,
		Data:    data,
	}

	return response
}
