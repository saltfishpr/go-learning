## 513. 找树左下角的值

给定一个二叉树的根节点 `root`，请找出该二叉树的 **最底层** **最左边** 节点的值。

假设二叉树中至少有一个节点。

### Think

广度优先搜索

### Solution

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var res int

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		res = queue[0].Val
		var nextQueue []*TreeNode
		for _, node := range queue {
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		if len(nextQueue) == 0 {
			break
		}
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

func Test_findBottomLeftValue(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例 1",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
			want: 1,
		},
		{
			name: "示例 2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:  5,
							Left: &TreeNode{Val: 7},
						},
						Right: &TreeNode{Val: 6},
					},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findBottomLeftValue(tt.args.root))
		})
	}
}

var result int

func Benchmark_findBottomLeftValue(b *testing.B) {
	var res int

	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}

	for n := 0; n < b.N; n++ {
		res = findBottomLeftValue(root)
	}

	result = res
}
```

```plaintext
=== RUN   Test_findBottomLeftValue
=== RUN   Test_findBottomLeftValue/示例_1
=== RUN   Test_findBottomLeftValue/示例_2
--- PASS: Test_findBottomLeftValue (0.00s)
    --- PASS: Test_findBottomLeftValue/示例_1 (0.00s)
    --- PASS: Test_findBottomLeftValue/示例_2 (0.00s)
PASS
coverage: 100.0% of statements
ok  	learning/513	0.002s	coverage: 100.0% of statements
```

```plaintext
goos: linux
goarch: amd64
pkg: learning/513
cpu: AMD Ryzen 7 5800H with Radeon Graphics
Benchmark_findBottomLeftValue
Benchmark_findBottomLeftValue-16    	24237619	        49.57 ns/op	      24 B/op	       2 allocs/op
PASS
coverage: 100.0% of statements
ok  	learning/513	1.257s
```
