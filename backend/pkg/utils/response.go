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
	Meta    interface{}    `json:"meta,omitempty"`
	Code    int            `json:"code"`
}

type PaginationMeta struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

func SuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := APIResponse{
		Status:  SuccessStatus,
		Message: message,
		Data:    data,
		Code:    statusCode,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func ErrorResponse(w http.ResponseWriter, statusCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := APIResponse{
		Status:  ErrorStatus,
		Message: errMsg,
		Code:    statusCode,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func SuccessPaginatedResponse(w http.ResponseWriter, statusCode int, message string, data interface{}, meta interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := APIResponse{
		Status:  SuccessStatus,
		Message: message,
		Data:    data,
		Meta:    meta,
		Code:    statusCode,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
