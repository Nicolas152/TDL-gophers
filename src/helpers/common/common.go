package common

import (
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"strconv"
)

// Crea una key con el formato xxx-xxx-xxx
func CreateKey() string {
	const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const keyLength = 3

	key := ""

	for i := 0; i < keyLength; i++ {
		for j := 0; j < keyLength; j++ {
			key += string(letters[rand.Intn(len(letters))])
		}

		if i < keyLength - 1 {
			key += "-"
		}
	}

	return key
}

// Headers

func SendWelcomeMessage(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte("================== Welcolme to 'GoChat' =================="))
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


func SendErrorMessage(ws *websocket.Conn, msgErr error) {
	if err := ws.WriteMessage(websocket.TextMessage, []byte(msgErr.Error())); err != nil {
		log.Println("Error al enviar el mensaje de error: ", err)
	}
}


// Ask

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
