package middlewares

import (
	"context"
	"gochat/src/models/request"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	UserId int    `json:"userId"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

func AuthenticationMiddleware(target http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Cargo la request del cliente
		var userRequest request.UserRequest
		if err := userRequest.ReadRequest(r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Realizo validaciones base
		if !userRequest.HasTokenAccess() || !userRequest.IsTokenAccessValid() {
			// TODO: @jesusphilipraiz no debería ser 401?
			http.Error(w, "Could not load server context. Reason: Resource not found", http.StatusNotFound)
			return
		}

		// Agrego el UserID para que el controlador lo pueda usar
		userId := userRequest.GetUserId()
		ctx := context.WithValue(r.Context(), "UserId", userId)

		// Pasar al siguiente controlador con el contexto modificado
		target(w, r.WithContext(ctx))
	}
}
