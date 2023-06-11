package main

import (
	"github.com/gorilla/mux"
	"gochat/src/controllers"
	"log"
	"net/http"
)

func AddRoutes(myRouter *mux.Router)  {
	controllers.AddAuthenticationsController(myRouter)
	controllers.AddWorkspaceController(myRouter)
}


func main()  {
	log.Printf("Starting Server...")
	myRouter := mux.NewRouter().StrictSlash(true)
	AddRoutes(myRouter)

	if err := http.ListenAndServe(":8080", myRouter); err != nil {
		log.Printf("Error al iniciar el servidor: ", err)
	}
}
