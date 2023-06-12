package authMiddleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

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

func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
	claims := JWTClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secetKey := []byte("super-secretKey") // TODO: move to config file
	tokenString, err := token.SignedString(secetKey)
	if err != nil {
		// log detail of error
		log.Println("Error generating JWT: " + err.Error())
		return "", err
	}
	return tokenString, nil
}
