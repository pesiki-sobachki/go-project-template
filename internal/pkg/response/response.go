package response

import (
	"encoding/json"
	"net/http"
)

type successResponse struct {
	Data interface{} `json:"data,omitempty"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(successResponse{Data: data})
	}
}

func WithError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
}
