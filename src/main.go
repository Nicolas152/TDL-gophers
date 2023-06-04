package main

import (
	"gochat/src/controllers"
	"log"
	"net/http"
)

func AddRoutes() {
	controllers.AddWebsocketController()
}

func main() {
	log.Printf("Starting Server...")
	AddRoutes()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Error al iniciar el servidor: %v", err)
	}
}
