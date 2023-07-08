package subscription

import "gochat/src/models/message"

var BroadcastMessages = make(chan message.Message)

func HandlerBroadcastMessages() {
	subscriptor := GetSubscriptionInstance()

	for {
		// Espero a que llegue un mensaje al canal
		msg := <-BroadcastMessages

		// TODO: Pinta que esto sera un cuello de botella. Meterlo dentro de una goroutina y que se encargue de enviar el mensaje a los clientes
		// Obtengo la WorkspaceKey y el ChannelKey del mensaje
		chatId := msg.ChatId

		// Obtengo las conexiones websocket suscriptas al canal
		clients := subscriptor.GetSubscriptions(chatId)

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				subscriptor.Unsubscribe(client, chatId)
			}
		}
	}
}
