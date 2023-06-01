package messages

import (
	"errors"
	"github.com/gorilla/websocket"
	"gochat/src/helpers/common"
	"gochat/src/helpers/workspace"
	"gochat/src/models/request"
	"gochat/src/models/user"
)

func HandlerMessages(ws *websocket.Conn, user *user.User) error {
	var userRequest request.UserRequest
	userRequest.LoadUserId(user.Id)

	for {
		common.SendMenuMessage(ws)

		// Recibo el mensaje con la acción a realizar
		if err := userRequest.ReadRequest(ws); err != nil {
			return err
		}

		switch userRequest.GetModel() {
			case request.WorkspaceModel:
				if err := workspace.HandlerWorkspaceMessages(ws, userRequest); err != nil {
					return err
				}

			default:
				return errors.New("Invalid 'model' option provided")
		}
	}
}