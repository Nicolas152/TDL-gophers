package subscription

import (
	"github.com/gorilla/websocket"
	"sync"
)

// Lock que protege la subcripcion de un cliente a un canal
var clientSubscriptionLock = &sync.Mutex{}

type SubscriptorInterface struct {
	Subscribe func()
	Unsubscribe func()
	GetSubscriptions func()
}

// TODO: Tiene que ser solo un map de un nivel.
type Subscriptor struct {
	Subscriptions map[string]map[string]map[*websocket.Conn]bool
}

func (s *Subscriptor) Subscribe(ws *websocket.Conn, workspaceKey string, channelKey string) {
	// Protejo la subscripcion de un cliente a un canal
	clientSubscriptionLock.Lock()
	defer clientSubscriptionLock.Unlock()
	if s.Subscriptions[workspaceKey] == nil {
		s.Subscriptions[workspaceKey] = make(map[string]map[*websocket.Conn]bool)
	}

	if s.Subscriptions[workspaceKey][channelKey] == nil {
		s.Subscriptions[workspaceKey][channelKey] = make(map[*websocket.Conn]bool)
	}

	// Agrergo la conexion websocket al mapa de conexiones
	s.Subscriptions[workspaceKey][channelKey][ws] = true
}

func (s *Subscriptor) Unsubscribe(ws *websocket.Conn, workspaceKey string, channelKey string) {
	// Protejo la des subscripcion de un cliente a un canal
	clientSubscriptionLock.Lock()
	defer clientSubscriptionLock.Unlock()
	if s.Subscriptions[workspaceKey] == nil {
		return
	}

	if s.Subscriptions[workspaceKey][channelKey] == nil {
		return
	}

	// Elimino la conexion websocket del mapa de conexiones
	delete(s.Subscriptions[workspaceKey][channelKey], ws)
}

func (s *Subscriptor) GetSubscriptions(workspaceKey string, channelKey string) map[*websocket.Conn]bool {
	// Protejo la obtencion de las conexiones websocket suscriptas a un canal
	clientSubscriptionLock.Lock()
	defer clientSubscriptionLock.Unlock()
	return s.Subscriptions[workspaceKey][channelKey]
}