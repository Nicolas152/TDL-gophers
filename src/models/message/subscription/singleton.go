package subscription

import (
	"github.com/gorilla/websocket"
	"sync"
)

// Lock que protege la creacion de la instancia
var subscriptionInstanceLock = &sync.Mutex{}

// Referencia a la instancia de subcriptor
var subscriptionInstance *Subscriptor

// Singleton que devuelve la instancia de subcriptor
func GetSubscriptionInstance() *Subscriptor {
	if subscriptionInstance == nil {
		subscriptionInstanceLock.Lock()
		defer subscriptionInstanceLock.Unlock()

		if subscriptionInstance == nil {
			subscriptionInstance = &Subscriptor{
				Subscriptions: make(map[string]map[string]map[*websocket.Conn]bool),
			}
		}
	}
	return subscriptionInstance
}

