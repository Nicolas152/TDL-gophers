package subscription

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Lock que protege la subcripcion de un cliente a un canal
var clientSubscriptionLock = &sync.Mutex{}

type SubscriptorInterface struct {
	Subscribe        func()
	Unsubscribe      func()
	GetSubscriptions func()
}

// TODO: Tiene que ser solo un map de un nivel.
type Subscriptor struct {
	Subscriptions map[int]map[*websocket.Conn]bool
}

func (s *Subscriptor) Subscribe(ws *websocket.Conn, chatId int) {
	// Protejo la subscripcion de un cliente a un canal
	clientSubscriptionLock.Lock()
	defer clientSubscriptionLock.Unlock()
	if s.Subscriptions[chatId] == nil {
		s.Subscriptions[chatId] = make(map[*websocket.Conn]bool)
	}

	// Agrergo la conexion websocket al mapa de conexiones
	s.Subscriptions[chatId][ws] = true
}

func (s *Subscriptor) Unsubscribe(ws *websocket.Conn, chatId int) {
	// Protejo la des subscripcion de un cliente a un canal
	clientSubscriptionLock.Lock()
	defer clientSubscriptionLock.Unlock()
	if s.Subscriptions[chatId] == nil {
		return
	}

	// Elimino la conexion websocket del mapa de conexiones
	delete(s.Subscriptions[chatId], ws)
}

func (s *Subscriptor) GetSubscriptions(chatId int) map[*websocket.Conn]bool {
	// Protejo la obtencion de las conexiones websocket suscriptas a un canal
	clientSubscriptionLock.Lock()
	defer clientSubscriptionLock.Unlock()
	return s.Subscriptions[chatId]
}
