package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))

	hub := NewHub(logger)
	go hub.Start(ctx)

	app := fiber.New()
	app.Use(pprof.New())
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})
	app.Get("/ws", websocket.New(hub.HandleConn, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
		RecoverHandler: func(conn *websocket.Conn) {
			if err := recover(); err != nil {
				_ = conn.WriteJSON(fiber.Map{
					"error": fmt.Sprintf("%+v", err),
				})
			}
		},
	}))

	go func() {
		if err := app.Listen("localhost:3000"); err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()

	if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
		logger.Error("shutdown error", slog.Any("err", err))
	}
}

type Client struct {
	logger *slog.Logger

	hub *Hub

	conn     *websocket.Conn
	userId   string
	deviceId string
}

func (c *Client) Listen() {
	for {
		messageType, payload, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.logger.Error("read error", slog.Any("err", err))
			}
			return
		}

		switch messageType {
		case websocket.PingMessage:
			if err := c.conn.WriteMessage(websocket.PongMessage, nil); err != nil {
				c.logger.Error("write error", slog.Any("err", err))
			}
		case websocket.TextMessage:
			c.logger.Info("websocket payload received", slog.String("payload", string(payload)))
			// TODO 给单个人发消息
			// TODO 给群组发消息
		}
	}
}

type Hub struct {
	logger *slog.Logger

	registerCh   chan *Client
	unregisterCh chan *Client

	clients map[*Client]struct{}
}

func NewHub(logger *slog.Logger) *Hub {
	return &Hub{
		registerCh:   make(chan *Client),
		unregisterCh: make(chan *Client),

		clients: make(map[*Client]struct{}),
	}
}

func (h *Hub) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // 服务退出
			h.shutdown()
			return
		case client := <-h.registerCh:
			h.register(client)
		case client := <-h.unregisterCh:
			h.unregister(client)
		}
	}
}

func (h *Hub) shutdown() {
	for client := range h.clients {
		client.conn.Close()
	}
}

func (h *Hub) register(client *Client) {
	h.clients[client] = struct{}{}
}

func (h *Hub) unregister(client *Client) {
	client.conn.Close()
	delete(h.clients, client)
}

// HandleConn handles websocket connections.
func (h *Hub) HandleConn(conn *websocket.Conn) {
	var userId string
	var deviceId string

	client := &Client{
		logger:   h.logger.With(slog.String("userId", userId), slog.String("deviceId", deviceId)),
		hub:      h,
		conn:     conn,
		userId:   userId,
		deviceId: deviceId,
	}

	defer func() {
		h.unregisterCh <- client
	}()
	h.registerCh <- client

	client.Listen()
}

func (h *Hub) SendMessageToUser(userId string, message any) error {
	panic("implement me")
}

func (h *Hub) SendMessageToGroup(groupId string, message any) error {
	panic("implement me")
}
