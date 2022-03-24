// @description: 数据库表数据结构定义
// @file: model.go
// @date: 2022/3/21

package data

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

type User struct {
	Base
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:user_udx_delete"`

	Username *string `gorm:"size:32;not null;uniqueIndex:user_udx_delete"`
	Password *string `gorm:"size:32;not null"`
	Phone    *string `gorm:"size:32"`
	Email    *string `gorm:"size:32"`
	Nickname *string `gorm:"size:32;not null"`
	Address  *string `gorm:"size:128"`

	Hubs    []*Hub  `gorm:"many2many:user_hub"`
	Friends []*User `gorm:"many2many:user_friends"`
}

type Hub struct {
	Base
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:hub_udx_delete"`

	HID  *string `gorm:"size:8;not null;uniqueIndex:hub_udx_delete"`
	Name *string `gorm:"size:32;not null"`
	Size *int    `gorm:"not null"`

	Users []*User `gorm:"many2many:user_hub"`
}

type Message struct {
	Base

	From    string `gorm:"size:32;not null"`
	Topic   string `gorm:"size:36;not null"`
	Payload []byte `gorm:"size:5120;not null"`
	Mode    int    `gorm:"not null"`
}
