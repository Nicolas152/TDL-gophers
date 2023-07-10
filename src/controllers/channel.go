package controllers

import (
	"encoding/json"
	channelDTO "gochat/src/controllers/DTOs/channel"
	"gochat/src/middlewares"
	"gochat/src/models/request"
	"gochat/src/services/channel"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddChannelController(myRouter *mux.Router) {
	// Get channels by workspace
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel", middlewares.AuthenticationMiddleware(getChannelsByWorkspace)).Methods("GET")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel", middlewares.AuthenticationMiddleware(createChannel)).Methods("POST")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel/{id}", middlewares.AuthenticationMiddleware(updateChannel)).Methods("PUT")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel/{id}", middlewares.AuthenticationMiddleware(deleteChannel)).Methods("DELETE")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel/{id}/join", middlewares.AuthenticationMiddleware(joinToChannel)).Methods("POST")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel/{id}/members", middlewares.AuthenticationMiddleware(membersOfChannel)).Methods("GET")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel/{id}/leave", middlewares.AuthenticationMiddleware(leaveChannel)).Methods("POST")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel/{id}/messages", middlewares.AuthenticationMiddleware(messages)).Methods("GET")

}

func getChannelsByWorkspace(w http.ResponseWriter, r *http.Request) {
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]

	channels, err, statusErr := channel.GetChannelsByWorkspace(workspaceKey, userRequest.GetUserId())

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(channels)
}

func createChannel(w http.ResponseWriter, r *http.Request) {

	// get user context
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]

	var channelDTO channelDTO.ChannelDTO
	_ = json.NewDecoder(r.Body).Decode(&channelDTO)

	if channelDTO.Name == "" {
		http.Error(w, "Channel name is required", http.StatusBadRequest)
		return
	}

	err, statusErr := channel.CreateChannel(channelDTO.Name, channelDTO.Password, workspaceKey, userRequest.GetUserId())

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Channel created successfully"))
}

func updateChannel(w http.ResponseWriter, r *http.Request) {

	// get user context
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]
	channelId, _ := strconv.Atoi(vars["id"])

	var channelDTO channelDTO.ChannelDTO
	_ = json.NewDecoder(r.Body).Decode(&channelDTO)

	if channelDTO.Name == "" {
		http.Error(w, "Channel name is required", http.StatusBadRequest)
		return
	}

	err, statusErr := channel.UpdateChannel(channelId, channelDTO.Name, channelDTO.Password, workspaceKey, userRequest.GetUserId())

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Channel updated successfully"))
}

func deleteChannel(w http.ResponseWriter, r *http.Request) {

	// get user context
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]
	channelId, _ := strconv.Atoi(vars["id"])

	err, statusErr := channel.DeleteChannel(channelId, workspaceKey, userRequest.GetUserId())

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Channel deleted successfully"))
}

func joinToChannel(w http.ResponseWriter, r *http.Request) {

	// get user context
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]
	channelId, _ := strconv.Atoi(vars["id"])

	var channelDTO channelDTO.ChannelDTO
	_ = json.NewDecoder(r.Body).Decode(&channelDTO)

	err, statusErr := channel.JoinToChannel(channelId, channelDTO.Password, workspaceKey, userRequest.GetUserId())

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Channel joined successfully"))
}

func membersOfChannel(w http.ResponseWriter, r *http.Request) {

	// get user context
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]
	channelId, _ := strconv.Atoi(vars["id"])

	members, err, statusErr := channel.MembersOfChannel(channelId, workspaceKey, userRequest.GetUserId())

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(members)
}

func leaveChannel(w http.ResponseWriter, r *http.Request) {

	// get user context
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]
	channelId, _ := strconv.Atoi(vars["id"])

	err, statusErr := channel.LeaveChannel(channelId, workspaceKey, userRequest.GetUserId())

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Channel left successfully"))
}

func messages(w http.ResponseWriter, r *http.Request) {

	// get user context
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]
	channelId, _ := strconv.Atoi(vars["id"])

	messages, err, statusErr := channel.Messages(channelId, workspaceKey, userRequest.GetUserId())

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(messages)
}
