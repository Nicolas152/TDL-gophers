package message

import (
	"gochat/src/models/chat"
	"gochat/src/models/message"
	"gochat/src/models/message/subscription"
	"gochat/src/models/request"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func HandlerMessages(w http.ResponseWriter, r *http.Request) {
	// Cargo la request del cliente
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtengo la WorkspaceKey y el ChannelKey de la request
	//TODO add validation for workspaceKey and channelKey
	//workspaceKey := mux.Vars(r)["workspaceKey"]
	channelKey, _ := strconv.Atoi(mux.Vars(r)["channelKey"])
	dmKey, _ := strconv.Atoi(mux.Vars(r)["dmKey"])

	chatId, err := chat.Resolve(channelKey, dmKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Validar que el usuario pertenezca a la workspace y al canal
	// TODO: Obtener el ChatKey a partir de la WorkspaceKey y el ChannelKey

	// Transformamos la conexion http en una conexion websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	// TODO: La subcripcion se tiene que hacer con el ChatKey, NO con el WorkspaceKey y el ChannelKey
	// Subscribo la conexion websocket al canal correspondiente
	subscriptor := subscription.GetSubscriptionInstance()
	subscriptor.Subscribe(ws, chatId)

	for {
		var msg message.Message

		// Leo el mensaje recibido
		err := ws.ReadJSON(&msg)
		if err != nil {
			subscriptor.Unsubscribe(ws, chatId)
			return
		}

		// TODO: Lo mismo que arriba. Se tiene que usar el ChatKey
		// Agrego la WorkspaceKey y el ChannelKey al mensaje
		msg.ChatId = chatId
		msg.UserId = userRequest.GetUserId()

		// save the message in the database
		err = msg.Save()
		if err != nil {
			println(err.Error())
		}

		// Envio el mensaje a todos los clientes subscriptos al canal
		subscription.BroadcastMessages <- msg
	}
}
