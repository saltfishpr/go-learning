// @file: hubs.go
// @date: 2021/11/18

// Package hubs 处理 hubs 接口相关的 http 请求
package hubs

import (
	"fmt"
	"log"
	"net/http"

	"learning/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/tidwall/gjson"
)

var hubs = make(map[string]*model.Hub)

func Create(c *fiber.Ctx) error {
	id := gjson.GetBytes(c.Body(), "id").String()
	if _, ok := hubs[id]; ok {
		return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("hub %s already exist", id))
	}
	h := model.NewHub(id)
	hubs[id] = h
	go h.Run()
	return c.SendStatus(http.StatusOK)
}

func Read(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("%v", hubs))
}

func Update(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func Delete(c *fiber.Ctx) error {
	_ = hubs[gjson.GetBytes(c.Body(), "id").String()]
	return c.SendStatus(http.StatusOK)
}

func EnterHubHandler(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}
	return c.Render("web/template/home.html", fiber.Map{"hub": c.Params("hubs")})
}

func HubHandler(c *websocket.Conn) {
	h := hubs[c.Params("hubs")]
	defer func() {
		h.Unregister(c)
	}()

	h.Register(c)

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure,
			) {
				log.Println("read error:", err)
			}
			return // Calls the deferred function, i.e. closes the connection on error
		}

		if messageType == websocket.TextMessage {
			h.Broadcast(message)
		} else {
			log.Println("websocket message received of type", messageType)
		}
	}
}
