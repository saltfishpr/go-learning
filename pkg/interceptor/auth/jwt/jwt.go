package jwt

import (
	"context"
	"errors"
	"strings"

	jwtV4 "github.com/golang-jwt/jwt/v4"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/metadata"
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

func AuthFunc(verify func(string) (*jwtV4.MapClaims, error)) grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		tokenStr, err := getTokenFromContext(ctx)
		if err != nil {
			return nil, err
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		claims, err := verify(tokenStr)
		if err != nil {
			return nil, err
		}

		return NewContext(ctx, claims), nil
	}
}

func getTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata in context")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return "", errors.New("no authorization")
	}

	return values[0], nil
}

var contextKeyClaims int

func NewContext(ctx context.Context, claims *jwtV4.MapClaims) context.Context {
	return context.WithValue(ctx, &contextKeyClaims, claims)
}

func FromContext(ctx context.Context) (*jwtV4.MapClaims, bool) {
	claims, ok := ctx.Value(&contextKeyClaims).(*jwtV4.MapClaims)
	return claims, ok
}
