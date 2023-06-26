package controllers

import (
	"gochat/src/middlewares"
	"gochat/src/services/message"

	"github.com/gorilla/mux"
)

func AddWebsocketController(myRouter *mux.Router) {
	// Handler para manipular la conexion websocket
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel/{channelKey}/message", middlewares.AuthenticationMiddleware(message.HandlerMessages))
}
