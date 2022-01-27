package utils

//JSON success response model
type SuccessResponse struct {
	Status  int         `json:"status_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type InsertSuccessResponse struct {
	Status     int    `json:"status_code"`
	Message    string `json:"message"`
	Request_id int    `json:"request_id"`
}

//Success response builder
func ResponseJSON(status int, message string, data interface{}) SuccessResponse {
	result := SuccessResponse{
		Status:  status,
		Message: "success",
		Data:    data,
	}
	return result
}

func InsertResponseJSON(status int, message string, data int) InsertSuccessResponse {
	result := InsertSuccessResponse{
		Status:     status,
		Message:    message,
		Request_id: data,
	}
	return result
}
