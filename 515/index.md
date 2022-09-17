## 515. 在每个树行中找最大值

给定一棵二叉树的根节点 `root`，请找出该二叉树中每一层的最大值。

### Think

层序遍历

### Solution

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		max := queue[0].Val
		var nextQueue []*TreeNode
		for _, node := range queue {
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		res = append(res, max)
		queue = nextQueue
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

func Test_largestValues(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例 1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 9},
					},
				},
			},
			want: []int{1, 3, 9},
		},
		{
			name: "示例 2",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
			},
			want: []int{1, 3},
		},
		{
			name: "示例 3",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, largestValues(tt.args.root))
		})
	}
}

var result []int

func Benchmark_largestValues(b *testing.B) {
	var res []int

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 9},
		},
	}

	for n := 0; n < b.N; n++ {
		res = largestValues(root)
	}

	result = res
}
```

```plaintext
=== RUN   Test_largestValues
=== RUN   Test_largestValues/示例_1
=== RUN   Test_largestValues/示例_2
=== RUN   Test_largestValues/示例_3
--- PASS: Test_largestValues (0.00s)
    --- PASS: Test_largestValues/示例_1 (0.00s)
    --- PASS: Test_largestValues/示例_2 (0.00s)
    --- PASS: Test_largestValues/示例_3 (0.00s)
PASS
coverage: 100.0% of statements
ok  	learning/515	0.004s	coverage: 100.0% of statements
```

```plaintext
goos: linux
goarch: amd64
pkg: learning/515
cpu: AMD Ryzen 7 5800H with Radeon Graphics
Benchmark_largestValues
Benchmark_largestValues-16    	 3543174	       355.2 ns/op	     136 B/op	       8 allocs/op
PASS
coverage: 94.1% of statements
ok  	learning/515	1.612s
```
