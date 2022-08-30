## 787. K 站中转内最便宜的航班

有 `n` 个城市通过一些航班连接。给你一个数组 `flights`，其中 `flights[i] = [fromi, toi, pricei]`，表示该航班都从城市 `fromi` 开始，以价格 `pricei` 抵达 toi。

现在给定所有的城市和航班，以及出发城市 `src` 和目的地 `dst`，你的任务是找到出一条最多经过 `k` 站中转的路线，使得从 `src` 到 `dst` 的**价格最便宜**，并返回该价格。如果不存在这样的路线，则输出 `-1`。

### Think

动态规划，TODO

### Solution

```go
package main

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 10000*101 + 1

	// dp[t][i] 表示通过恰好 t 次航班，从出发城市 src 到达城市 i 需要的最小花费
	// k 次中转，也就是最多搭乘 k+1 次航班，再加上 base state，0 <= t < k+2
	dp := make([][]int, k+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][src] = 0

	for t := 1; t < k+2; t++ {
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			dp[t][i] = min(dp[t][i], dp[t-1][j]+cost)
		}
	}

	ans := inf
	for t := 1; t <= k+1; t++ {
		ans = min(ans, dp[t][dst])
	}
	if ans == inf {
		ans = -1
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

### Tests

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findCheapestPrice(t *testing.T) {
	type args struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例 1",
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
				src:     0,
				dst:     2,
				k:       1,
			},
			want: 200,
		},
		{
			name: "示例 2",
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
				src:     0,
				dst:     2,
				k:       0,
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findCheapestPrice(tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.k))
		})
	}
}

var result int

func Benchmark_findCheapestPrice(b *testing.B) {
	var res int
	for n := 0; n < b.N; n++ {
		res = findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1)
	}
	result = res
}
```

```plaintext
=== RUN   Test_findCheapestPrice
=== RUN   Test_findCheapestPrice/示例_1
=== RUN   Test_findCheapestPrice/示例_2
--- PASS: Test_findCheapestPrice (0.00s)
    --- PASS: Test_findCheapestPrice/示例_1 (0.00s)
    --- PASS: Test_findCheapestPrice/示例_2 (0.00s)
PASS
coverage: 95.0% of statements
ok  	leetcode/787	0.003s	coverage: 95.0% of statements
```

```plaintext
goos: linux
goarch: amd64
pkg: leetcode/787
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
Benchmark_findCheapestPrice
Benchmark_findCheapestPrice-4   	 4898914	       238.1 ns/op	     152 B/op	       4 allocs/op
PASS
coverage: 95.0% of statements
ok  	leetcode/787	1.425s
```
