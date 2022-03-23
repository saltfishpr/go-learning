// @description: 聊天信息数据结构
// @file: message.go
// @date: 2021/11/22

package model

import "fmt"

type RecvMessage struct {
	From    string      `json:"from"    validate:"required"` // 发送消息的用户
	Topic   string      `json:"topic"   validate:"required"` // 发送给谁 前缀:ID 的格式
	Payload []byte      `json:"payload" validate:"required"` // 数据
	Mode    PayloadMode `json:"mode"    validate:"required"` // 消息类型
}

func (m RecvMessage) String() string {
	return fmt.Sprintf("%s send to %s, content: %s", m.From, m.Topic, string(m.Payload))
}

type SendMessage struct {
	From    string      `json:"from"`
	Payload []byte      `json:"payload"`
	Mode    PayloadMode `json:"mode"`
}

type MessagesPaginationRequest struct {
	Topic  string
	Offset int
	Limit  int
}

type MessagesPaginationResponse struct {
	Data   []SendMessage
	Offset int
	Limit  int
	Count  int
}
