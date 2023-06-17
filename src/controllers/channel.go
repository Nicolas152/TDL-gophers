package controllers

import (
	"gochat/src/controllers/authentication/authMiddleware"
	"gochat/src/controllers/authentication/userContext"
	"gochat/src/services/channel"
	"net/http"

	"github.com/gorilla/mux"
)

func AddChannelController(myRouter *mux.Router) {
	// Get channels by workspace
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel", authMiddleware.VerifyTokenMiddleware(getChannelsByWorkspace)).Methods("GET")
	// myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel/id", authMiddleware.VerifyTokenMiddleware(getChannelsByWorkspace)).Methods("GET")
	// myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channel", authMiddleware.VerifyTokenMiddleware(getChannelsByWorkspace)).Methods("POST")

	// myRouter.HandleFunc("/gophers/channels/{}", authMiddleware.VerifyTokenMiddleware(getChannelByKey)).Methods("GET")
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

// func getChannelByKey(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Get Channel by Key"))
// }
