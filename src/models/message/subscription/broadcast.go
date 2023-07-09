package subscription

import (
	"gochat/src/models/message"
	"time"
)

var BroadcastMessages = make(chan message.Message)

func handleBroadcast(subscriptor *Subscriptor, msg message.Message) {
	chatId := msg.ChatId
	// client message
	clientMsg := message.ClientMessage{Message: msg.Message, UserEmail: msg.UserEmail, UserName: msg.UserName, CreatedAt: time.Now().Format("2006-01-02 15:04:05")}

	// Obtain the clients subscribed to the chat
	clients := subscriptor.GetSubscriptions(chatId)
	for client := range clients {
		err := client.WriteJSON(clientMsg)
		if err != nil {
			subscriptor.Unsubscribe(client, chatId)
		}
	}
}

func HandlerBroadcastMessages() {
	subscriptor := GetSubscriptionInstance()

	for {
		// Wait for a message to be broadcasted
		msg := <-BroadcastMessages

		// Handle the broadcast
		go handleBroadcast(subscriptor, msg)
	}
}
