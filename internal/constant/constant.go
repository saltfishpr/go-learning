// @description: 常量配置。
// @file: constant.go
// @date: 2021/11/2

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

	RefreshTokenPrefix    = "refresh_token_jti_"
	DisposableTokenPrefix = "disposable_token_jti_"
)

const (
	ShutdownTime = 5 * time.Second
)

const (
	TopicPrefixLen = 3
)
