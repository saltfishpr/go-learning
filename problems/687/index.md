## 687. 最长同值路径

给定一个二叉树的 `root`，返回最长的路径的长度，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。

两个节点之间的路径长度由它们之间的边数表示。

### Think

DFS，后续遍历

当前节点的最长同值路径 = 左子节点的最长同值路径 + 右子节点的最长同值路径

节点作为子节点时的最长同值路径 = max(左子节点的最长同值路径, 右子节点的最长同值路径)

要获得最长同值路径，我们需要一个值来保存**最长**的值，

### Solution

```go
func longestUnivaluePath(root *TreeNode) int {
	var res int

	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := helper(root.Left)
		right := helper(root.Right)
		var left_res, right_res int
		if root.Left != nil && root.Left.Val == root.Val {
			left_res = left + 1
		}
		if root.Right != nil && root.Right.Val == root.Val {
			right_res = right + 1
		}
		res = max(res, left_res+right_res)
		return max(left_res, right_res)
	}

	helper(root)
	return res
}

func max(a, b int) int {
	if a > b {
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

func Test_longestUnivaluePath(t *testing.T) {
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
					Val: 5,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 1},
					},
					Right: &TreeNode{
						Val:   5,
						Right: &TreeNode{Val: 5},
					},
				},
			},
			want: 2,
		},
		{
			name: "示例 2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   5,
						Right: &TreeNode{Val: 5},
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestUnivaluePath(tt.args.root))
		})
	}
}

var result int

func Benchmark_longestUnivaluePath(b *testing.B) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 5},
		},
	}

	var res int
	for i := 0; i < b.N; i++ {
		res = longestUnivaluePath(root)
	}
	result = res
}
```

```plaintext
=== RUN   Test_longestUnivaluePath
=== RUN   Test_longestUnivaluePath/示例_1
=== RUN   Test_longestUnivaluePath/示例_2
--- PASS: Test_longestUnivaluePath (0.00s)
    --- PASS: Test_longestUnivaluePath/示例_1 (0.00s)
    --- PASS: Test_longestUnivaluePath/示例_2 (0.00s)
PASS
coverage: 100.0% of statements
ok  	leetcode/687	0.031s	coverage: 100.0% of statements
```

```plaintext
goos: linux
goarch: amd64
pkg: leetcode/687
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
Benchmark_longestUnivaluePath
Benchmark_longestUnivaluePath-4   	  184054	      6774 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 89.5% of statements
ok  	leetcode/687	1.343s
```
