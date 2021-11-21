// @description: 聊天信息数据结构
// @file: chat.go
// @date: 2021/11/22

package model

type Message struct {
	Mode    *ChatMode `json:"mode" validate:"required"`    // 对象 0:user 1:hub
	Content *string   `json:"content" validate:"required"` // 消息内容
}
