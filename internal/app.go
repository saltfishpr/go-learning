// @file: app.go
// @date: 2021/11/2

package internal

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/websocket/v2"

	"learning/config"
	v1 "learning/internal/api/v1"
	"learning/logger"
)

func NewApp() *fiber.App {
	app := fiber.New()
	app.Static("/web/", "web")

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
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
		apiV1.Post("/h", v1.CreateHub)
		apiV1.Get("/h", v1.GetAllHubs)
		apiV1.Put("/h", v1.UpdateHub)
		apiV1.Delete("/h", v1.DeleteHub)
		apiV1.Get("/h/:hid", v1.GetHubInfo)
	}
	{
		apiV1.Get("/u", v1.GetAllUsers)
		apiV1.Put("/u", v1.UpdateUser)
		apiV1.Delete("/u", v1.DeleteUser)
		apiV1.Get("/u/:account", v1.GetUserInfo)

		apiV1.Post("/hubs", v1.Join)
		apiV1.Get("/hubs", v1.Joined)
		apiV1.Delete("/hubs", v1.Leave)

		apiV1.Post("/friends", v1.Follow)
		apiV1.Get("/friends", v1.Following)
		apiV1.Delete("/friends", v1.Unfollow)
	}

	ws := app.Group("/ws")
	ws.Get("/:hub", websocket.New(v1.HubHandler))
	return app
}
