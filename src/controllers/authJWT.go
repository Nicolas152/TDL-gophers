package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	authDTO "gochat/src/controllers/DTOs/auth"
	"gochat/src/controllers/authentication/authMiddleware"
	"gochat/src/models/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func AddAuthJWTController(myRouter *mux.Router) {
	// Handler of signin
	myRouter.HandleFunc("/gophers/signin", func(w http.ResponseWriter, r *http.Request) {
		HandlerSignIn(w, r)
	})

	// Handler of login
	myRouter.HandleFunc("/gophers/loginJWT", func(w http.ResponseWriter, r *http.Request) {
		HandlerLogIn(w, r)
	})
}

func HandlerLogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginDTO authDTO.LoginDTO
	_ = json.NewDecoder(r.Body).Decode(&loginDTO)

	// Validate user credentials
	userModel := user.User{Email: loginDTO.Email, Password: loginDTO.Password}
	if authenticated := userModel.Authenticate(); !authenticated {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	tokenString, err := generateJWT(userModel.Email)
	if err != nil {
		http.Error(w, "Failed to generate JWT", http.StatusInternalServerError)
		return
	}
	response := map[string]string{"token": tokenString}
	json.NewEncoder(w).Encode(response)
	return
}

func generateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(time.Hour) // Token expires in 1 hour
	claims := &authMiddleware.JWTClaims{
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

func HandlerSignIn(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
