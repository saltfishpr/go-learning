// @file: middleware.go
// @date: 2021/11/22

// Package middleware 提供中间件.
package middleware

import (
	"learning/internal/common/rediscache"
	"learning/internal/constant"
	"learning/internal/constant/e"
	"learning/internal/log"
	"learning/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"
)

var Recover = recover.New(
	recover.Config{
		EnableStackTrace: true,
	},
)

func Logger(l *log.Logger) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		l.Info("request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("ip", c.IP()),
			zap.String("user-agent", c.Get("User-Agent")),
		)
		return c.Next()
	}
}

var Pprof = pprof.New()

var JwtAuth = jwtware.New(
	jwtware.Config{
		AuthScheme: constant.AuthScheme,
		ContextKey: constant.ContextKey,
		SigningKey: []byte(constant.SigningKey),
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
		var username string
		if err := rediscache.Get(
			constant.DisposableTokenPrefix+claims["jti"].(string), &username,
		); err != nil || claims["username"].(string) != username {
			return c.Status(fiber.StatusUnauthorized).JSON(e.Failed(e.Unauthorized))
		}
		if err := rediscache.Del(constant.DisposableTokenPrefix + claims["jti"].(string)); err != nil {
			zap.S().Error("delete disposable token error: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.Error))
		}
		c.Locals("username", username)
		return c.Next()
	}
	return c.SendStatus(fiber.StatusUpgradeRequired)
}
