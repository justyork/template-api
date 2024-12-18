package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justyork/api-template/internal/handlers"
	"github.com/justyork/api-template/internal/middleware"
)

// RegisterRoutes initializes all application routes
func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Apply CORS middleware to all routes
	r.Use(middleware.CORSMiddleware)

	// Health check route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API is running"))
	})

	// Public routes
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Protected route with middleware
	r.Handle("/protected", middleware.AuthMiddleware(http.HandlerFunc(handlers.ProtectedHandler))).Methods("GET")

	return r
}
