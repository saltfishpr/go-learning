## Ctx.JSON()

**Fiber version**

github.com/gofiber/fiber/v2 v2.22.0

**Issue description**

c.JSON() panic but encoding/json.Marshal works well.

**Code snippet**

model.User:

```go
type User struct {
Account  *string `json:"account" validate:"required"`
Password *string `json:"password" validate:"required"`
Nickname *string `json:"nickname" validate:"required"`
Address  *string `json:"address,omitempty"`

Hubs    []*Hub  `json:"hubs,omitempty"`
Friends []*User `json:"friends,omitempty"`
}
```

service.GetUserInfo:

```go
func GetUserInfo(c *fiber.Ctx) error {
account := c.Params("account")
if len(account) == 0 {
return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "缺少用户ID"})
}

user, err := service.GetUserByAccount(account)
if err != nil {
logger.Error("get user error: ", err)
return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "没有此用户"})
}

return c.Status(fiber.StatusOK).JSON(user) // add breakpoint here
}
```

`user` data at breakpoint:

```
*learning/internal/model.User {Account: *"saltfish", Password: *"123456", Nickname: *"咸鱼", Address: *string nil, Hubs: []*learning/internal/model.Hub len: 1, cap: 1, [*(*"learning/internal/model.Hub")(0x4000542348)], Friends: []*learning/internal/model.User len: 1, cap: 1, [*(*"learning/internal/model.User")(0x400008e460)]}
```

user.Hubs:

```
[]*learning/internal/model.Hub len: 1, cap: 1, [*{HID: *"alpha", Name: *"聊天室1", Size: *100}]
```

user.Friends:

```
[]*learning/internal/model.User len: 1, cap: 1, [*{Account: *"saltfishpr", Password: *"123456", Nickname: *"咸鱼硕2", Address: *string nil, Hubs: []*learning/internal/model.Hub len: 0, cap: 0, nil, Friends: []*learning/internal/model.User len: 0, cap: 0, nil}]
```

json encode works well until marshal user.Friends

```
"{\"account\":\"saltfish\",\"password\":\"123456\",\"nickname\":\"咸鱼\",\"hubs\":[{\"hid\":\"alpha\",\"name\":\"聊天室1\",\"size\":100}],\"friends\":[{\"account\":true,\"password\":"
```

then panic happen

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0xa pc=0x8038f0]

goroutine 24 [running]:
github.com/gofiber/fiber/v2/internal/go-json/encoder/vm.ptrToString(...)
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/internal/go-json/encoder/vm/util.go:77
github.com/gofiber/fiber/v2/internal/go-json/encoder/vm.Run(0x40005160d0, {0x400049a400, 0x0, 0x400}, 0x400007ac40)
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/internal/go-json/encoder/vm/vm.go:3859 +0x26d40
github.com/gofiber/fiber/v2/internal/go-json.encodeRunCode(0x40005160d0, {0x400049a400, 0x0, 0x400}, 0x400007ac40)
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/internal/go-json/encode.go:307 +0x170
github.com/gofiber/fiber/v2/internal/go-json.encode(0x40005160d0, {0xc44ae0, 0x40000b41e0})
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/internal/go-json/encode.go:232 +0x284
github.com/gofiber/fiber/v2/internal/go-json.marshal({0xc44ae0, 0x40000b41e0}, {0x0, 0x0, 0x0})
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/internal/go-json/encode.go:147 +0xbc
github.com/gofiber/fiber/v2/internal/go-json.MarshalWithOption(...)
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/internal/go-json/json.go:186
github.com/gofiber/fiber/v2/internal/go-json.Marshal({0xc44ae0, 0x40000b41e0})
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/internal/go-json/json.go:171 +0x38
github.com/gofiber/fiber/v2.(*Ctx).JSON(0x40001582c0, {0xc44ae0, 0x40000b41e0})
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/ctx.go:650 +0x40
learning/internal/api/v1.GetUserInfo(0x40001582c0)
        /home/saltfish/Programming/Go/github.com/SaltFishPr/go-learning/internal/api/v1/users.go:72 +0x290
