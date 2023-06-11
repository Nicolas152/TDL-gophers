package controllers

import (
	"github.com/gorilla/mux"
	"gochat/src/helpers/authentication"
	"gochat/src/helpers/workspace"
	"net/http"
)

func AddWorkspaceController(myRouter *mux.Router)  {
	// Handler para obtener y crear un workspace
	myRouter.HandleFunc("/gophers/workspace", func(w http.ResponseWriter, r *http.Request) {
		// Valido que el usuario este logueado
		userModel := authentication.HandlerAccess(w, r)
		if userModel == nil { return }

		switch r.Method {
			case http.MethodGet:
				workspace.HandlerListWorkspace(userModel.Id)

			case http.MethodPost:
				println("Create Workspace")

			default:
				println("Method not allowed")
		}

		w.Write([]byte("Workspace"))
	})

	// Handler para obtener, actualizar y eliminar un workspace
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}", func(writer http.ResponseWriter, r *http.Request) {
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

		writer.Write([]byte("Workspace"))
	})
}
