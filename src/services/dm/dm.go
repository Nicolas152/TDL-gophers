package dm

import (
	"encoding/json"
	"errors"
	"gochat/src/controllers/authentication/userContext"
	"gochat/src/models/dm"
	"gochat/src/models/workspace"
	"net/http"
)


func GetDMsByUserAndWorkspace(userContext *userContext.UserContext, workspaceKey string) ([]byte, error, int) {
	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return nil, errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := DMValidations(workspaceModel, userContext); err != nil {
		return nil, err, statusErr
	}

	dms, err := dm.GetDMsByUserAndWorkspace(userContext.Id, workspaceModel.Id)
	if err != nil {
		return nil, errors.New("Error getting dms: " + err.Error()), http.StatusInternalServerError
	}

	dmsJson, err := json.Marshal(dms)
	if err != nil {
		return nil, errors.New("Error marshalling dms: " + err.Error()), http.StatusInternalServerError
	}

	return dmsJson, err, 0
}

// channelValidations performs validations to determine if the user has access to the workspace.
// It returns an error and a corresponding HTTP status code based on the validation results.
func DMValidations(workspaceModel workspace.Workspace, userContext *userContext.UserContext) (error, int) {
	if workspaceModel.Id == 0 {
		return errors.New("Workspace does not exists"), http.StatusBadRequest
	}

	// validate if user is a member of workspace
	exists, err := workspaceModel.HasMember(userContext.Id);
	if err != nil {
		return errors.New("Error validating if user is member of workspace: " + err.Error()), http.StatusInternalServerError
	}

	if !exists {
		return errors.New("User is not member of workspace"), http.StatusUnauthorized
	}
	
	return nil, 0
}

func CreateDM(workspaceKey string, userContext *userContext.UserContext) (error, int) {
	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := DMValidations(workspaceModel, userContext); err != nil {
		return nil, statusErr
	}

	dmModel := dm.DM{
		WorkspaceId: workspaceModel.Id,
	}

	err = dmModel.Create()
	if err != nil {
		return errors.New("Error creating dm: " + err.Error()), http.StatusInternalServerError
	}

	return nil, 0
}

func LeaveDM(id int, workspaceKey string, userContext *userContext.UserContext) (error, int) {
	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := DMValidations(workspaceModel, userContext); err != nil {
		return nil, statusErr
	}

	dmModel := dm.DM{
		Id: id,
		WorkspaceId: workspaceModel.Id,
	}

	dmModel, err = dmModel.Get()
	if err != nil {
		return errors.New("Could not leave DM. Reason:" + err.Error()), http.StatusBadRequest
	}

	//leave dm
	if err := dmModel.Leave(userContext.Id); err != nil {
		return errors.New("Error leaving DM: " + err.Error()), http.StatusInternalServerError
	}

	return nil, 0
}

func JoinDM(id int, workspaceKey string, userContext *userContext.UserContext) (error, int) {
	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := DMValidations(workspaceModel, userContext); err != nil {
		return nil, statusErr
	}

	dmModel := dm.DM{
		Id: id,
		WorkspaceId: workspaceModel.Id,
	}

	dmModel, err = dmModel.Get()
	if err != nil {
		return errors.New("Could not join DM. Reason:" + err.Error()), http.StatusBadRequest
	}

	//join dm
	if err := dmModel.Join(userContext.Id); err != nil {
		return errors.New("Error joining DM: " + err.Error()), http.StatusInternalServerError
	}

	return nil, 0
}

/*
func MembersOfDM(id int, workspaceKey string, userContext *userContext.UserContext) ([]byte, error, int) {
	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return nil, errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := DMValidations(workspaceModel, userContext); err != nil {
		return nil, err, statusErr
	}

	dmModel := dm.DM{
		Id: id,
		WorkspaceId: workspaceModel.Id,
	}

	dmModel, err = dmModel.Get()
	if err != nil {
		return nil, errors.New("Could not get members of DM. Reason:" + err.Error()), http.StatusBadRequest
	}

	members, err := dmModel.GetMembers()
	if err != nil {
		return nil, errors.New("Error getting members of DM: " + err.Error()), http.StatusInternalServerError
	}

	membersJson, err := json.Marshal(members)
	if err != nil {
		return nil, errors.New("Error marshalling members of DM: " + err.Error()), http.StatusInternalServerError
	}

	return membersJson, err, 0
}
*/