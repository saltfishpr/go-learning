// @description: 从上下文获取jwt信息
// @file: jwt.go
// @date: 2021/11/20

package utils

import (
	"learning/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserAccountFromCtx(c *fiber.Ctx) string {
	user := c.Locals(config.ContextKey).(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	account := claims["account"].(string)
	return account
}

func GetUserAccountWebsocketConn(c *websocket.Conn) string {
	user := c.Locals(config.ContextKey).(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	account := claims["account"].(string)
	return account
}
