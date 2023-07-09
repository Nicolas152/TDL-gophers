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
	// Print a log message indicating the server is starting
	log.Printf("Starting Server...")

	// Create a new router using mux package
	myRouter := mux.NewRouter().StrictSlash(true)

	// Call the function AddRoutes to add routes to the router
	AddRoutes(myRouter)

	// Start a goroutine to handle broadcasting messages for subscription
	go subscription.HandlerBroadcastMessages()

	// Listen and serve HTTP requests on port 8080 using the created router
	if err := http.ListenAndServe(":8080", myRouter); err != nil {
		// If there's an error starting the server, print an error log message
		log.Printf("Error al iniciar el servidor: %v", err)
	}
}
