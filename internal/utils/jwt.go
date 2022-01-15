// @description: 从上下文获取jwt信息
// @file: jwt.go
// @date: 2021/11/20

package utils

import (
	"errors"
	"fmt"
	"time"

	"learning/config"
	"learning/internal/common/rediscache"
	"learning/internal/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Account string `json:"account"`
}

func GetUserAccountFromCtx(c *fiber.Ctx) (string, bool) {
	t := c.Locals(config.ContextKey)
	if t == nil {
		return "", false
	}
	user := t.(*jwt.Token)
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return "", ok
	}
	account, ok := claims["account"].(string)
	return account, ok
}

func MustGetUserAccountFromCtx(c *fiber.Ctx) string {
	user := c.Locals(config.ContextKey).(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	account := claims["account"].(string)
	return account
}

func GenerateToken(account string) (string, error) {
	claims := &CustomClaims{Account: account}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(config.TokenExpireTime))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SigningKey))
}

func GenerateRefreshToken(account string) (string, error) {
	id := NewNode().Generate().String()
	err := rediscache.Set(config.RefreshTokenPrefix+id, account, config.RefreshTokenExpireTime)
	if err != nil {
		return "", err
	}
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.RefreshTokenExpireTime)),
		ID:        id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SigningKey))
}

func GenerateTokenPair(account string) (fiber.Map, error) {
	t, err := GenerateToken(account)
	if err != nil {
		logger.Error("generate token error: ", err)
		return nil, errors.New("generate token error")
	}
	rt, err := GenerateRefreshToken(account)
	if err != nil {
		logger.Error("generate refresh token error: ", err)
		return nil, errors.New("generate refresh token error")
	}
	return fiber.Map{"token": t, "refresh_token": rt}, nil
}

func GenerateDisposableToken(account string) (string, error) {
	claims := &CustomClaims{Account: account}
	id := NewNode().Generate().String()
	err := rediscache.Set(config.DisposableTokenPrefix+id, account, config.TokenExpireTime)
	if err != nil {
		return "", err
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(config.TokenExpireTime))
	claims.ID = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SigningKey))
}

func VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("can not convert claims to jwt.MapClaims")
	}
	if !token.Valid {
		return nil, errors.New("token not valid")
	}
	return claims, nil
}
