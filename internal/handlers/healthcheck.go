package handlers

import (
	"net/http"
)

// HealthCheckHandler returns the status of the API
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is running"))
}
