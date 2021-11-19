// @description: 聊天室数据结构
// @file: model.go
// @date: 2021/11/1

package model

import (
	"learning/config"
	"learning/logger"

	"github.com/gofiber/websocket/v2"
)

type Hub struct {
	ID   string
	Size string

	clients    map[*websocket.Conn]Client
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	broadcast  chan []byte
}

func NewHub(id string) *Hub {
	return &Hub{
		ID: id,

		clients:    make(map[*websocket.Conn]Client),
		register:   make(chan *websocket.Conn, config.BufferedChan),
		unregister: make(chan *websocket.Conn, config.BufferedChan),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case message := <-h.broadcast:
			logger.Infof("[%s] message received: %s", h.ID, string(message))
			for connection := range h.clients {
				if err := connection.WriteMessage(websocket.TextMessage, message); err != nil {
					logger.Error("write error:", err)
					connection.WriteMessage(websocket.CloseMessage, []byte{})
					connection.Close()
					delete(h.clients, connection)
				}
			}

		case connection := <-h.register:
			h.clients[connection] = Client{}
			logger.Infof("[%s] connection registered", h.ID)

		case connection := <-h.unregister:
			delete(h.clients, connection)
			logger.Infof("[%s] connection unregistered", h.ID)
		}
	}
}

func (h *Hub) Register(c *websocket.Conn) {
	h.register <- c
}

func (h *Hub) Unregister(c *websocket.Conn) {
	h.unregister <- c
	c.Close()
}

func (h *Hub) Broadcast(payload []byte) {
	h.broadcast <- payload
}

func (h *Hub) Close() {
	for c := range h.clients {
		h.Unregister(c)
	}
}
