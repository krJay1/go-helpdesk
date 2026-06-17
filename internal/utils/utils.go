package utils

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func NewApiResponse() *ApiResponse {
	return &ApiResponse{
		Success: false,
		Status:  http.StatusInternalServerError,
	}
}

func (r *ApiResponse) Send(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(r)
}

func SendResponse(w http.ResponseWriter, success bool, status int, message string, data interface{}, error interface{}, contentType string) {
	r := &ApiResponse{
		Status:  status,
		Success: success,
		Message: message,
		Data:    data,
		Error:   error,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}
