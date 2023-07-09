package dm

import (
	"encoding/json"
	"errors"
	"gochat/src/models/chat"
	"gochat/src/models/dm"
	"gochat/src/models/user"
	"gochat/src/models/workspace"
	"net/http"
)

func GetDMsByWorkspace(userId int, workspaceKey string) ([]byte, error, int) {
	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return nil, errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := DMValidations(workspaceModel, userId); err != nil {
		return nil, err, statusErr
	}

	dms, err := dm.GetDMsByWorkspaceId(workspaceModel.Id, userId)
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
func DMValidations(workspaceModel workspace.Workspace, userId int) (error, int) {
	if workspaceModel.Id == 0 {
		return errors.New("Workspace does not exists"), http.StatusBadRequest
	}

	// validate if user is a member of workspace
	exists, err := workspaceModel.HasMember(userId)
	if err != nil {
		return errors.New("Error validating if user is member of workspace: " + err.Error()), http.StatusInternalServerError
	}

	if !exists {
		return errors.New("User is not member of workspace"), http.StatusUnauthorized
	}

	return nil, 0
}

func DMUsersValidations(workspaceModel workspace.Workspace, senderId, receiverId int) (error, int) {
	// Validar si el modelo de espacio de trabajo existe
	if workspaceModel.Id == 0 {
		return errors.New("El espacio de trabajo no existe"), http.StatusBadRequest
	}

	// Validar si el usuario remitente es miembro del espacio de trabajo
	exists, err := workspaceModel.HasMember(senderId)
	if err != nil {
		return errors.New("Error validando si el usuario es miembro del espacio de trabajo: " + err.Error()), http.StatusInternalServerError
	}

	if !exists {
		return errors.New("El usuario no es miembro del espacio de trabajo"), http.StatusUnauthorized
	}

	// Validar si el usuario receptor es miembro del espacio de trabajo
	exists, err = workspaceModel.HasMember(receiverId)
	if err != nil {
		return errors.New("Error validando si el usuario receptor es miembro del espacio de trabajo: " + err.Error()), http.StatusInternalServerError
	}

	if !exists {
		return errors.New("El usuario receptor no es miembro del espacio de trabajo"), http.StatusUnauthorized
	}

	return nil, 0
}

func CreateDM(workspaceKey string, senderID int, receiverEmail string) (error, int) {
	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	receiverID, err := user.GetUserIDByEmail(receiverEmail)

	if err != nil {
		return errors.New("Error obtaining receiver ID: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := DMUsersValidations(workspaceModel, senderID, receiverID); err != nil {
		return err, statusErr
	}

	dmModel := dm.DM{
		WorkspaceId: workspaceModel.Id,
	}

	err = dmModel.Create(senderID, receiverID)
	if err != nil {
		return errors.New("Error creating dm: " + err.Error()), http.StatusInternalServerError
	}

	return nil, 0
}

func LeaveDM(id int, workspaceKey string, userId int) (error, int) {
	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := DMValidations(workspaceModel, userId); err != nil {
		return nil, statusErr
	}

	dmModel := dm.DM{
		Id:          id,
		WorkspaceId: workspaceModel.Id,
	}

	dmModel, err = dmModel.Get()
	if err != nil {
		return errors.New("Could not leave DM. Reason:" + err.Error()), http.StatusBadRequest
	}

	//leave dm
	if err := dmModel.Leave(userId); err != nil {
		return errors.New("Error leaving DM: " + err.Error()), http.StatusInternalServerError
	}

	return nil, 0
}

func JoinDM(id int, workspaceKey string, userId int) (error, int) {
	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := DMValidations(workspaceModel, userId); err != nil {
		return nil, statusErr
	}

	dmModel := dm.DM{
		Id:          id,
		WorkspaceId: workspaceModel.Id,
	}

	dmModel, err = dmModel.Get()
	if err != nil {
		return errors.New("Could not join DM. Reason:" + err.Error()), http.StatusBadRequest
	}

	//join dm
	if err := dmModel.Join(userId); err != nil {
		return errors.New("Error joining DM: " + err.Error()), http.StatusInternalServerError
	}

	return nil, 0
}

func Messages(id int, workspaceKey string, userId int) ([]byte, error, int) {

	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return nil, errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := DMValidations(workspaceModel, userId); err != nil {
		return nil, err, statusErr
	}

	chatId, err := chat.Resolve(0, id)
	messages, err := chat.GetMessages(chatId)
	if err != nil {
		return nil, errors.New("Could not get messages of DM. Reason:" + err.Error()), http.StatusBadRequest
	}

	messagesJson, err := json.Marshal(messages)
	if err != nil {
		return nil, errors.New("Error marshalling messages: " + err.Error()), http.StatusInternalServerError
	}

	return messagesJson, nil, 0
}

/*
func MembersOfDM(id int, workspaceKey string, userId int) ([]byte, error, int) {
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
