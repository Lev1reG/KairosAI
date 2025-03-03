package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseStatus string

const (
	SuccessStatus ResponseStatus = "success"
	ErrorStatus   ResponseStatus = "error"
)

type APIResponse struct {
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data,omitempty"`
	Code    int            `json:"code"`
}

func SuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(APIResponse{
		Status:  SuccessStatus,
		Message: message,
		Data:    data,
		Code:    statusCode,
	})
}

func ErrorResponse(w http.ResponseWriter, statusCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(APIResponse{
		Status:  ErrorStatus,
		Message: errMsg,
		Code:    statusCode,
	})
}
