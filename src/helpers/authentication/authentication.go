package authentication

import (
	"encoding/json"
	"gochat/src/models/authentication"
	"gochat/src/models/session"
	"gochat/src/models/user"
	"log"
	"net/http"
)

const SessionTokenName = "session-token"

func HandlerSignIn(w http.ResponseWriter, r *http.Request) {
	var userCredentials authentication.UserCredentials

	// Autentico al usuario
	userCredentials.SetAuthenticationType(authentication.SignIn)
	HandlerUserAuthentication(w, r, &userCredentials)
}

func HandlerLogIn(w http.ResponseWriter, r *http.Request) {
	var userCredentials authentication.UserCredentials
	userCredentials.SetAuthenticationType(authentication.LogIn)

	// Autentico al usuario
	log.Println("LogIn")
	HandlerUserAuthentication(w, r, &userCredentials)

	// check if userCredenttials is nil
	if userCredentials.Email == nil {
		return
	}

	// Agrero el token de sesion
	userSession := session.Session{Email: *userCredentials.Email, Token: "123456789"}
	userSession.Add()
}

func HandlerAccess(w http.ResponseWriter, r *http.Request) *user.User {
	// Obtengo el token de sesion
	userSession := session.Get(r.Header.Get(SessionTokenName))
	if userSession == nil {
		http.Error(w, "Could not authenticate user. Reason: Invalid session token", http.StatusBadRequest)
		return nil
	}

	// Resuelvo el email del usuario
	userModel := user.User{Email: (*userSession).Email}
	if err := userModel.GetContext(); err != nil {
		http.Error(w, "Could not authenticate user. Reason: "+err.Error(), http.StatusBadRequest)
		return nil
	}

	return &userModel
}

func HandlerUserAuthentication(w http.ResponseWriter, r *http.Request, userCredentials *authentication.UserCredentials) {
	// Obtengo el usuario y la contraseña del request
	if err := json.NewDecoder(r.Body).Decode(userCredentials); err != nil {
		http.Error(w, "Could not authenticate user. Reason: Empty payload detected", http.StatusBadRequest)
		return
	}

	// Autentico al usuario
	if err := userCredentials.Authenticate(); err != nil {
		http.Error(w, "Could not authenticate user. Reason: "+err.Error(), http.StatusBadRequest)
		return
	}
}
