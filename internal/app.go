// @description: 定义路由
// @file: app.go
// @date: 2021/11/2

package internal

import (
	v1 "learning/internal/api/v1"
	"learning/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func NewApp() *fiber.App {
	app := fiber.New()
	app.Static("/web/", "web")

	app.Use(middleware.Recover)
	app.Use(middleware.Timer)

	apiV1 := app.Group("/api/v1")
	{
		apiV1.Post("/register", v1.Register)
		apiV1.Post("/login", v1.Login)
	}
	apiV1.Use(middleware.JwtAuth)
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
	}

	chat := app.Group("/chat")
	chat.Get("/message", websocket.New(v1.ChatHandler))
	return app
}
