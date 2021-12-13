// @file: middleware.go
// @date: 2021/11/22

// Package middleware 提供中间件
package middleware

import (
	"fmt"

	"learning/config"
	"learning/internal/common/connstorage"
	"learning/internal/common/gocache"
	"learning/internal/constant/e"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v4"
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
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(e.Failed(e.Unauthorized))
		},
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
		sid := c.Params("sid")
		account, ok := gocache.Get(sid)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(e.Failed(e.InvalidSID))
		}

		token, err := jwt.Parse(
			c.Query("token"), func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(config.SigningKey), nil
			},
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				e.Failed(
					e.TokenParse, e.WithMessage("invalid token"),
				),
			)
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && claims["account"] == account {
			if _, ok := connstorage.Get(sid); ok {
				return c.Status(fiber.StatusInternalServerError).JSON(e.Failed(e.SIDInUse))
			}
			c.Locals("sid", sid)
			c.Locals("account", account)
			return c.Next()
		}
		return c.Status(fiber.StatusUnauthorized).JSON(e.Failed(e.Unauthorized))
	}

	return c.SendStatus(fiber.StatusUpgradeRequired)
}
