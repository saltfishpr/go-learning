## 427. 建立四叉树

给你一个 `n*n` 矩阵 grid，矩阵由若干 `0` 和 `1` 组成。请你用四叉树表示该矩阵 `grid`。

你需要返回能表示矩阵的四叉树的根结点。

注意，当 `isLeaf` 为 `false` 时，你可以把 `true` 或者 `false` 赋值给节点，两种值都会被判题机制接受。

四叉树数据结构中，每个内部节点只有四个子节点。此外，每个节点都有两个属性：

- `Val`：储存叶子结点所代表的区域的值。`1` 对应 `true`，`0` 对应 `false`；
- `IsLeaf`: 当这个节点是一个叶子结点时为 `true`，如果它有 4 个子节点则为 `false` 。

**四叉树格式：**

输出为使用层序遍历后四叉树的序列化形式，其中 `null` 表示路径终止符，其下面不存在节点。

它与二叉树的序列化非常相似。唯一的区别是节点以列表形式表示 `[IsLeaf, Val]`。

如果 `IsLeaf` 或者 `Val` 的值为 `true` ，则表示它在列表  `[IsLeaf, Val]` 中的值为 `1`；如果 `IsLeaf` 或者 `Val` 的值为 `false` ，则表示值为 `0`。

### Think

正方形由左上角 `A(x0, y0)` 和右上角 `B(x1, y1)` 两个点确定

判断正方形是否是叶子结点，只需要判断正方形内元素的和是否等于正方形面积

preSum 指 `(0, 0)` 与 `(x, y)` 包含的矩形面积

则正方形 AB 的面积可以由下面的公式计算：

`Area(x0, y0, x1, y1) = preSum(x1, y1) + preSum(x0, y0) - preSum(x1, y0) - preSum(x0, y1)`

### Solution

```go
package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	m, n := len(grid), len(grid[0])
	preSum := make([][]int, m+1)
	preSum[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		preSum[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			preSum[i+1][j+1] = preSum[i+1][j] + preSum[i][j+1] - preSum[i][j] + grid[i][j]
		}
	}

	var dfs func(x0, y0, x1, y1 int) *Node
	dfs = func(x0, y0, x1, y1 int) *Node {
		if diff := preSum[x1][y1] - preSum[x1][y0] - preSum[x0][y1] + preSum[x0][y0]; diff == 0 {
			return &Node{false, true, nil, nil, nil, nil}
		} else if diff == (x1-x0)*(y1-y0) {
			return &Node{Val: true, IsLeaf: true, TopLeft: nil, TopRight: nil, BottomLeft: nil, BottomRight: nil}
		}
		hx, hy := (x0+x1)/2, (y0+y1)/2
		return &Node{
			Val:         true,
			IsLeaf:      false,
			TopLeft:     dfs(x0, y0, hx, hy),
			TopRight:    dfs(x0, hy, hx, y1),
			BottomLeft:  dfs(hx, y0, x1, hy),
			BottomRight: dfs(hx, hy, x1, y1),
		}
	}

	return dfs(0, 0, m, n)
}
```

### Tests

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_construct(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "示例 1",
			args: args{grid: [][]int{
				{0, 1},
				{1, 0},
			}},
			want: &Node{
				Val:         true,
				IsLeaf:      false,
				TopLeft:     &Node{Val: false, IsLeaf: true},
				TopRight:    &Node{Val: true, IsLeaf: true},
				BottomLeft:  &Node{Val: true, IsLeaf: true},
				BottomRight: &Node{Val: false, IsLeaf: true},
			},
		},
		{
			name: "",
			args: args{grid: [][]int{
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
			}},
			want: &Node{
				Val: true, IsLeaf: false,
				TopLeft: &Node{Val: true, IsLeaf: true},
				TopRight: &Node{
					Val: true, IsLeaf: false,
					TopLeft:     &Node{Val: false, IsLeaf: true},
					TopRight:    &Node{Val: false, IsLeaf: true},
					BottomLeft:  &Node{Val: true, IsLeaf: true},
					BottomRight: &Node{Val: true, IsLeaf: true},
				},
				BottomLeft:  &Node{Val: true, IsLeaf: true},
				BottomRight: &Node{Val: false, IsLeaf: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, construct(tt.args.grid))
		})
	}
}

var result *Node

func Benchmark_construct(b *testing.B) {
	grid := [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	var res *Node
	for n := 0; n < b.N; n++ {
		res = construct(grid)
	}
	result = res
}
```

```plaintext
=== RUN   Test_construct
=== RUN   Test_construct/示例_1
=== RUN   Test_construct/#00
--- PASS: Test_construct (0.00s)
    --- PASS: Test_construct/示例_1 (0.00s)
    --- PASS: Test_construct/#00 (0.00s)
PASS
ok  	leetcode/427	0.439s
```

```plaintext
goos: darwin
goarch: arm64
pkg: leetcode/427
Benchmark_construct
Benchmark_construct-10    	 2237152	       535.4 ns/op	    1376 B/op	      19 allocs/op
PASS
ok  	leetcode/427	1.843s
```
