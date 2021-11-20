// @file: app.go
// @date: 2021/11/2

package internal

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/websocket/v2"

	"learning/config"
	v1 "learning/internal/api/v1"
	"learning/logger"
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

	apiV1 := app.Group("/api/v1")
	{
		apiV1.Post("/register", v1.Register)
		apiV1.Post("/login", v1.Login)
	}

	apiV1.Use(jwtware.New(jwtware.Config{
		AuthScheme: config.AuthScheme,
		ContextKey: config.ContextKey,
		SigningKey: []byte(config.SigningKey),
	}))

	{
		apiV1.Post("/hubs", v1.CreateHub)
		apiV1.Get("/hubs", v1.GetAllHubs)
		apiV1.Put("/hubs", v1.UpdateHub)
		apiV1.Delete("/hubs", v1.DeleteHub)

		apiV1.Post("/hubs/:hid", v1.JoinHub)
		apiV1.Get("/hubs/:hid", v1.GetHubInfo)
		apiV1.Delete("/hubs/:hid", v1.LeaveHub)
	}

	{
		apiV1.Get("/users", v1.GetAllUsers)
		apiV1.Put("/users", v1.UpdateUser)
		apiV1.Delete("/users", v1.DeleteUser)
		apiV1.Get("/users/:account", v1.GetUserInfo)
	}

	ws := app.Group("/ws")
	ws.Get("/:hub", websocket.New(v1.HubHandler))
	return app
}
