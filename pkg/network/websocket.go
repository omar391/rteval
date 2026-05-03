package network

import (
	"context"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/omar391/rteval/pkg/event"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Router struct {
	mu          sync.RWMutex
	connections map[string]*websocket.Conn
}

func NewRouter() *Router {
	return &Router{
		connections: make(map[string]*websocket.Conn),
	}
}

func (r *Router) HandleConnection(w http.ResponseWriter, req *http.Request, clientID string) error {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		return err
	}
	r.mu.Lock()
	r.connections[clientID] = conn
	r.mu.Unlock()
	return nil
}

func (r *Router) Route(ctx context.Context, e event.Event[any]) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Example: Broadcaster or targeted routing logic can be added here.
	for _, conn := range r.connections {
		if err := conn.WriteJSON(e); err != nil {
			// Handle broken connections in a real scenario
		}
	}
	return nil
}
