// @description: 从上下文获取jwt信息
// @file: jwt.go
// @date: 2021/11/20

package utils

import (
	"learning/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserAccountFromCtx(c *fiber.Ctx) (string, bool) {
	t := c.Locals(config.ContextKey)
	if t == nil {
		return "", false
	}
	user := t.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	account := claims["account"].(string)
	return account, true
}

func MustGetUserAccountFromCtx(c *fiber.Ctx) string {
	user := c.Locals(config.ContextKey).(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	account := claims["account"].(string)
	return account
}
