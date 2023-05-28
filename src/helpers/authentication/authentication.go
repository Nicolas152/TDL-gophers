package authentication

import (
	"fiuba/concurrent/gochat/src/helpers/messages"
	handler "fiuba/concurrent/gochat/src/helpers/user"
	model "fiuba/concurrent/gochat/src/models/user"
	"github.com/gorilla/websocket"
)

type FunctionHandler func(ws *websocket.Conn) (model.User, error)

func HandlerSignIn(ws *websocket.Conn) (model.User, error) {
	messages.SendSignInMessage(ws)

	// Creo el usuario
	user, err := handler.HandlerCreateUser(ws)
	if err != nil {
		messages.SendErrorInvalidUserMessage(ws)
		return model.User{}, err
	}

	return user, nil
}

func HandlerLogIn(ws *websocket.Conn) (model.User, error) {
	messages.SendLogInMessage(ws)

	// Valido el usuario
	user, err := handler.HandlerValidateUser(ws)
	if err != nil {
		messages.SendErrorInvalidUserMessage(ws)
		return model.User{}, err
	}

	return user, nil
}

