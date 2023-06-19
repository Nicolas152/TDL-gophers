package controllers

import (
	"github.com/gorilla/mux"
	"gochat/src/middlewares"
	"gochat/src/services"
)

func AddWorkspaceController(myRouter *mux.Router) {
	// Handler para manipular los workspaces
	myRouter.HandleFunc("/gophers/workspace", middlewares.AuthenticationMiddleware(services.HandlerGetWorkspace)).Methods("GET")
	myRouter.HandleFunc("/gophers/workspace", middlewares.AuthenticationMiddleware(services.HandlerCreateWorkspace)).Methods("POST")

	//// Handler para manipular un workspace en particular
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}", middlewares.AuthenticationMiddleware(services.HandlerUpdateWorkspace)).Methods("PUT")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}", middlewares.AuthenticationMiddleware(services.HandlerDeleteWorkspace)).Methods("DELETE")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/join", middlewares.AuthenticationMiddleware(services.HandlerJoinWorkspace)).Methods("POST")
}
