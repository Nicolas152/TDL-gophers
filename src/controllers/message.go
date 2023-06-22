package controllers

import (
	"github.com/gorilla/mux"
	"gochat/src/services/message"
)

// TODO: Agregar el middleware de autenticacion
func AddWebsocketController(myRouter *mux.Router) {
	// Handler para manipular la conexion websocket
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel/{channelKey}/message", message.HandlerMessages)
}
