package messages

import (
	"github.com/gorilla/websocket"
	"strconv"
)

// Headers

func SendWelcomeMessage(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte("================== Welcolme to 'GoChat' =================="))
}

func SendLogInMessage(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte("================== LogIn =================="))
}

func SendSignInMessage(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte("================== SignIn =================="))
}

func SendMenuMessage(ws *websocket.Conn)  {
	ws.WriteMessage(websocket.TextMessage, []byte("================== Menu =================="))
}

func SendWorkspaceMessage(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte("================== Workspace =================="))
}

// Messages

func SendConnectedMessage(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte("Conectado!"))
}


// Errors

func SendErrorEmptyMessage(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte("No se ingresó ningún mensaje..."))
}

func SendErrorInvalidUserMessage(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte("No se pudo completar la operación. Usuario inválido!"))
}

// Ask

func AskSignInMessage(ws *websocket.Conn) (string, string, string) {
	ws.WriteMessage(websocket.TextMessage, []byte("Ingrese las credenciales..."))

	email := AskEmailMessage(ws)
	name := AskNameMessage(ws)
	password := AskPasswordMessage(ws)

	return email, name, password
}

func AskLogInMessage(ws *websocket.Conn) (string, string) {
	ws.WriteMessage(websocket.TextMessage, []byte("Ingrese sus credenciales..."))

	email := AskEmailMessage(ws)
	password := AskPasswordMessage(ws)

	return email, password
}

func AskSelectWorkspaceMessage(ws *websocket.Conn) int {
	ws.WriteMessage(websocket.TextMessage, []byte("Ingrese el ID del workspace:"))
	_, msg, _ := ws.ReadMessage()

	rawWorkflowId := string(msg)
	workflowId, _ := strconv.Atoi(rawWorkflowId)

	return workflowId
}

func AskCreateWorkspaceMessage(ws *websocket.Conn) (string, string, string) {
	ws.WriteMessage(websocket.TextMessage, []byte("Ingrese los datos del workspace..."))

	name := AskNameMessage(ws)
	description := AskDescriptionMessage(ws)
	password := AskPasswordMessage(ws)

	return name, description, password
}

func AskModifyWorkspaceMessage(ws *websocket.Conn) (string, string, string) {
	ws.WriteMessage(websocket.TextMessage, []byte("Ingrese los datos del workspace..."))

	name := AskNameMessage(ws)
	description := AskDescriptionMessage(ws)
	password := AskPasswordMessage(ws)

	return name, description, password
}

func AskEmailMessage(ws *websocket.Conn) string {
	ws.WriteMessage(websocket.TextMessage, []byte("Ingrese el Email:"))
	_, msg, _ := ws.ReadMessage()

	return string(msg)
}

func AskNameMessage(ws *websocket.Conn) string {
	ws.WriteMessage(websocket.TextMessage, []byte("Ingrese el Nombre:"))
	_, msg, _ := ws.ReadMessage()

	return string(msg)
}

func AskDescriptionMessage(ws *websocket.Conn) string {
	ws.WriteMessage(websocket.TextMessage, []byte("Ingrese la Descripción:"))
	_, msg, _ := ws.ReadMessage()

	return string(msg)
}

func AskPasswordMessage(ws *websocket.Conn) string {
	ws.WriteMessage(websocket.TextMessage, []byte("Ingrese la Contraseña:"))
	_, msg, _ := ws.ReadMessage()

	return string(msg)
}
