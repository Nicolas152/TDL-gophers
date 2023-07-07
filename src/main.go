package main

import (
	"gochat/src/controllers"
	"gochat/src/models/message/subscription"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AddRoutes(myRouter *mux.Router) {
	controllers.AddAuthenticationsController(myRouter)

	controllers.AddWorkspaceController(myRouter)
	controllers.AddChannelController(myRouter)
	controllers.AddDMController(myRouter)
	controllers.AddWebsocketController(myRouter)
}

func main() {
	log.Printf("Starting Server...")
	myRouter := mux.NewRouter().StrictSlash(true)
	AddRoutes(myRouter)

	go subscription.HandlerBroadcastMessages()

	if err := http.ListenAndServe(":8080", myRouter); err != nil {
		log.Printf("Error al iniciar el servidor: %v", err)
	}
}
