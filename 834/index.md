## 834. 树中距离之和

给定一个无向、连通的树。树中有 `n` 个标记为 `0...n-1` 的节点以及 `n-1` 条边 。

给定整数 `n` 和数组 `edges`，`edges[i] = [ai, bi]` 表示树中的节点 `ai` 和 `bi` 之间有一条边。

返回长度为 `n` 的数组 `answer`，其中 `answer[i]` 是树中第 `i` 个节点与所有其他节点之间的距离之和。

### Think

回溯，换根，TODO

### Solution

```go
package main

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make([][]int, n)
	for _, edge := range edges {
		i, j := edge[0], edge[1]
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], i)
	}

	sz := make([]int, n) // 存储子节点数(包括自己)
	dp := make([]int, n) // dp[i] 表示以 i 为根的树，所有子节点到它的距离之和
	var dfs func(int, int)
	// u 当前节点，f 父节点
	dfs = func(u, f int) {
		sz[u] = 1
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			dfs(v, u)
			dp[u] += dp[v] + sz[v]
			sz[u] += sz[v]
		}
	}
	dfs(0, -1)

	res := make([]int, n) // 存储最终结果
	var dfs2 func(int, int)
	dfs2 = func(u, f int) {
		res[u] = dp[u]
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			// 存下当前结果
			pu, pv := dp[u], dp[v]
			su, sv := sz[u], sz[v]

			// 换根
			dp[u] -= dp[v] + sz[v]
			sz[u] -= sz[v]
			dp[v] += dp[u] + sz[u]
			sz[v] += sz[u]

			// 当前状态，v 已经为根节点，再进一步将 v 的子节点作为根节点，计算 ans
			dfs2(v, u)

			// 恢复
			dp[u], dp[v] = pu, pv
			sz[u], sz[v] = su, sv
		}
	}

	dfs2(0, -1)
	return res
}
```

### Tests

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumOfDistancesInTree(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例 1",
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}},
			},
			want: []int{8, 12, 6, 10, 10, 10},
		},
		{
			name: "示例 2",
			args: args{
				n:     1,
				edges: [][]int{},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sumOfDistancesInTree(tt.args.n, tt.args.edges))
		})
	}
}

var result []int

func Benchmark_sumOfDistancesInTree(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = sumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}})
	}
	result = res
}
```

```plaintext
=== RUN   Test_sumOfDistancesInTree
=== RUN   Test_sumOfDistancesInTree/示例_1
=== RUN   Test_sumOfDistancesInTree/示例_2
--- PASS: Test_sumOfDistancesInTree (0.00s)
    --- PASS: Test_sumOfDistancesInTree/示例_1 (0.00s)
    --- PASS: Test_sumOfDistancesInTree/示例_2 (0.00s)
PASS
coverage: 100.0% of statements
ok  	leetcode/834	0.004s	coverage: 100.0% of statements
```

```plaintext
goos: linux
goarch: amd64
pkg: leetcode/834
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
Benchmark_sumOfDistancesInTree
Benchmark_sumOfDistancesInTree-4   	 1729839	       680.4 ns/op	     400 B/op	      13 allocs/op
PASS
coverage: 100.0% of statements
ok  	leetcode/834	1.888s
```
