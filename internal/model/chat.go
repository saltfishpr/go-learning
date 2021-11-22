// @description: 聊天信息数据结构
// @file: chat.go
// @date: 2021/11/22

package model

type RecvMessage struct {
	Mode    *ChatMode `json:"mode" validate:"required"`    // 对象类型 0:user 1:hub
	ID      *string   `json:"id" validate:"required"`      // 对象ID
	Content *string   `json:"content" validate:"required"` // 消息内容
}

type SendMessage struct {
	Content string
}
