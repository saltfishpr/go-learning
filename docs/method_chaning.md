GORM 的方法链功能可实现平滑流畅的编码风格。 Here's an example:

```go
db.Where("name = ?", "jinzhu").Where("age = ?", 18).First(&user)
```

## Method Categories

GORM 将方法分为三大类： `Chain Methods`, `Finisher Methods`, and `New Session Methods`.

### Chain Methods

用于修改或追加目前 `Clauses` 的 `Statement`

```go
var chainingMethods = map[string]bool{
	"Create":          true,
	"CreateInBatches": true,
	"Save":            true,
	"First":           true,
	"Take":            true,
	"Last":            true,
	"Find":            true,
	"FindInBatches":   true,
	"FirstOrInit":     true,
	"FirstOrCreate":   true,
	"Update":          true,
	"Updates":         true,
	"UpdateColumn":    true,
	"UpdateColumns":   true,
	"Delete":          true,
	"Count":           true,
	"Row":             true,
	"Rows":            true,
	"Scan":            true,
	"Pluck":           true,
	"ScanRows":        true,
	"Connection":      true,
	"Transaction":     true,
	"Begin":           true,
	"Commit":          true,
	"Rollback":        true,
	"SavePoint":       true,
	"RollbackTo":      true,
	"Exec":            true,
}
```

For a comprehensive list, visit [GORM Chainable API](https://github.com/go-gorm/gorm/blob/master/chainable_api.go). Also, the [SQL Builder](sql_builder.html) documentation offers more details about `Clauses`.

### Finisher Methods

终结方法是即时的，执行生成和运行 SQL 命令的注册回调

```go
var finisherMethods = map[string]bool{
	"Model":      true,
	"Clauses":    true,
	"Table":      true,
	"Distinct":   true,
	"Select":     true,
	"Omit":       true,
	"MapColumns": true,
	"Where":      true,
	"Not":        true,
	"Or":         true,
	"Joins":      true,
	"InnerJoins": true,
	"joins":      true,
	"Group":      true,
	"Having":     true,
	"Order":      true,
	"Limit":      true,
	"Offset":     true,
	"Scopes":     true,
	"Preload":    true,
	"Attrs":      true,
	"Assign":     true,
	"Unscoped":   true,
	"Raw":        true,
}
```

### New Session Methods

```go
var newSessionMethods = map[string]bool{
	"Session":     true,
	"WithContext": true,
	"Debug":       true,
}
```

## Examples for Clarity

- **Example 1: Safe Instance Reuse**

```go
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 'db' is a newly initialized `*gorm.DB`, which is safe to reuse.

db.where ("name = ?", "jinzhu").Where("age = ?", 18). ind(&用户)
// 第一个`Where ("name = ?", "jinzhu")`是一个启动一个 `*gorm.DB` 实例或`*gorm.Statement`的链式方法。
// The second `Where("age = ?", 18)` call adds a new condition to the existing `*gorm.Statement`.
// `Find(&users)` is a finisher method, executing registered Query Callbacks, generating and running:
// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18;

db.Where("name = ?", "jinzhu2").Where("age = ?", 20).Find(&users)
// Here, `Where("name = ?", "jinzhu2")` starts a new chain, creating a fresh `*gorm.Statement`.
// `Where("age = ?", 20)` adds to this new statement.
// `Find(&users)` again finalizes the query, executing and generating:
// SELECT * FROM users WHERE name = 'jinzhu2' AND age = 20;

db.Find(&users)
// Directly calling `Find(&users)` without any `Where` starts a new chain and executes:
// SELECT * FROM users;
```

在这个例子中，每个方法调用链都是独立的，确保干净、无污染的 SQL 查询。

- **(Bad) 示例2：不安全的实例重用**

```go
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 'db' is a newly initialized *gorm.DB, safe for initial reuse.

tx := db.Where("name = ?", "jinzhu")
// `Where("name = ?", "jinzhu")` initializes a `*gorm.Statement` instance, which should not be reused across different logical operations.

// Good case
tx.Where("age = ?", 18).Find(&users)
// Reuses 'tx' correctly for a single logical operation, executing:
// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18

// Bad case
tx.Where("age = ?", 28).Find(&users)
// Incorrectly reuses 'tx', compounding conditions and leading to a polluted query:
// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18 AND age = 28;
```

In this bad example, reusing the `tx` variable leads to compounded conditions, which is generally not desirable.

- **例3：使用新会话方法安全重新使用**

```go
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 'db' is a newly initialized *gorm.DB, safe to reuse.

tx := db.Where("name = ?", "jinzhu").Session(&gorm.Session{})
tx := db.Where("name = ?", "jinzhu").WithContext(context.Background())
tx := db.Where("name = ?", "jinzhu").Debug()
// `Session`, `WithContext`, `Debug` methods return a `*gorm.DB` instance marked as safe for reuse. They base a newly initialized `*gorm.Statement` on the current conditions.

// Good case
tx.Where("age = ?", 18).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18

// Good case
tx.Where("age = ?", 28).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' AND age = 28;
```

在这个例子中，为每个逻辑操作正确使用新建会话方法 `Session`, `WithContext`, `Debug` 初始化一个 `*gorm.DB`实例，从而防止了条件污染，确保每个查询都是独立的，并且基于所提供的特定条件。

Overall, these examples illustrate the importance of understanding GORM's behavior with respect to method chaining and instance management to ensure accurate and efficient database querying.