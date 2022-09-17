// @file: e.go
// @date: 2021/11/25

// Package e 定义错误码与错误信息.
package e

type ErrorCode int

const (
	Error = 10000 + iota
	Unauthorized
	InvalidParams

	ExistHub
	HubNotFound

	ExistUsername
	LoginFailed
	TokenParseFailed
)

var errorMessage = map[ErrorCode]string{
	Error:            "fail",
	Unauthorized:     "unauthorized or token expired",
	InvalidParams:    "invalid params",
	ExistHub:         "hub already exists",
	HubNotFound:      "no such hub",
	ExistUsername:    "username already exists",
	LoginFailed:      "wrong username or password",
	TokenParseFailed: "parse token error",
}

type ErrorResult struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func Failed(code ErrorCode, opts ...Option) *ErrorResult {
	result := &ErrorResult{Code: code, Message: errorMessage[code]}
	for _, opt := range opts {
		opt(result)
	}
	return result
}

type Option func(*ErrorResult)

func WithMessage(message string) Option {
	return func(result *ErrorResult) {
		result.Message = message
	}
}
