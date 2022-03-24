// @description: 定义路由
// @file: app.go
// @date: 2021/11/2

package internal

import (
	v1 "learning/internal/api/v1"
	"learning/internal/data"
	"learning/internal/log"
	"learning/internal/middleware"
	"learning/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func NewApp(logger *log.Logger) *fiber.App {
	app := fiber.New(
		fiber.Config{
			JSONEncoder: utils.JsonMarshal,
			JSONDecoder: utils.JsonUnmarshal,
		},
	)

	conn, err := data.NewPostgres()
	if err != nil {
		logger.Fatal("connect to database error: ", err)
	}

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("logger", logger)
		c.Locals("dbconn", conn)
		return c.Next()
	})

	app.Use(middleware.Recover, middleware.Pprof, middleware.Logger(logger))
	app.Use(middleware.CORS)

	app.Post("/auth/login", v1.Login)          // 登录/注册
	app.Get("/auth/token/refresh", v1.Refresh) // 刷新token

	apiV1 := app.Group("/api/v1")
	apiV1.Use(middleware.JwtAuth)
	apiV1.Use(middleware.Cache)
	{
		apiV1.Post("/hubs", v1.CreateHub)        // 创建群组
		apiV1.Get("/hubs/:hid", v1.GetHub)       // 获取群组信息
		apiV1.Put("/hubs/:hid", v1.UpdateHub)    // 更新群组信息
		apiV1.Delete("/hubs/:hid", v1.DeleteHub) // 删除群组
		apiV1.Get("/hubs/items", v1.GetHubs)     // 获取群组列表

		apiV1.Post("/users", v1.CreateUser)            // 创建用户
		apiV1.Get("/users/:account", v1.GetUser)       // 获取用户信息
		apiV1.Put("/users/:account", v1.UpdateUser)    // 更新用户信息
		apiV1.Delete("/users/:account", v1.DeleteUser) // 删除用户
		apiV1.Get("/users/items", v1.GetUsers)         // 获取用户列表

		apiV1.Put("/hub/:hid/members/:account", v1.Join)     // 加入群组
		apiV1.Delete("/hub/:hid/members/:account", v1.Leave) // 离开群组
		apiV1.Get("/hub/:hid/members", v1.Joined)            // 查询群组成员

		apiV1.Put("/user/following/:account", v1.Follow)      // 关注
		apiV1.Delete("/user/following/:account", v1.Unfollow) // 取消关注
		apiV1.Get("/user/following", v1.Following)            // 查询关注

		apiV1.Get("/chat/auth", v1.ChatAuth)               // 聊天验证
		apiV1.Get("/chat/:topic/messages", v1.GetMessages) // 获取聊天记录
	}

	app.Get("/ws", middleware.WebSocket, websocket.New(v1.ChatHandler)) // 建立websocket连接

	return app
}
