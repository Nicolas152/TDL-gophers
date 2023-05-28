package workspace

import (
	"encoding/json"
	"fiuba/concurrent/gochat/src/helpers/messages"
	"fiuba/concurrent/gochat/src/models/user"
	"fiuba/concurrent/gochat/src/models/workspace"
	"github.com/gorilla/websocket"
	"log"
)

func HandlerWorkspaceMessages(ws *websocket.Conn, user *user.User) error {
	selectedWorkspace := workspace.EmptyWorkspace()

	for {
		messages.SendWorkspaceMessage(ws)
		_, msg, _ := ws.ReadMessage()

		switch string(msg) {

			case "list":
				workspaces := workspace.Get()
				message, _ := json.Marshal(workspaces)

				ws.WriteMessage(websocket.TextMessage, message)

			case "select":
				selectedWorkspaceId := messages.AskSelectWorkspaceMessage(ws)
				selectedWorkspace, _ = workspace.GetWorkspaceById(selectedWorkspaceId)

			case "create":
				name, description, password := messages.AskCreateWorkspaceMessage(ws)
				(*user).CreateWorkspace(name, description, password)

			case "modify":
				if selectedWorkspace.IsEmpty() {
					log.Println("No se seleccionó ningún workspace. Selecciona uno con 'select'")
					continue
				}

				name, description, password := messages.AskModifyWorkspaceMessage(ws)
				(*user).ModifyWorkspace(&selectedWorkspace, name, description, password)

			case "delete":
				if selectedWorkspace.IsEmpty() {
					log.Println("No se seleccionó ningún workspace. Selecciona uno con 'select'")
					continue
				}

				(*user).DeleteWorkspace(&selectedWorkspace)
				selectedWorkspace = workspace.EmptyWorkspace()

			case "back":
				return nil

		default:
			log.Println("Opcion invalida. Ingresa 'list', 'select', 'create', 'modify', 'delete' o 'back'")

		}
	}
}
