package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestAuthMiddleware(t *testing.T) {
	SetJWTKey([]byte("test_secret"))

	validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "testuser",
		"exp":      time.Now().Add(1 * time.Hour).Unix(),
	})
	validTokenString, _ := validToken.SignedString([]byte("test_secret"))

	tests := []struct {
		name           string
		token          string
		expectedStatus int
	}{
		{
			name:           "Valid token",
			token:          validTokenString,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Missing token",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid token",
			token:          "invalid_token",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a dummy handler to wrap with middleware
			dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Protected resource"))
			})

			// Wrap dummy handler with AuthMiddleware
			handler := AuthMiddleware(dummyHandler)

			// Create a test request
			req, _ := http.NewRequest("GET", "/protected", nil)
			if tt.token != "" {
				req.AddCookie(&http.Cookie{
					Name:  "token",
					Value: tt.token,
				})
			}
			rr := httptest.NewRecorder()

			// Serve the request
			handler.ServeHTTP(rr, req)

			// Check response status
			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}