github.com/gofiber/fiber/v2.(*App).next(0x4000397520, 0x40001582c0)
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/router.go:127 +0x1c0
github.com/gofiber/fiber/v2.(*Ctx).Next(0x40001582c0)
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/ctx.go:747 +0x70
github.com/gofiber/jwt/v3.makeCfg.func1(0x40001582c0)
        /home/saltfish/go/pkg/mod/github.com/gofiber/jwt/v3@v3.2.1/config.go:124 +0x28
github.com/gofiber/jwt/v3.New.func1(0x40001582c0)
        /home/saltfish/go/pkg/mod/github.com/gofiber/jwt/v3@v3.2.1/main.go:60 +0x378
github.com/gofiber/fiber/v2.(*App).next(0x4000397520, 0x40001582c0)
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/router.go:127 +0x1c0
github.com/gofiber/fiber/v2.(*Ctx).Next(0x40001582c0)
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/ctx.go:747 +0x70
learning/internal/middleware.glob..func1(0x40001582c0)
        /home/saltfish/Programming/Go/github.com/SaltFishPr/go-learning/internal/middleware/middleware.go:28 +0xc0
github.com/gofiber/fiber/v2.(*App).next(0x4000397520, 0x40001582c0)
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/router.go:127 +0x1c0
github.com/gofiber/fiber/v2.(*App).handler(0x4000397520, 0x40000de000)
        /home/saltfish/go/pkg/mod/github.com/gofiber/fiber/v2@v2.22.0/router.go:155 +0x104
github.com/valyala/fasthttp.(*Server).serveConn(0x400017f400, {0xe60ba0, 0x4000138018})
        /home/saltfish/go/pkg/mod/github.com/valyala/fasthttp@v1.31.0/server.go:2278 +0x106c
github.com/valyala/fasthttp.(*workerPool).workerFunc(0x40000dc000, 0x4000464000)
        /home/saltfish/go/pkg/mod/github.com/valyala/fasthttp@v1.31.0/workerpool.go:223 +0x7c
github.com/valyala/fasthttp.(*workerPool).getCh.func1(0x40000dc000, 0x4000464000, {0xc41960, 0x4000464000})
        /home/saltfish/go/pkg/mod/github.com/valyala/fasthttp@v1.31.0/workerpool.go:195 +0x30
created by github.com/valyala/fasthttp.(*workerPool).getCh
        /home/saltfish/go/pkg/mod/github.com/valyala/fasthttp@v1.31.0/workerpool.go:194 +0x1f0
```

A simple test file.

```go
package main

import (
    "encoding/json"
    "log"

    "github.com/gofiber/fiber/v2"
)

type User struct {
    Account  *string `json:"account" validate:"required"`
    Password *string `json:"password" validate:"required"`
    Nickname *string `json:"nickname" validate:"required"`
    Address  *string `json:"address,omitempty"`

    Hubs    []*Hub  `json:"hubs,omitempty"`
    Friends []*User `json:"friends,omitempty"`
}

type Hub struct {
    HID  *string `json:"hid" validate:"required"`
    Name *string `json:"name" validate:"required"`
    Size *int    `json:"size" validate:"required"`
}

func main() {
    app := fiber.New()
    app.Get("/test", TestHandler)
    app.Listen(":9091")
}

func TestHandler(c *fiber.Ctx) error {
    user1Account := "saltfish"
    user1Password := "123456"
    user1Nickname := "咸鱼"
    user1 := &User{Account: &user1Account, Password: &user1Password, Nickname: &user1Nickname, Address: nil}
    user2Account := "saltfishpr"
    user2Password := "123456"
    user2Nickname := "咸鱼硕"
    user2 := &User{Account: &user2Account, Password: &user2Password, Nickname: &user2Nickname, Address: nil}
    hubID := "alpha"
    hubName := "聊天室1"
    hubSize := 100
    hub := &Hub{HID: &hubID, Name: &hubName, Size: &hubSize}
    hubs := []*Hub{hub}
    friends := []*User{user2}
    user1.Hubs = hubs
    user1.Friends = friends

    data, _ := json.Marshal(user1)
    log.Println(string(data))

    return c.Status(fiber.StatusOK).JSON(user1)
}
```
