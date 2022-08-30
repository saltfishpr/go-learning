## 1396. 设计地铁系统

地铁系统跟踪不同车站之间的乘客出行时间，并使用这一数据来计算从一站到另一站的平均时间。

实现 `UndergroundSystem` 类：

- `void checkIn(int id, string stationName, int t)`
  通行卡 ID 等于 `id` 的乘客，在时间 `t` ，从 `stationName` 站进入
  乘客一次只能从一个站进入
- `void checkOut(int id, string stationName, int t)`
  通行卡 ID 等于 `id` 的乘客，在时间 `t` ，从 `stationName` 站离开
- `double getAverageTime(string startStation, string endStation)`
  返回从 `startStation` 站到 `endStation` 站的平均时间
  平均时间会根据截至目前所有从 `startStation` 站 直接 到达 `endStation` 站的行程进行计算，也就是从 `startStation` 站进入并从 `endStation` 离开的行程
  从 `startStation` 到 `endStation` 的行程时间与从 `endStation` 到 `startStation` 的行程时间可能不同
  在调用 `getAverageTime` 之前，至少有一名乘客从 `startStation` 站到达 `endStation` 站

你可以假设对 `checkIn` 和 `checkOut` 方法的所有调用都是符合逻辑的。如果一名乘客在时间 `t1` 进站、时间 `t2` 出站，那么 `t1 < t2`。所有时间都按时间顺序发生。

### Think

`map[id]CheckInInfo` 保存用户的入站信息，在用户出站时计算耗时 `Cost`

`map["A->B"]Cost` 其中 `"A->B"` 为 Key，`Cost` 是包含总时间与人数的结构体

### Solution

```go
package main

import "strings"

type CheckInInfo struct {
	Station string
	Time    int
}

type Cost struct {
	Sum int // 总耗时
	Num int // 人数
}

type UndergroundSystem struct {
	check map[int]CheckInInfo
	data  map[string]Cost
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		check: make(map[int]CheckInInfo),
		data:  make(map[string]Cost),
	}
}

func (us *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	us.check[id] = CheckInInfo{Station: stationName, Time: t}
}

func (us *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	checkInInfo := us.check[id]
	travel := strings.Join([]string{checkInInfo.Station, stationName}, "->")
	oldCost := us.data[travel]
	us.data[travel] = Cost{
		Sum: oldCost.Sum + t - checkInInfo.Time,
		Num: oldCost.Num + 1,
	}
}

func (us *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	travel := strings.Join([]string{startStation, endStation}, "->")
	if cost, ok := us.data[travel]; ok {
		return float64(cost.Sum) / float64(cost.Num)
	}
	return 0
}
```

### Tests

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndergroundSystem(t *testing.T) {
	t.Run("示例 1", func(t *testing.T) {
		us := Constructor()

		us.CheckIn(45, "Leyton", 3)
		us.CheckIn(32, "Paradise", 8)
		us.CheckIn(27, "Leyton", 10)
		us.CheckOut(45, "Waterloo", 15)
		us.CheckOut(27, "Waterloo", 20)
		us.CheckOut(32, "Cambridge", 22)
		assert.Equal(t, 14., us.GetAverageTime("Paradise", "Cambridge"))
		assert.Equal(t, 11., us.GetAverageTime("Leyton", "Waterloo"))
		us.CheckIn(10, "Leyton", 24)
		assert.Equal(t, 11., us.GetAverageTime("Leyton", "Waterloo"))
		us.CheckOut(10, "Waterloo", 38)
		assert.Equal(t, 12., us.GetAverageTime("Leyton", "Waterloo"))
	})

	t.Run("示例 2", func(t *testing.T) {
		us := Constructor()

		us.CheckIn(10, "Leyton", 3)
		us.CheckOut(10, "Paradise", 8)
		assert.Equal(t, 5./1, us.GetAverageTime("Leyton", "Paradise"))
		us.CheckIn(5, "Leyton", 10)
		us.CheckOut(5, "Paradise", 16)
		assert.Equal(t, 11./2, us.GetAverageTime("Leyton", "Paradise"))

		us.CheckIn(2, "Leyton", 21)
		us.CheckOut(2, "Paradise", 30)
		assert.Equal(t, 20./3, us.GetAverageTime("Leyton", "Paradise"))
	})
}

var result float64

func BenchmarkUndergroundSystem(b *testing.B) {
	var res float64
	us := Constructor()
	startStation := "A"
	endStation := "B"
	for n := 0; n < b.N; n++ {
		us.CheckIn(n, startStation, 5)
		res = us.GetAverageTime(startStation, endStation)
		us.CheckOut(n, endStation, 10)
	}
	result = res
}
```

```plaintext
=== RUN   TestUndergroundSystem
=== RUN   TestUndergroundSystem/示例_1
=== RUN   TestUndergroundSystem/示例_2
--- PASS: TestUndergroundSystem (0.00s)
    --- PASS: TestUndergroundSystem/示例_1 (0.00s)
    --- PASS: TestUndergroundSystem/示例_2 (0.00s)
PASS
coverage: 90.0% of statements
ok  	leetcode/1396	0.004s	coverage: 90.0% of statements
```

```plaintext
goos: linux
goarch: amd64
pkg: leetcode/1396
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
BenchmarkUndergroundSystem
BenchmarkUndergroundSystem-4   	 1913522	       644.4 ns/op	     177 B/op	       2 allocs/op
PASS
coverage: 100.0% of statements
ok  	leetcode/1396	1.894s
```
