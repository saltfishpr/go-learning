// @description: 用户数据结构
// @file: user.go
// @date: 2021/11/18

package model

// User .
type User struct {
	Username *string `json:"username"           validate:"required"`
	Password *string `json:"password"           validate:"required"`
	Phone    *string `json:"phone,omitempty"`
	Email    *string `json:"email,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	Address  *string `json:"address,omitempty"`
}
