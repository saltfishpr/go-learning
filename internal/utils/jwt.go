// @description: jwt token util
// @file: jwt.go
// @date: 2021/11/20

package utils

import (
	"errors"
	"fmt"
	"time"

	"learning/internal/common/rediscache"
	"learning/internal/constant"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

func GetUsernameFromCtx(c *fiber.Ctx) (string, bool) {
	t := c.Locals(constant.ContextKey)
	if t == nil {
		return "", false
	}
	user := t.(*jwt.Token)
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return "", ok
	}
	username, ok := claims["username"].(string)
	return username, ok
}

func MustGetUsernameFromCtx(c *fiber.Ctx) string {
	user := c.Locals(constant.ContextKey).(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return username
}

func GenerateToken(username string) (string, error) {
	claims := &CustomClaims{Username: username}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(constant.TokenExpireTime))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SigningKey))
}

func GenerateRefreshToken(username string) (string, error) {
	id := NewNode().Generate().String()
	err := rediscache.Set(constant.RefreshTokenPrefix+id, username, constant.RefreshTokenExpireTime)
	if err != nil {
		return "", err
	}
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(constant.RefreshTokenExpireTime)),
		ID:        id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SigningKey))
}

func GenerateTokenPair(username string) (fiber.Map, error) {
	t, err := GenerateToken(username)
	if err != nil {
		zap.S().Error("generate token error: ", err)
		return nil, errors.New("generate token error")
	}
	rt, err := GenerateRefreshToken(username)
	if err != nil {
		zap.S().Error("generate refresh token error: ", err)
		return nil, errors.New("generate refresh token error")
	}
	return fiber.Map{"token": t, "refresh_token": rt}, nil
}

func GenerateDisposableToken(username string) (string, error) {
	claims := &CustomClaims{Username: username}
	id := NewNode().Generate().String()
	err := rediscache.Set(constant.DisposableTokenPrefix+id, username, constant.TokenExpireTime)
	if err != nil {
		return "", err
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(constant.TokenExpireTime))
	claims.ID = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SigningKey))
}

func VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constant.SigningKey), nil
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
