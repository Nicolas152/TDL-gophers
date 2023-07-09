package controllers

import (
	"gochat/src/middlewares"
	"gochat/src/services/workspace"

	"github.com/gorilla/mux"
)

func AddWorkspaceController(myRouter *mux.Router) {
	// Handler to manipulate the workspaces
	myRouter.HandleFunc("/gophers/workspace", middlewares.AuthenticationMiddleware(workspace.HandlerGetWorkspace)).Methods("GET")
	myRouter.HandleFunc("/gophers/workspace", middlewares.AuthenticationMiddleware(workspace.HandlerCreateWorkspace)).Methods("POST")

	//// Handler to manipulate a specific workspace
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}", middlewares.AuthenticationMiddleware(workspace.HandlerUpdateWorkspace)).Methods("PUT")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}", middlewares.AuthenticationMiddleware(workspace.HandlerDeleteWorkspace)).Methods("DELETE")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/join", middlewares.AuthenticationMiddleware(workspace.HandlerJoinWorkspace)).Methods("POST")
}
