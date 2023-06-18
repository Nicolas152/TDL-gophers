package middlewares

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"gochat/src/models/request"
	"net/http"
)

type JWTClaims struct {
	UserId 	int 	`json:"userId"`
	Email 	string 	`json:"email"`
	Name 	string 	`json:"name"`
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
