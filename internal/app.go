// @description: 定义路由
// @file: app.go
// @date: 2021/11/2

package internal

import (
	v1 "learning/internal/api/v1"
	"learning/internal/middleware"
	"learning/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// @title           Chat App API
// @version         1.0
// @description     This is the chat app server.

// @license.name  MIT

func NewApp() *fiber.App {
	app := fiber.New(
		fiber.Config{
			JSONEncoder: utils.JsonMarshal,
			JSONDecoder: utils.JsonUnmarshal,
		},
	)
	app.Use(middleware.Recover, middleware.Pprof, middleware.Logger)
	app.Use(middleware.CORS)
	app.Post("/login", v1.Login)
	app.Get("/refresh", v1.Refresh)

	apiV1 := app.Group("/api/v1")
	apiV1.Use(middleware.JwtAuth)
	apiV1.Use(middleware.Cache)
	{
		apiV1.Post("/h", v1.CreateHub)
		apiV1.Get("/h", v1.GetAllHubs)
		apiV1.Put("/h", v1.UpdateHub)
		apiV1.Delete("/h", v1.DeleteHub)
		apiV1.Get("/h/:hid", v1.GetHubInfo)

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

		apiV1.Get("/chat", v1.ChatAuth)
	}

	app.Get("/ws", middleware.WebSocket, websocket.New(v1.ChatHandler))
	return app
}
