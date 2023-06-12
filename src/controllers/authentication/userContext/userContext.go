package userContext

import (
	"gochat/src/controllers/authentication/authMiddleware"

	"net/http"
)

type UserContext struct {
	Email string `json:"email"`
}

func GetUserContext(r *http.Request) *UserContext {

	claims := r.Context().Value("userClaims").(*authMiddleware.JWTClaims)
	userContext := UserContext{Email: claims.Email}

	return &userContext
}
