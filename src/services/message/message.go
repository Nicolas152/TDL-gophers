package message

import (
	"errors"
	"gochat/src/models/chat"
	"gochat/src/models/message"
	"gochat/src/models/message/subscription"
	"gochat/src/models/request"
	"gochat/src/models/user"
	"gochat/src/services/channel"
	"gochat/src/services/dm"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func HandlerMessages(w http.ResponseWriter, r *http.Request) {
	// Load the client's request
	var userRequest request.UserRequest
	if err := userRequest.ReadRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the WorkspaceKey, ChannelKey and DmKey from the request
	workspaceKey := mux.Vars(r)["workspaceKey"]
	channelKey, _ := strconv.Atoi(mux.Vars(r)["channelKey"])
	dmKey, _ := strconv.Atoi(mux.Vars(r)["dmKey"])

	err := validateChat(workspaceKey, channelKey, dmKey, userRequest.GetUserId())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Resolve the chatId
	chatId, err := chat.Resolve(channelKey, dmKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Transform the HTTP connection into a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	// Subscribe the WebSocket connection to the corresponding channel
	subscriber := subscription.GetSubscriptionInstance()
	subscriber.Subscribe(ws, chatId)

	for {
		var msg message.Message

		// Read the received message
		err := ws.ReadJSON(&msg)
		if err != nil {
			subscriber.Unsubscribe(ws, chatId)
			return
		}

		// Add the WorkspaceKey and ChannelKey to the message
		msg.ChatId = chatId
		msg.UserId = userRequest.GetUserId()
		user, err := user.GetUserById(msg.UserId)
		if err != nil {
			println(err.Error())
		} else {
			msg.UserName = user.Name
			msg.UserEmail = user.Email
		}

		// Save the message in the database
		err = msg.Save()
		if err != nil {
			println(err.Error())
		}

		// Send the message to all clients subscribed to the channel
		subscription.BroadcastMessages <- msg
	}
}

func validateChat(workspaceKey string, channelKey, dmKey int, userId int) error {
	if channelKey == 0 && dmKey == 0 {
		return errors.New("ChannelKey or DmKey must be provided")
	}

	if workspaceKey == "" {
		return errors.New("WorkspaceKey must be provided")
	}

	if channelKey != 0 {
		_, err, _ := channel.GetChannel(channelKey, workspaceKey, userId)
		return err
	} else {
		_, err, _ := dm.GetDM(dmKey, workspaceKey, userId)
		return err
	}
}
