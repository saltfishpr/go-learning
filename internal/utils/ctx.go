// @description: get value from context
// @file: ctx.go
// @date: 2022/3/25

package utils

import (
	"learning/internal/data"
	"learning/internal/log"

	"github.com/gofiber/fiber/v2"
)

func MustGetConnectionFromContext(ctx *fiber.Ctx) data.Connection {
	conn, ok := ctx.Locals("dbconn").(data.Connection)
	if !ok {
		panic("cannot get connection from context")
	}
	return conn
}

func MustGetLoggerFromContext(ctx *fiber.Ctx) *log.Logger {
	logger, ok := ctx.Locals("logger").(*log.Logger)
	if !ok {
		panic("cannot get logger from context")
	}
	return logger
}
