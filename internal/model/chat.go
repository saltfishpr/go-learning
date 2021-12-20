// @description: 聊天信息数据结构
// @file: chat.go
// @date: 2021/11/22

package model

import "fmt"

type Message struct {
	From    string   `json:"from" validate:"required"`
	To      string   `json:"to" validate:"required"`
	Payload []byte   `json:"payload" validate:"required"` // 消息内容
	Mode    ChatMode `json:"mode" validate:"required"`    // 对象类型 0:user 1:hub
}

func (m Message) String() string {
	return fmt.Sprintf("[Mode: %d] %s send to %s, content: %s", m.Mode, m.From, m.To, string(m.Payload))
}
