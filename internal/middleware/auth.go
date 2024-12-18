package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/justyork/api-template/internal/utils"
)

var jwtKey []byte

// SetJWTKey sets the JWT secret key for the middleware
func SetJWTKey(key []byte) {
	jwtKey = key
}

// AuthMiddleware checks the validity of a JWT token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		claims := &jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			utils.WriteError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Add claims to context for further use
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
