package subscription

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Lock that protects the creation of the instance
var clientSubscriptionLock = &sync.Mutex{}

type SubscriptorInterface struct {
	Subscribe        func()
	Unsubscribe      func()
	GetSubscriptions func()
}

type Subscriptor struct {
	Subscriptions map[int]map[*websocket.Conn]bool
}

func (s *Subscriptor) Subscribe(ws *websocket.Conn, chatId int) {
	// Protect the subscription of a client to a channel
	clientSubscriptionLock.Lock()
	defer clientSubscriptionLock.Unlock()
	if s.Subscriptions[chatId] == nil {
		s.Subscriptions[chatId] = make(map[*websocket.Conn]bool)
	}

	// Add the websocket connection to the map of connections
	s.Subscriptions[chatId][ws] = true
}

func (s *Subscriptor) Unsubscribe(ws *websocket.Conn, chatId int) {
	// Protect the unsubscription of a client to a channel
	clientSubscriptionLock.Lock()
	defer clientSubscriptionLock.Unlock()
	if s.Subscriptions[chatId] == nil {
		return
	}

	// Delete the websocket connection from the map of connections
	delete(s.Subscriptions[chatId], ws)
}

func (s *Subscriptor) GetSubscriptions(chatId int) map[*websocket.Conn]bool {
	// Protect the access to the map of connections
	clientSubscriptionLock.Lock()
	defer clientSubscriptionLock.Unlock()
	return s.Subscriptions[chatId]
}
