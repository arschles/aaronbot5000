package http

import (
	"reflect"
)

// hub maintains the set of active clients and broadcasts messages to the clients.
type hub struct {
	// Registered clients.
	clients map[*client]bool

	// Inbound messages from the clients.
	broadcast chan WebsocketMessage

	// Register requests from the clients.
	register chan *client

	// Unregister requests from clients.
	unregister chan *client
}

func newhub() *hub {
	return &hub{
		broadcast:  make(chan WebsocketMessage),
		register:   make(chan *client),
		unregister: make(chan *client),
		clients:    make(map[*client]bool),
	}
}

func (h *hub) BroadcastMessage(msg Message) error {
	t := reflect.TypeOf(msg)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	m := WebsocketMessage{
		Type:    t.String(),
		Message: msg,
	}

	h.broadcast <- m
	return nil
}

func (h *hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
