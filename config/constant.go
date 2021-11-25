// @description: 常量配置。
// @file: constant.go
// @date: 2021/11/2

package config

import "time"

// BufferedChan chan size
const BufferedChan = 1

// JWT config
const (
	TokenExpireTime = 15 * time.Minute
	SigningKey      = "saltfish"
	ContextKey      = "user"
	AuthScheme      = "Bearer"
)
