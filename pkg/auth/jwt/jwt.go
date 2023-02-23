package jwt

import (
	"context"

	jwtV4 "github.com/golang-jwt/jwt/v4"
)

// Auth 自定义认证信息.
type Auth struct {
	Token string
}

// GetRequestMetadata 获取认证信息.
func (c Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": c.Token,
	}, nil
}

// RequireTransportSecurity 是否需要安全传输.
func (c Auth) RequireTransportSecurity() bool {
	return false
}

var contextKeyClaims int

func NewContext(ctx context.Context, claims *jwtV4.MapClaims) context.Context {
	return context.WithValue(ctx, &contextKeyClaims, claims)
}

func FromContext(ctx context.Context) (*jwtV4.MapClaims, bool) {
	claims, ok := ctx.Value(&contextKeyClaims).(*jwtV4.MapClaims)
	return claims, ok
}
