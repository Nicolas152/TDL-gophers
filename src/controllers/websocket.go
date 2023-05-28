package controllers

import (
	"fiuba/concurrent/gochat/src/helpers/authentication"
	"fiuba/concurrent/gochat/src/helpers/messages"
	"fiuba/concurrent/gochat/src/helpers/workspace"
	model "fiuba/concurrent/gochat/src/models/user"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func AddWebsocketController() {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Agrego el controlador de websocket
	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error al conectarse al websocket: ", err)
		}

		// Delego el manejo del mensaje a otra función
		Handler(ws, authentication.HandlerSignIn)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error al conectarse al websocket: ", err)
		}

		// Delego el manejo del mensaje a otra función
		Handler(ws, authentication.HandlerLogIn)
	})
}

func Handler(ws *websocket.Conn, fn authentication.FunctionHandler) {
	defer ws.Close()

	// Envío un mensaje de bienvenida
	messages.SendWelcomeMessage(ws)

	// Valido el UserID del cliente
	user, err := fn(ws)
	if err != nil {
		return
	}

	messages.SendConnectedMessage(ws)
	HandlerMessages(ws, &user)
}

func HandlerMessages(ws *websocket.Conn, user *model.User) {
	for {
		messages.SendMenuMessage(ws)
		_, msg, _ := ws.ReadMessage()

		switch string(msg) {
			case "workspace":
				workspace.HandlerWorkspaceMessages(ws, user)

			case "exit":
				return
			default:
				log.Println("Opcion invalida. Ingrese una de las siguientes opciones: 'workspace', 'exit'")
		}
	}
}