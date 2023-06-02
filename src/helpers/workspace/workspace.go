package workspace

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"gochat/src/models/request"
	"gochat/src/models/workspace"
)

func HandlerWorkspaceMessages(ws *websocket.Conn, userRequest request.UserRequest) error {

	switch userRequest.GetAction() {
		case request.ListAction:
			message, err := HandlerListWorkspaceMessage(userRequest)
			if err != nil {
				return err
			}

			ws.WriteMessage(websocket.TextMessage, message)

	case request.CreateAction:
		if err := HandlerCreateWorkspaceMessage(userRequest); err != nil {
			return err
		}

	case request.UpdateAction:
		if err := HandlerUpdateWorkspaceMessage(userRequest); err != nil {
			return err
		}

	case request.DeleteAction:
		if err := HandlerDeleteWorkspaceMessage(userRequest); err != nil {
			return err
		}

	default:
		return errors.New("Invalid 'action' for 'workspace' model provided")
	}

	return nil
}

func HandlerListWorkspaceMessage(userRequest request.UserRequest) ([]byte, error) {
	workspaces := workspace.Get()
	return json.Marshal(workspaces)
}


func HandlerCreateWorkspaceMessage(userRequest request.UserRequest) error {
	// Valido que el request tenga parametros
	if !userRequest.HasParameters() {
		return errors.New("Parameters are required")
	}

	// Obtengo los parametros del request
	var workspaceModel workspace.Workspace
	if err := json.Unmarshal(userRequest.GetParameters(), &workspaceModel); err != nil {
		return err
	}

	// Creo el workspace en DB
	return workspaceModel.Create(userRequest.GetUserId())
}

func HandlerUpdateWorkspaceMessage(userRequest request.UserRequest) error {
	if !userRequest.HasId() {
		return errors.New("Id is required")
	}

	if !userRequest.HasParameters() {
		return errors.New("Parameters are required")
	}

	// Obtengo el workspace actual
	currentWorkspaceModel, err := workspace.GetWorkspaceByKey(userRequest.GetId())
	if err != nil {
		return err
	}

	// Valido que el usuario sea el creador del workspace
	if currentWorkspaceModel.Creator != userRequest.GetUserId() {
		return errors.New("Only the creator of the workspace can update it")
	}

	// Obtengo los parametros a actualizar del request
	var workspaceModel workspace.Workspace
	if err := json.Unmarshal(userRequest.GetParameters(), &workspaceModel); err != nil {
		return err
	}

	// Actualizo el workspace en DB
	return currentWorkspaceModel.Update(userRequest.GetUserId(), workspaceModel)
}

func HandlerDeleteWorkspaceMessage(userRequest request.UserRequest) error {
	if !userRequest.HasId() {
		return errors.New("Id is required")
	}

	// Obtengo el workspace actual
	currentWorkspaceModel, err := workspace.GetWorkspaceByKey(userRequest.GetId())
	if err != nil {
		return err
	}

	// Valido que el usuario sea el creador del workspace
	if currentWorkspaceModel.Creator != userRequest.GetUserId() {
		return errors.New("Only the creator of the workspace can delete it")
	}

	if err := currentWorkspaceModel.Delete(userRequest.GetUserId()); err != nil {
		return err
	}
	return nil
}