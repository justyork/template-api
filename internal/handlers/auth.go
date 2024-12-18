package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/justyork/api-template/internal/utils"
)

var jwtKey []byte

func SetJWTKey(key []byte) {
	jwtKey = key
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// LoginHandler godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce text/plain
// @Param credentials body Credentials true "User credentials"
// @Success 200 {string} string "Login successful"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 401 {string} string "Invalid username or password"
// @Router /login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if creds.Username != "testuser" || creds.Password != "password" {
		utils.WriteError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	w.Write([]byte("Login successful"))
}

// ProtectedHandler godoc
// @Summary Access protected resource
// @Description Access a protected route using a JWT token
// @Tags auth
// @Produce text/plain
// @Success 200 {string} string "Welcome"
// @Failure 401 {string} string "Unauthorized"
// @Router /protected [get]
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			utils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		utils.WriteError(w, http.StatusBadRequest, "Bad request")
		return
	}

	tokenStr := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		utils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	w.Write([]byte("Welcome, " + claims.Username))
}
