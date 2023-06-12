package controllers

import (
	"gochat/src/controllers/authentication/authMiddleware"
	"gochat/src/controllers/authentication/userContext"
	"gochat/src/helpers/workspace"
	"gochat/src/models/user"

	"net/http"

	"github.com/gorilla/mux"
)

func AddWorkspaceController(myRouter *mux.Router) {
	// Handler para obtener y crear un workspace
	myRouter.HandleFunc("/gophers/workspace", authMiddleware.VerifyTokenMiddleware(workspaceHandler))

	// Handler para obtener, actualizar y eliminar un workspace
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}", authMiddleware.VerifyTokenMiddleware(workspaceHandlerByKey))
}

func workspaceHandler(w http.ResponseWriter, r *http.Request) {
	// userModel := authentication.HandlerAccess(w, r)
	// if userModel == nil {
	// 	return
	// }

	// Get user context
	userContext := userContext.GetUserContext(r)

	userModel := user.User{Email: (*userContext).Email}
	if err := userModel.GetContext(); err != nil {
		http.Error(w, "Error with user context: "+err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		workspace.HandlerListWorkspace(userModel.Id)

	case http.MethodPost:
		println("Create Workspace")

	default:
		println("Method not allowed")
	}

	w.Write([]byte("Workspace"))
}

func workspaceHandlerByKey(w http.ResponseWriter, r *http.Request) {
	// Obtengo el workspaceKey de la URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]

	switch r.Method {
	case http.MethodGet:
		println("Get Workspace", workspaceKey)

	case http.MethodPut:
		println("Update Workspace", workspaceKey)

	case http.MethodDelete:
		println("Delete Workspace", workspaceKey)

	default:
		println("Method not allowed")
	}

	w.Write([]byte("Workspace"))
}
