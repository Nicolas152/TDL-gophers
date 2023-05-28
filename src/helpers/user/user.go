package user

import (
	"errors"
	"fiuba/concurrent/gochat/src/helpers/messages"
	"fiuba/concurrent/gochat/src/models/user"
	"github.com/gorilla/websocket"
)

func HandlerCreateUser(ws *websocket.Conn) (user.User, error) {
	// Pido el contexto necesario para crear un nuevo usuario
	email, name, password := messages.AskSignInMessage(ws)

	if len(email) == 0 || len(password) == 0 {
		messages.SendErrorEmptyMessage(ws)
		return user.User{}, errors.New("Credenciales vacias")
	}

	if len(name) == 0 {
		messages.SendErrorEmptyMessage(ws)
		return user.User{}, errors.New("Nombre vacio")
	}

	// Valido que no exista en DB
	exists, _ := user.Exists(email)
	if exists {
		return user.User{}, errors.New("El usuario ya existe")
	}

	return user.Create(email, name, password)
}

func HandlerValidateUser(ws *websocket.Conn) (user.User, error) {
	// Pido el contexto necesario para autenticar al usuario
	email, password := messages.AskLogInMessage(ws)

	if len(email) == 0 || len(password) == 0 {
		messages.SendErrorEmptyMessage(ws)
		return user.User{}, errors.New("Credenciales vacias")
	}

	// Valido que el User exista en DB
	exists, _ := user.Exists(email)
	if !exists {
		return user.User{}, errors.New("El usuario no existe")
	}

	user, err := user.GetUser(email)
	if err != nil {
		return user, err
	}

	// Valido que el password sea correcto
	if !user.ValidatePassword(password) {
		return user, errors.New("El password es incorrecto")
	}

	return user, nil
}