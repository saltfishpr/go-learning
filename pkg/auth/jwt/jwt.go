package jwt

import (
	"context"
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
