package utils

type ErrorResponse struct {
	Status_code    int    `json:"status_code"`
	Status_message string `json:"status_message"`
}

func ErrorJSON(err string, status int) ErrorResponse {
	res := ErrorResponse{
		Status_code:    status,
		Status_message: err,
	}
	return res
}