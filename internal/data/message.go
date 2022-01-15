// @description: 聊天记录表字段
// @file: message.go
// @date: 2022/01/05

// Package data
package data

type Message struct {
	Base

	From    string `gorm:"size:32;not null"`
	Topic   string `gorm:"size:36;not null"`
	Payload []byte `gorm:"size:5120;not null"`
	Mode    int    `gorm:"not null"`
}

func CreateMessage(message *Message) error {
	res := db.Create(message)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetMessagesPagination(topic string, offset int, limit int) (int, []*Message, error) {
	var messages []*Message

	res := db.Offset(offset).Limit(limit).Where("topic = ?", topic).Find(&messages)
	if res.Error != nil {
		return 0, nil, res.Error
	}
	var count int64
	db.Model(&Message{}).Where("topic = ?", topic).Count(&count)
	return int(count), messages, nil
}
