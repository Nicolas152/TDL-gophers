package workspace

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gochat/src/models/relationship"
	"gochat/src/models/request"
	"gochat/src/models/workspace"
	"net/http"
)

func HandlerGetWorkspace(w http.ResponseWriter, r *http.Request) {
	// Cargo la request del cliente
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtengo los Workspaces del usuario
	workspaces, err := workspace.Get(userRequest.GetUserId())
	if err != nil {
		http.Error(w, "Could not get Workspaces. Reason: " + err.Error(), http.StatusInternalServerError)
		return
	}

	// Convierto los Workspaces a JSON
	data, err := json.Marshal(workspaces)
	if err != nil {
		http.Error(w, "Coudl not get Workspaces. Reason: " + err.Error(), http.StatusInternalServerError)
		return
	}

	// Envio la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func HandlerCreateWorkspace(w http.ResponseWriter, r *http.Request) {
	// Cargo la request del cliente
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, "Could not create Workspace. Reason: " + err.Error(), http.StatusBadRequest)
		return
	}

	// Valido que tenga parametros
	if !userRequest.HasParameters() {
		http.Error(w, "Could not create Workspace. Reason: Parameters are required", http.StatusBadRequest)
		return
	}

	// Obtengo los parametros del request
	var workspaceModel workspace.Workspace
	if err := json.Unmarshal(userRequest.GetParameters(), &workspaceModel); err != nil {
		http.Error(w, "Coudl not create Workspace. Reason: " + err.Error(), http.StatusBadRequest)
		return
	}

	// Creo el Workspace
	if err := workspaceModel.Create(userRequest.GetUserId()); err != nil {
		http.Error(w, "Could not create Workspace. Reason: " + err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandlerUpdateWorkspace(w http.ResponseWriter, r *http.Request) {
	// Cargo la request del cliente
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtengo el workspaceKey de las variables de la URL
	workspaceKey := mux.Vars(r)["workspaceKey"]
	currentWorkspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		http.Error(w, "Could not update Workspace. Reason: " + err.Error(), http.StatusBadRequest)
		return
	}

	// Valido que el usuario sea el owner del workspace
	if !currentWorkspaceModel.IsOwner(userRequest.GetUserId()) {
		http.Error(w, "Could not update Workspace. Reason: Not authorized", http.StatusBadRequest)
		return
	}

	// Valido que tenga parametros
	if !userRequest.HasParameters() {
		http.Error(w, "Could not update Workspace. Reason: Parameters are required", http.StatusBadRequest)
		return
	}

	// Obtengo los parametros del request
	var workspaceModel workspace.Workspace
	if err := json.Unmarshal(userRequest.GetParameters(), &workspaceModel); err != nil {
		http.Error(w, "Could not update Workspace. Reason: " + err.Error(), http.StatusBadRequest)
		return
	}

	// Actualizo el Workspace
	if err := currentWorkspaceModel.Update(userRequest.GetUserId(), workspaceModel); err != nil {
		http.Error(w, "Could not update Workspace. Reason: " + err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandlerDeleteWorkspace(w http.ResponseWriter, r *http.Request) {
	// Cargo la request del cliente
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtengo el workspaceKey de las variables de la URL
	workspaceKey := mux.Vars(r)["workspaceKey"]

	// Obtengo el workspace de interes
	currentWorkspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		http.Error(w, "Could not delete Workspace. Reason: " + err.Error(), http.StatusBadRequest)
		return
	}

	// Valido que el usuario sea el owner del workspace
	if !currentWorkspaceModel.IsOwner(userRequest.GetUserId()) {
		http.Error(w, "Could not delete Workspace. Reason: Not authorized", http.StatusBadRequest)
		return
	}

	// Elimino el Workspace
	if err := currentWorkspaceModel.Delete(userRequest.GetUserId()); err != nil {
		http.Error(w, "Could not delete Workspace. Reason: " + err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandlerJoinWorkspace(w http.ResponseWriter, r *http.Request) {
	// Cargo la request del cliente
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtengo el workspaceKey de las variables de la URL
	workspaceKey := mux.Vars(r)["workspaceKey"]
	currentWorkspaceModel := workspace.Workspace{WorkspaceKey: workspaceKey}

	// Obtengo workspace de interes
	currentWorkspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		http.Error(w, "Could not join to Workspace. Reason: " + err.Error(), http.StatusBadRequest)
		return
	}

	// Valido que el usuario tenga permisos para unirse al workspace
	if !currentWorkspaceModel.IsPublic() {
		if !userRequest.HasParameters() {
			http.Error(w, "Could not join to Workspace. Reason: Parameters are required", http.StatusBadRequest)
			return
		}

		// Obtengo los parametros del request
		var workspaceModel workspace.Workspace
		if err := json.Unmarshal(userRequest.GetParameters(), &workspaceModel); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Valido la contraseña del workspace
		if !currentWorkspaceModel.Authenticate(workspaceModel) {
			http.Error(w, "Could not join to Workspace. Reason: Not authorized", http.StatusBadRequest)
			return
		}
	}

	// Agrego al usuario al workspace
	userWorkflow := relationship.UserWorkspaceRelationship{UserId: userRequest.GetUserId(), WorkspaceId: currentWorkspaceModel.GetId()}
	if err := userWorkflow.Create(); err != nil {
		http.Error(w, "Could not join to Workspace. Reason: " + err.Error(), http.StatusInternalServerError)
		return
	}
}