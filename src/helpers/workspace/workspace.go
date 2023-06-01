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
			workspaces := workspace.Get()
			message, _ := json.Marshal(workspaces)

			ws.WriteMessage(websocket.TextMessage, message)

		case request.CreateAction:
			if !userRequest.HasParameters() {
				return errors.New("Parameters are required")
			}

			var workspaceModel workspace.Workspace
			if err := json.Unmarshal(userRequest.GetParameters(), &workspaceModel); err != nil {
				return err
			}

			if err := workspaceModel.Create(userRequest.GetUserId()); err != nil {
				return err
			}

		case request.UpdateAction:
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

			// Obtengo los parametros a actualizar del request
			var workspaceModel workspace.Workspace
			if err := json.Unmarshal(userRequest.GetParameters(), &workspaceModel); err != nil {
				return err
			}

			if err := currentWorkspaceModel.Update(userRequest.GetUserId(), workspaceModel); err != nil {
				return err
			}

		case request.DeleteAction:
			if !userRequest.HasId() {
				return errors.New("Id is required")
			}

			// Obtengo el workspace actual
			currentWorkspaceModel, err := workspace.GetWorkspaceByKey(userRequest.GetId())
			if err != nil {
				return err
			}

			if err := currentWorkspaceModel.Delete(userRequest.GetUserId()); err != nil {
				return err
			}

		default:
			return errors.New("Invalid 'action' for 'workspace' model provided")
	}

	return nil
}
