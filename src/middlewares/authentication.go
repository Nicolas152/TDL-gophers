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
		// Load the client's request
		var userRequest request.UserRequest
		if err := userRequest.ReadRequest(r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Perform basic validations
		if !userRequest.HasTokenAccess() || !userRequest.IsTokenAccessValid() {
			http.Error(w, "Could not load server context. Reason: Resource not found", http.StatusNotFound)
			return
		}

		// Add the UserID so that the controller can use it
		userId := userRequest.GetUserId()
		ctx := context.WithValue(r.Context(), "UserId", userId)

		// Proceed to the next handler with the modified context
		target(w, r.WithContext(ctx))
	}
}
