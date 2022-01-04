// @description: 数据库表通用字段
// @file: base.go
// @date: 2021/11/19

package data

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
