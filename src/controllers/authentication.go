package controllers

import (
	"encoding/json"
	"errors"
	"gochat/src/helpers/authentication"
	"net/http"

	authDTO "gochat/src/controllers/DTOs/auth"
	"gochat/src/models/user"

	"github.com/gorilla/mux"
)

func AddAuthenticationsController(myRouter *mux.Router) {
	// Handler of signin
	myRouter.HandleFunc("/gophers/signin", func(w http.ResponseWriter, r *http.Request) {
		HandlerSignIn(w, r)
	})

	// Handler of login
	myRouter.HandleFunc("/gophers/login", func(w http.ResponseWriter, r *http.Request) {
		HandlerLogIn(w, r)
	})
}

func HandlerLogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginDTO authDTO.LoginDTO
	_ = json.NewDecoder(r.Body).Decode(&loginDTO)

	// Valido las credenciales del usuario
	userModel := user.User{Email: loginDTO.Email, Password: loginDTO.Password}
	if authenticated := userModel.Authenticate(); !authenticated {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Obtengo el contexto del usuario
	err := userModel.GetContext()
	if err != nil {
		http.Error(w, "Error with user context: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Genero el token
	tokenString, err := authentication.CreateJWTHeader(userModel.Id, userModel.Email, userModel.Name)
	if err != nil {
		http.Error(w, "Failed to generate JWT", http.StatusInternalServerError)
		return
	}

	// Respondo con el token
	response := map[string]string{"access_token": tokenString}
	json.NewEncoder(w).Encode(response)
	return
}

func HandlerSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var signinDTO authDTO.SignInDTO
	_ = json.NewDecoder(r.Body).Decode(&signinDTO)

	// Validate sign in data is not empty
	if err := validateSignInData(&signinDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userModel := user.User{Email: signinDTO.Email, Password: signinDTO.Password, Name: signinDTO.Name}
	if err := userModel.Create(); err != nil {
		http.Error(w, "Failed to create user:"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "User created successfully"}
	json.NewEncoder(w).Encode(response)
	return
}

func validateSignInData(signinDTO *authDTO.SignInDTO) error {
	// Validate sign in data is not empty
	if signinDTO.Email == "" || signinDTO.Password == "" || signinDTO.Name == "" {
		return errors.New("invalid data")
	}
	return nil
}
