## 429. N 叉树的层序遍历

给定一个 `N` 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 `null` 值分隔（参见示例）。

### Solution

```go
package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var queue []*Node
	queue = append(queue, root)
	var res [][]int
	for len(queue) != 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			node := queue[i]
			temp = append(temp, node.Val)
			queue = append(queue, node.Children...)
		}
		res = append(res, temp)
		queue = queue[length:]
	}
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

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "示例 1",
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{
							Val: 3,
							Children: []*Node{
								{Val: 5},
								{Val: 6},
							},
						},
						{Val: 2},
						{Val: 4},
					},
				},
			},
			want: [][]int{{1}, {3, 2, 4}, {5, 6}},
		},
		{
			name: "示例 2",
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{Val: 2},
						{
							Val: 3,
							Children: []*Node{
								{Val: 6},
								{
									Val: 7,
									Children: []*Node{
										{
											Val: 11,
											Children: []*Node{
												{Val: 14},
											},
										},
									},
								},
							},
						},
						{
							Val: 4,
							Children: []*Node{
								{
									Val: 8,
									Children: []*Node{
										{Val: 12},
									},
								},
							},
						},
						{
							Val: 5,
							Children: []*Node{
								{
									Val: 9,
									Children: []*Node{
										{Val: 13},
									},
								},
								{Val: 10},
							},
						},
					},
				},
			},
			want: [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, levelOrder(tt.args.root))
		})
	}
}

var result [][]int

func Benchmark_levelOrder(b *testing.B) {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
			},
			{Val: 2},
			{Val: 4},
		},
	}
	var res [][]int
	for n := 0; n < b.N; n++ {
		res = levelOrder(root)
	}
	result = res
}
```

```plaintext
=== RUN   Test_levelOrder
=== RUN   Test_levelOrder/示例_1
=== RUN   Test_levelOrder/示例_2
--- PASS: Test_levelOrder (0.00s)
    --- PASS: Test_levelOrder/示例_1 (0.00s)
    --- PASS: Test_levelOrder/示例_2 (0.00s)
PASS
coverage: 93.3% of statements
ok  	leetcode/429	0.506s	coverage: 93.3% of statements
```

```plaintext
goos: darwin
goarch: arm64
pkg: leetcode/429
Benchmark_levelOrder
Benchmark_levelOrder-10    	 3533253	       330.1 ns/op	     344 B/op	      12 allocs/op
PASS
coverage: 93.3% of statements
ok  	leetcode/429	1.811s
```
