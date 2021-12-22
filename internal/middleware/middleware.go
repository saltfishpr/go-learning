// @file: middleware.go
// @date: 2021/11/22

// Package middleware 提供中间件
package middleware

import (
	"learning/config"
	"learning/internal/common/rediscache"
	"learning/internal/constant/e"
	"learning/internal/utils"
	"learning/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/websocket/v2"
)

var Recover = recover.New(
	recover.Config{
		EnableStackTrace: true,
	},
)

var Logger = fiberlogger.New()

var JwtAuth = jwtware.New(
	jwtware.Config{
		AuthScheme: config.AuthScheme,
		ContextKey: config.ContextKey,
		SigningKey: []byte(config.SigningKey),
	},
)

var CORS = cors.New(
	cors.Config{
		AllowOrigins: "http://localhost:9090",
	},
)

var Cache = cache.New()

var WebSocket = func(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		claims, err := utils.VerifyToken(c.Query("token"))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(e.Failed(e.Unauthorized))
		}
		var account string
		if err := rediscache.Get(
			config.DisposableTokenPrefix+claims["jti"].(string), &account,
		); err != nil || claims["account"].(string) != account {
			return c.Status(fiber.StatusUnauthorized).JSON(e.Failed(e.Unauthorized))
		}
		if err := rediscache.Del(config.DisposableTokenPrefix + claims["jti"].(string)); err != nil {
			logger.Error("delete disposable token error: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.Error))
		}
		c.Locals("account", account)
		return c.Next()
	}
	return c.SendStatus(fiber.StatusUpgradeRequired)
}
