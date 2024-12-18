package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	response := ErrorResponse{Message: message}
	json.NewEncoder(w).Encode(response)
}
