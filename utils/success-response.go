package utils

//JSON success response model
type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//Success response builder
func ResponseJSON(status int, message string, data interface{}) SuccessResponse {
	result := SuccessResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return result
}
