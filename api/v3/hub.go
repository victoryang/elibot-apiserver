package v3

import (
        "fmt"
)
type Corporate interface {
        PushMsg([]byte) error 
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
        // Registered clients.
        clients map[*Client]bool

        // Inbound messages from the clients.
        broadcast chan []byte

        // Register requests from the clients.
        register chan *Client

        // Unregister requests from clients.
        unregister chan *Client
}

func (h *Hub) PushMsg(msg []byte) error {
        fmt.Println("Push messages to all client")
        h.broadcast <- msg
        return nil
}

func (h *Hub) Run() {
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

func NewHub() *Hub {
        return &Hub{
                broadcast:  make(chan []byte),
                register:   make(chan *Client),
                unregister: make(chan *Client),
                clients:    make(map[*Client]bool),
        }
}