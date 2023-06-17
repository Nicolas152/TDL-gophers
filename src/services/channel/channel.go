package channel

import (
	"encoding/json"
	"errors"
	"gochat/src/controllers/authentication/userContext"
	"gochat/src/models/channel"
	"gochat/src/models/user"
	"gochat/src/models/workspace"
	"net/http"
)

func GetChannelsByWorkspace(workspaceKey string, userContext *userContext.UserContext) ([]byte, error, int) {

	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return nil, errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := getChannelValidations(workspaceModel, userContext); err != nil {
		return nil, err, statusErr
	}

	channels, err := channel.GetChannelsByWorkspaceId(workspaceModel.Id)
	if err != nil {
		return nil, errors.New("Error getting channels: " + err.Error()), http.StatusInternalServerError
	}

	// print channels
	for _, channel := range channels {
		println(channel.Id, channel.Name)
	}
	channelsJson, err := json.Marshal(channels)
	if err != nil {
		return nil, errors.New("Error marshalling channels: " + err.Error()), http.StatusInternalServerError
	}

	return channelsJson, err, 0
}

// getChannelValidations performs validations to determine if the user has access to the workspace.
// It returns an error and a corresponding HTTP status code based on the validation results.
func getChannelValidations(workspaceModel workspace.Workspace, userContext *userContext.UserContext) (error, int) {

	if workspaceModel.Id == 0 {
		return errors.New("Workspace does not exists"), http.StatusBadRequest
	}

	// get userId by email
	userModel := user.User{Email: (*userContext).Email}
	if err := userModel.GetContext(); err != nil {
		return errors.New("Error getting user: " + err.Error()), http.StatusInternalServerError
	}

	// validate if user is member of workspace
	exists, err2 := workspaceModel.HasMember(userModel.Id)
	if err2 != nil {
		return errors.New("Error validating if user is member of workspace: " + err2.Error()), http.StatusInternalServerError
	}
	if !exists {
		return errors.New("User is not member of workspace"), http.StatusUnauthorized
	}

	return nil, 0
}
