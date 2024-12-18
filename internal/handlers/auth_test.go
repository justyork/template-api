package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestLoginHandler(t *testing.T) {
	SetJWTKey([]byte("test_secret"))

	tests := []struct {
		name           string
		credentials    Credentials
		expectedStatus int
	}{
		{
			name: "Valid login",
			credentials: Credentials{
				Username: "testuser",
				Password: "password",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Invalid login",
			credentials: Credentials{
				Username: "wronguser",
				Password: "wrongpassword",
			},
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.credentials)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(LoginHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}

func TestProtectedHandler(t *testing.T) {
	SetJWTKey([]byte("test_secret"))

	validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: "testuser",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	})
	tokenString, _ := validToken.SignedString([]byte("test_secret"))

	tests := []struct {
		name           string
		token          string
		expectedStatus int
	}{
		{
			name:           "Valid token",
			token:          tokenString,
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
			req, _ := http.NewRequest("GET", "/protected", nil)
			if tt.token != "" {
				req.AddCookie(&http.Cookie{
					Name:  "token",
					Value: tt.token,
				})
			}
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(ProtectedHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}
