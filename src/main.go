package main

import (
	"gochat/src/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AddRoutes(myRouter *mux.Router) {
	controllers.AddAuthenticationsController(myRouter)
	controllers.AddWorkspaceController(myRouter)

	controllers.AddAuthJWTController(myRouter)
	// controllers.AddChannelController(myRouter)
}

func main() {
	log.Printf("Starting Server...")
	myRouter := mux.NewRouter().StrictSlash(true)
	AddRoutes(myRouter)

	if err := http.ListenAndServe(":8080", myRouter); err != nil {
		log.Printf("Error al iniciar el servidor: %v", err)
	}
}
