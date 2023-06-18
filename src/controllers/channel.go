package controllers

import (
	"encoding/json"
	channelDTO "gochat/src/controllers/DTOs/channel"
	"gochat/src/controllers/authentication/authMiddleware"
	"gochat/src/controllers/authentication/userContext"
	"gochat/src/services/channel"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddChannelController(myRouter *mux.Router) {
	// Get channels by workspace
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel", authMiddleware.VerifyTokenMiddleware(getChannelsByWorkspace)).Methods("GET")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel", authMiddleware.VerifyTokenMiddleware(createChannel)).Methods("POST")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel/{id}", authMiddleware.VerifyTokenMiddleware(updateChannel)).Methods("PUT")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel/{id}", authMiddleware.VerifyTokenMiddleware(deleteChannel)).Methods("DELETE")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel/{id}/join", authMiddleware.VerifyTokenMiddleware(joinToChannel)).Methods("POST")
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel/{id}/members", authMiddleware.VerifyTokenMiddleware(membersOfChannel)).Methods("POST")

}

func getChannelsByWorkspace(w http.ResponseWriter, r *http.Request) {

	// get user context
	userContext := userContext.GetUserContext(r)

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]

	channels, err, statusErr := channel.GetChannelsByWorkspace(workspaceKey, userContext)

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}

	w.Write(channels)
}

func createChannel(w http.ResponseWriter, r *http.Request) {

	// get user context
	userContext := userContext.GetUserContext(r)

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]

	var channelDTO channelDTO.ChannelDTO
	_ = json.NewDecoder(r.Body).Decode(&channelDTO)

	if channelDTO.Name == "" {
		http.Error(w, "Channel name is required", http.StatusBadRequest)
		return
	}

	err, statusErr := channel.CreateChannel(channelDTO.Name, channelDTO.Password, workspaceKey, userContext)

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Channel created successfully"))
}

func updateChannel(w http.ResponseWriter, r *http.Request) {

	// get user context
	userContext := userContext.GetUserContext(r)

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

	err, statusErr := channel.UpdateChannel(channelId, channelDTO.Name, channelDTO.Password, workspaceKey, userContext)

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Channel updated successfully"))
}

func deleteChannel(w http.ResponseWriter, r *http.Request) {

	// get user context
	userContext := userContext.GetUserContext(r)

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]
	channelId, _ := strconv.Atoi(vars["id"])

	err, statusErr := channel.DeleteChannel(channelId, workspaceKey, userContext)

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Channel deleted successfully"))
}

func joinToChannel(w http.ResponseWriter, r *http.Request) {

	// get user context
	userContext := userContext.GetUserContext(r)

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]
	channelId, _ := strconv.Atoi(vars["id"])

	var channelDTO channelDTO.ChannelDTO
	_ = json.NewDecoder(r.Body).Decode(&channelDTO)

	err, statusErr := channel.JoinToChannel(channelId, channelDTO.Password, workspaceKey, userContext)

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Channel joined successfully"))
}

func membersOfChannel(w http.ResponseWriter, r *http.Request) {

	// get user context
	userContext := userContext.GetUserContext(r)

	// get workspaceKey from URL
	vars := mux.Vars(r)
	workspaceKey := vars["workspaceKey"]
	channelId, _ := strconv.Atoi(vars["id"])

	members, err, statusErr := channel.MembersOfChannel(channelId, workspaceKey, userContext)

	if err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(members)
}
