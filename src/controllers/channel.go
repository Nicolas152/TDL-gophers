package controllers

import (
	"errors"
	"gochat/src/middlewares"
	"gochat/src/models/request"
	"gochat/src/models/workspace"
	"net/http"

	"github.com/gorilla/mux"
)

func AddChannelController(myRouter *mux.Router) {
	// Get channels by workspace
	myRouter.HandleFunc("/gophers/workspace/{workspaceKey}/channels", middlewares.AuthenticationMiddleware(getChannelsByWorkspace)).Methods("GET")

	//myRouter.HandleFunc("/gophers/channels/{channelKey}", authMiddleware.AuthenticationMiddleware(getChannelByKey)).Methods("GET")
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

	if err, statusErr := getChannelValidations(workspaceKey, userRequest.GetUserId()); err != nil {
		http.Error(w, err.Error(), statusErr)
		return
	}

	// get channels by workspace
	println("Get Channels by Workspace", workspaceKey, userRequest.GetUserId())

	w.Write([]byte("Get Channels"))
}

// TODO: @Lescalante14 move this to a service
// getChannelValidations performs validations to determine if the user has access to the workspace.
// It returns an error and a corresponding HTTP status code based on the validation results.
func getChannelValidations(workspaceKey string, userId int) (error, int) {
	// validate if workspaceModel exists
	// workspaceModel := workspace.Workspace{WorkflowKey: workspaceKey}
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}
	if workspaceModel.Id == 0 {
		return errors.New("Workspace does not exists"), http.StatusBadRequest
	}

	// validate if user is member of workspace
	exists, err2 := workspaceModel.HasMember(userId)
	if err2 != nil {
		return errors.New("Error validating if user is member of workspace: " + err2.Error()), http.StatusInternalServerError
	}
	if !exists {
		return errors.New("User is not member of workspace"), http.StatusForbidden
	}

	return nil, 0
}

// func getChannelByKey(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Get Channel by Key"))
// }
