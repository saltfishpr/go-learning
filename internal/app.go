// @description: 定义路由
// @file: app.go
// @date: 2021/11/2

package internal

import (
	"learning/internal/api/v1/auth"
	"learning/internal/api/v1/chat"
	"learning/internal/api/v1/hub"
	"learning/internal/api/v1/user"
	"learning/internal/constant"
	"learning/internal/data"
	"learning/internal/log"
	"learning/internal/middleware"
	"learning/internal/service"
	"learning/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func NewApp(logger *log.Logger, connection data.Connection) *fiber.App {
	app := fiber.New(
		fiber.Config{
			JSONEncoder: utils.JsonMarshal,
			JSONDecoder: utils.JsonUnmarshal,
			ReadTimeout: constant.ReadTimeout,
		},
	)

	app.Use(middleware.Recover, middleware.Pprof, middleware.Logger(logger))
	app.Use(middleware.CORS)

	AuthHandler := auth.New(service.NewUser(connection), logger)
	UserHandler := user.New(service.NewUser(connection), logger)
	HubHandler := hub.New(service.NewHub(connection), logger)
	ChatHandler := chat.New(logger)

	app.Post("/auth/login", AuthHandler.Login)          // 登录/注册
	app.Get("/auth/token/refresh", AuthHandler.Refresh) // 刷新token

	apiV1 := app.Group("/api/v1")
	apiV1.Use(middleware.JwtAuth)
	apiV1.Use(middleware.Cache)
	{
		apiV1.Post("/hubs", HubHandler.CreateHub)        // 创建群组
		apiV1.Get("/hubs/:hid", HubHandler.GetHub)       // 获取群组信息
		apiV1.Put("/hubs/:hid", HubHandler.UpdateHub)    // 更新群组信息
		apiV1.Delete("/hubs/:hid", HubHandler.DeleteHub) // 删除群组
		apiV1.Get("/hubs/items", HubHandler.GetHubs)     // 获取群组列表

		apiV1.Post("/users", UserHandler.CreateUser)             // 创建用户
		apiV1.Get("/users/:username", UserHandler.GetUser)       // 获取用户信息
		apiV1.Put("/users/:username", UserHandler.UpdateUser)    // 更新用户信息
		apiV1.Delete("/users/:username", UserHandler.DeleteUser) // 删除用户
		apiV1.Get("/users/items", UserHandler.GetUsers)          // 获取用户列表

		apiV1.Put("/hub/:hid/members/:username", UserHandler.Join)     // 加入群组
		apiV1.Delete("/hub/:hid/members/:username", UserHandler.Leave) // 离开群组
		apiV1.Get("/hub/:hid/members", UserHandler.Joined)             // 查询群组成员

		apiV1.Put("/user/following/:username", UserHandler.Follow)      // 关注
		apiV1.Delete("/user/following/:username", UserHandler.Unfollow) // 取消关注
		apiV1.Get("/user/following", UserHandler.Following)             // 查询关注

		apiV1.Get("/chat/auth", ChatHandler.ChatAuth)               // 聊天验证
		apiV1.Get("/chat/:topic/messages", ChatHandler.GetMessages) // 获取聊天记录
	}

	app.Get("/ws", middleware.WebSocket, websocket.New(ChatHandler.ChatHandler)) // 建立websocket连接

	return app
}
