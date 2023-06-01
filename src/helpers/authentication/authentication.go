package authentication

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"gochat/src/models/authentication"
	"gochat/src/models/request"
	"gochat/src/models/user"
)

// Metodo que desencapsula las credenciales del usuario
func HandlerAuthentication(ws *websocket.Conn) (user.User, error) {
	var userRequest request.UserRequest
	var userCredentials authentication.UserCredentials
	var authenticate authentication.AuthenticationFunction

	// Recibo el mensaje con la acción a realizar
	if err := userRequest.ReadRequest(ws); err != nil {
		return user.User{}, err
	}

	// Obtengo la función a ejecutar en base a la acción
	switch userRequest.GetAction() {
		case request.SignInAction:
			authenticate = userCredentials.SignIn

		case request.LogInAction:
			authenticate = userCredentials.LogIn

		default:
			return user.User{}, errors.New("Invalid 'action' option provided")
	}

	// Si no tiene parametros, no se puede autenticar
	if !userRequest.HasParameters() {
		return user.User{}, errors.New("No parameters provided")
	}

	// Desencapsulo las credenciales del usuario
	if err := json.Unmarshal(userRequest.GetParameters(), &userCredentials); err != nil {
		return user.User{}, err
	}

	// El resto de la autenticacion la delego a otra funcion
	return HandlerUserAuthentication(userCredentials, authenticate)
}

// Metodo que valida las credenciales del usuario y obtiene el data del mismo
func HandlerUserAuthentication(userCredentials authentication.UserCredentials, authenticate authentication.AuthenticationFunction) (user.User, error) {
	// Autentico al usuario con la función correspondiente
	if err := authenticate(); err != nil {
		return user.User{}, err
	}

	// Obtengo contexto del usuario
	userModel := user.User{Email: *userCredentials.Email}
	if err := userModel.GetContext(); err != nil {
		return user.User{}, err
	}

	return userModel, nil
}
