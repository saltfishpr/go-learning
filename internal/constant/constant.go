// @file: constant.go
// @date: 2021/11/2

// Package constant 常量配置.
package constant

import "time"

// BufferedChan chan size
const BufferedChan = 1

// JWT config
const (
	SigningKey = "saltfish"
	ContextKey = "user"
	AuthScheme = "Bearer"

	TokenExpireTime        = 15 * time.Minute
	RefreshTokenExpireTime = 24 * time.Hour

	ShutdownTimeout = 15 * time.Second

	RefreshTokenPrefix    = "refresh_token_jti_"
	DisposableTokenPrefix = "disposable_token_jti_"
)

const (
	ReadTimeout = 15 * time.Second
)

const (
	TopicPrefixLen = 3
)
