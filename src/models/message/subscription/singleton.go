package subscription

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Lock that protects the creation of the instance
var subscriptionInstanceLock = &sync.Mutex{}

// Reference to the subscriber instance
var subscriptionInstance *Subscriptor

// Singleton that returns the subscriber instance
func GetSubscriptionInstance() *Subscriptor {
	if subscriptionInstance == nil {
		subscriptionInstanceLock.Lock()
		defer subscriptionInstanceLock.Unlock()

		if subscriptionInstance == nil {
			subscriptionInstance = &Subscriptor{
				Subscriptions: make(map[int]map[*websocket.Conn]bool),
			}
		}
	}
	return subscriptionInstance
}
