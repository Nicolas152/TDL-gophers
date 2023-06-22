package controllers

import (
	"github.com/gorilla/mux"
	"gochat/src/middlewares"
	"gochat/src/services/workspace"
)

func AddWorkspaceController(myRouter *mux.Router) {
	// Handler para manipular los workspaces
	myRouter.HandleFunc("/gophers/workspace", middlewares.AuthenticationMiddleware(workspace.HandlerGetWorkspace)).Methods("GET")
	myRouter.HandleFunc("/gophers/workspace", middlewares.AuthenticationMiddleware(workspace.HandlerCreateWorkspace)).Methods("POST")

	//// Handler para manipular un workspace en particular
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}", middlewares.AuthenticationMiddleware(workspace.HandlerUpdateWorkspace)).Methods("PUT")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}", middlewares.AuthenticationMiddleware(workspace.HandlerDeleteWorkspace)).Methods("DELETE")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/join", middlewares.AuthenticationMiddleware(workspace.HandlerJoinWorkspace)).Methods("POST")
}
