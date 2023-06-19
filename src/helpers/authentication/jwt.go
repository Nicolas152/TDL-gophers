package authentication

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

type JWTToken struct {
	UserId 	int 	`json:"userId"`
	Email 	string 	`json:"email"`
	Name 	string 	`json:"name"`
	jwt.StandardClaims
}

func GetJWTHeader(r *http.Request) string {
	prefix := "Bearer "
	authHeader := r.Header.Get("Authorization")
	return strings.TrimPrefix(authHeader, prefix)
}

func CreateJWTHeader(userId int, email string, name string) (string, error) {
	// Creo los claims
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := JWTToken{
		UserId: userId,
		Email: email,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Creo el token con los claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("super-secretKey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(jwtToken string) bool {
	token, _ := jwt.ParseWithClaims(jwtToken, &JWTToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("super-secretKey"), nil
	})

	return token != nil && token.Valid
}

func GetJWTUserId(jwtToken string) (int, error) {
	token, _ := jwt.ParseWithClaims(jwtToken, &JWTToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("super-secretKey"), nil
	})

	if claims, ok := token.Claims.(*JWTToken); ok && token.Valid {
		return claims.UserId, nil
	}

	return 0, errors.New("Invalid JWT token")
}

