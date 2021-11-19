// @description:
// @file: data.go
// @date: 2021/11/18

// Package data 提供数据访问
package data

import (
	"sync"

	"learning/logger"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	once sync.Once
	db   *gorm.DB
	err  error
)

func NewSqlite() *gorm.DB {
	once.Do(func() {
		db, err = gorm.Open(sqlite.Open("output/gorm.db"), &gorm.Config{
			Logger: gormLogger.Discard,
		})
		if err != nil {
			logger.Fatal("connect sqlite database error: ", err)
		}
	})
	return db
}

func NewPostgres() *gorm.DB {
	once.Do(func() {
		dsn := "host=localhost user=guest password=123456 dbname=chat port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: gormLogger.Discard,
		})
		if err != nil {
			logger.Fatal("connect postgras database error: ", err)
		}
	})
	return db
}
