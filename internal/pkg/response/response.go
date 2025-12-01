package response

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Data T `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(Response[interface{}]{Data: data}); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func WithError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}
