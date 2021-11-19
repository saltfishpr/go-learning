// @description: 定义通用字段
// @file: base.go
// @date: 2021/11/19

package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Base struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_delete"`
}
