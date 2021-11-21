// @description: 用户数据结构
// @file: user.go
// @date: 2021/11/18

package model

type User struct {
	Account  *string `json:"account" validate:"required"`
	Password *string `json:"password" validate:"required"`
	Nickname *string `json:"nickname" validate:"required"`
	Address  *string `json:"address,omitempty"`
}
