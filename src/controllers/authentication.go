package controllers

import (
	"github.com/gorilla/mux"
	"gochat/src/helpers/authentication"
	"net/http"
)

func AddAuthenticationsController(myRouter *mux.Router)  {
	// Handler para manejar el signin
	myRouter.HandleFunc("/gophers/signin", func(w http.ResponseWriter, r *http.Request) {
		authentication.HandlerSignIn(w, r)
	})

	// Handler para manejar el login
	myRouter.HandleFunc("/gophers/login", func(w http.ResponseWriter, r *http.Request) {
		authentication.HandlerLogIn(w, r)
	})
}
