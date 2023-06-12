package authMiddleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func VerifyTokenMiddleware(target http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := getJWTFromHeader(r)
		if tokenString == "" {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		secretKey := []byte("super-secretKey")
		// Verify token
		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
			return
		}

		// Add claims to context of request
		claims := token.Claims.(*JWTClaims)
		ctx := context.WithValue(r.Context(), "userClaims", claims)

		// Pasar al siguiente controlador con el contexto modificado
		target(w, r.WithContext(ctx))
	}
}

func getJWTFromHeader(r *http.Request) string {
	prefix := "Bearer "
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, prefix)
	return tokenString
}
