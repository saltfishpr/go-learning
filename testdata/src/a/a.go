package a

import (
	"gorm.io/gorm"
)

type User struct{}

var users []User

var DB *gorm.DB

// Good case
func f1() {
	db := DB

	db.Where("name = ?", "jinzhu").Where("age = ?", 18).Find(&users)
	// 第一个 `Where("name = ?", "jinzhu")` 启动了一个 `*gorm.DB` 实例或 `*gorm.Statement` 的链式调用。
	// 第二个 `Where("age = ?", 18)` 向现有的 `*gorm.Statement` 添加了一个新的条件。
	// `Find(&users)` 是一个终结方法，执行已注册的查询回调，生成并运行以下 SQL：
	// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18;

	db.Where("name = ?", "jinzhu2").Where("age = ?", 20).Find(&users)
	// 在这里，`Where("name = ?", "jinzhu2")` 启动了一个新的链式调用，创建了一个全新的 `*gorm.Statement`。
	// `Where("age = ?", 20)` 向这个新的 `*gorm.Statement` 添加条件。
	// `Find(&users)` 再次终结查询，生成并运行以下 SQL：
	// SELECT * FROM users WHERE name = 'jinzhu2' AND age = 20;

	db.Find(&users)
	// 直接调用 `Find(&users)` 而不使用任何 `Where` 条件，会启动一个新的链式调用并执行以下 SQL：
	// SELECT * FROM users;
}

// Bad case
func f2() {
	db := DB // 'db' is a newly initialized *gorm.DB, safe to reuse.

	tx := db.Where("name = ?", "jinzhu")
	// `Where("name = ?", "jinzhu")` initializes a `*gorm.Statement` instance, which should not be reused across different logical operations.

	tx.Where("age = ?", 18).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18

	// WARNING
	// 错误地重用了 'tx'，导致条件叠加，查询被污染
	tx.Where("age = ?", 28).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18 AND age = 28;
}

// Good case
func f3() {
	db := DB // 'db' is a newly initialized *gorm.DB, safe to reuse.

	// `Session`、`WithContext`、`Debug`` 方法返回一个标记为可安全重用的 *gorm.DB 实例。它们基于当前条件初始化一个新的 *gorm.Statement。
	tx := db.Where("name = ?", "jinzhu").Session(&gorm.Session{})
	// tx := db.Where("name = ?", "jinzhu").WithContext(context.Background())
	// tx := db.Where("name = ?", "jinzhu").Debug()

	tx.Where("age = ?", 18).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18

	tx.Where("age = ?", 28).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age = 28;
}

func setupQuery(db *gorm.DB) *gorm.DB {
	return db.Where("name = ?", "jinzhu")
}

// Bad case
func f4(db *gorm.DB) {
	tx := setupQuery(db)

	tx.Where("age = ?", 18).Find(&users)

	tx.Where("age = ?", 28).Find(&users)
}
