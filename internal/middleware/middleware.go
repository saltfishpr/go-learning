// @file: middleware.go
// @date: 2021/11/22

// Package middleware 提供中间件
package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"

	"learning/config"
	"learning/logger"
)

var Recover = recover.New(recover.Config{
	EnableStackTrace: true,
})

var Timer = func(c *fiber.Ctx) error {
	begin := time.Now().UnixNano()
	defer func() {
		end := time.Now().UnixNano()
		logger.Infof("url: %s%s%s use: %s%dns%s", config.BlueBold, c.Path(), config.Reset, config.Green, end-begin, config.Reset)
	}()
	return c.Next()
}

var JwtAuth = jwtware.New(jwtware.Config{
	AuthScheme: config.AuthScheme,
	ContextKey: config.ContextKey,
	SigningKey: []byte(config.SigningKey),
})
