package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/websocket/v2"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)

	// Create a new hub
	hub := NewHub()
	go hub.Start(cancelCtx)

	app := fiber.New()
	app.Use(pprof.New())
	app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	// Upgraded websocket request
	app.Get("/ws", websocket.New(hub.Handle))

	// ws://localhost:3000/ws
	go func() {
		if err := app.Listen("localhost:3000"); err != nil {
			log.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	select {
	case <-shutdown(app):
		log.Println("server shutdown")
		cancel()
	case <-time.After(5 * time.Second):
		log.Println("server shutdown timeout")
	}
}

func shutdown(app *fiber.App) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		if err := app.Shutdown(); err != nil {
			log.Println("shutdown error:", err)
		}
	}()
	return done
}

type Hub struct {
	registerCh   chan *websocket.Conn
	unregisterCh chan *websocket.Conn
	broadcastCh  chan []byte

	conns map[*websocket.Conn]bool
}

func NewHub() *Hub {
	return &Hub{
		registerCh:   make(chan *websocket.Conn),
		unregisterCh: make(chan *websocket.Conn),
		broadcastCh:  make(chan []byte),
		conns:        make(map[*websocket.Conn]bool),
	}
}

func (h *Hub) Start(ctx context.Context) {
	// broadcast message every 3 seconds
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				payload, _ := json.Marshal(map[string]string{
					"event":   "CHAT_MSG",
					"user":    "server",
					"message": "hello, this is server",
				})
				h.broadcastCh <- payload
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			for conn := range h.conns {
				conn.Close()
			}
			return
		case conn := <-h.registerCh:
			h.conns[conn] = true
		case conn := <-h.unregisterCh:
			delete(h.conns, conn)
		case message := <-h.broadcastCh:
			for conn := range h.conns {
				if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
					log.Println("write error:", err)
				}
			}
		}
	}
}

func (h *Hub) Handle(c *websocket.Conn) {
	// When the function returns, unregister the client and close the connection
	defer func() {
		h.unregisterCh <- c
		c.Close()
	}()

	// Register the client
	h.registerCh <- c

	for {
		messageType, payload, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("read error:", err)
			}
			return
		}

		switch messageType {
		case websocket.TextMessage:
			log.Printf("websocket payload received: %s", payload)
		}
	}
}
