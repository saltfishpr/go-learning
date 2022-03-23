// @description: 聊天室数据结构
// @file: model.go
// @date: 2021/11/1

package model

// Hub is a chat room that users can join and leave.
type Hub struct {
	HID  *string `json:"hid"  validate:"required"`
	Name *string `json:"name" validate:"required"`
	Size *int    `json:"size" validate:"required"`
}
