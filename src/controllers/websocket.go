package controllers

import (
	"errors"
	"github.com/gorilla/websocket"
	"gochat/src/helpers/authentication"
	"gochat/src/helpers/common"
	"gochat/src/helpers/messages"
	"log"
	"net/http"
)

func AddWebsocketController() {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Agrego el controlador de websocket
	http.HandleFunc("/gochat", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Could not connect with the websocket. Reason: ", err)
		}

		// Delego el manejo del websocket a otra función
		HandlerWebsocket(ws)
	})
}

func HandlerWebsocket(ws *websocket.Conn) {
	defer ws.Close()

	// Envío un mensaje de bienvenida
	common.SendWelcomeMessage(ws)

	// Valido el UserID del cliente
	user, err := authentication.HandlerAuthentication(ws)
	if err != nil {
		common.SendErrorMessage(ws, errors.New("Could not authenticate user. Reason: " + err.Error()))
		return
	}

	common.SendConnectedMessage(ws)

	// Delego el manejo de los mensajes a otra función
	if err := messages.HandlerMessages(ws, &user); err != nil {
		common.SendErrorMessage(ws, errors.New("Could not process message. Reason: " + err.Error()))
		return
	}
}