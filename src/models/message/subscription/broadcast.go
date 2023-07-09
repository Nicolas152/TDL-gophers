package subscription

import (
	"gochat/src/models/message"
)

var BroadcastMessages = make(chan message.Message)

func handleBroadcast(subscriptor *Subscriptor, msg message.Message) {
	chatId := msg.ChatId

	// Obtain the clients subscribed to the chat
	clients := subscriptor.GetSubscriptions(chatId)
	for client := range clients {
		err := client.WriteJSON(msg)
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
