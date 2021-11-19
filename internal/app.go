// @file: app.go
// @date: 2021/11/2

package internal

import (
	"time"

	"learning/config"
	"learning/internal/api/v1/hubs"
	"learning/internal/api/v1/users"
	"learning/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func NewApp() *fiber.App {
	app := fiber.New()
	app.Static("/web/", "web")
	app.Use(func(c *fiber.Ctx) error {
		begin := time.Now().UnixNano()
		defer func() {
			end := time.Now().UnixNano()
			logger.Infof("url: %s%s%s use: %s%dns%s", config.BlueBold, c.Path(), config.Reset, config.Green, end-begin, config.Reset)
		}()
		return c.Next()
	})

	api := app.Group("/api")

	hubsV1 := api.Group("/v1/hubs")
	{
		hubsV1.Post("", hubs.Create)
		hubsV1.Get("", hubs.Read)
		hubsV1.Put("", hubs.Update)
		hubsV1.Delete("", hubs.Delete)
	}

	usersV1 := api.Group("/v1/users")
	{
		usersV1.Post("", users.Create)
		usersV1.Get("", users.Read)
		usersV1.Get("/all", users.ReadAll)
		usersV1.Put("", users.Update)
		usersV1.Delete("", users.Delete)
	}

	ws := app.Group("/ws")
	ws.Get("/:hubs", hubs.EnterHubHandler, websocket.New(hubs.HubHandler))
	return app
}
