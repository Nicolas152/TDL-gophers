package controllers

import (
	"encoding/json"
	dmDTO "gochat/src/controllers/DTOs/dm"
	"gochat/src/middlewares"
	"gochat/src/models/request"
	"gochat/src/services/dm"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddDMController(myRouter *mux.Router) {
	// Get dms by workspace
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/dm", middlewares.AuthenticationMiddleware(getDMsByWorkspace)).Methods("GET")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/dm", middlewares.AuthenticationMiddleware(createDM)).Methods("POST")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/dm/{id}/messages", middlewares.AuthenticationMiddleware(messagesDM)).Methods("GET")
}

func getDMsByWorkspace(w http.ResponseWriter, r *http.Request) {
	// Cargo la request del cliente
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]

	dms, err, statusErr := dm.GetDMsByWorkspace(userRequest.GetUserId(), workspaceKey)

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}

	w.Write(dms)
}

func createDM(w http.ResponseWriter, r *http.Request) {
	// Cargo la request del cliente
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]

	var dmDTO dmDTO.DmDTO
	_ = json.NewDecoder(r.Body).Decode(&dmDTO)

	if dmDTO.Email == "" {
		http.Error(w, "Receiver email is required", http.StatusBadRequest)
		return
	}

	err, statusErr := dm.CreateDM(workspaceKey, userRequest.GetUserId(), dmDTO.Email)

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("DM created successfully"))
}

func messagesDM(w http.ResponseWriter, r *http.Request) {

	// get user context
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]
	dmId, _ := strconv.Atoi(vars["id"])

	messages, err, statusErr := dm.Messages(dmId, workspaceKey, userRequest.GetUserId())

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(messages)
}
