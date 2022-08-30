## 404. 左叶子之和

给定二叉树的根节点 `root`，返回所有左叶子之和。

### Think

递归

### Solution

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res int

	if isLeaf(root.Left) {
		res += root.Left.Val
	} else {
		res += sumOfLeftLeaves(root.Left)
	}

	res += sumOfLeftLeaves(root.Right)

	return res
}

func isLeaf(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return root.Left == nil && root.Right == nil
}
```

### Tests

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumOfLeftLeaves(t *testing.T) {
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
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: 24,
		},
		{
			name: "示例 2",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sumOfLeftLeaves(tt.args.root))
		})
	}
}

var result int

func Benchmark_sumOfLeftLeaves(b *testing.B) {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	var res int
	for n := 0; n < b.N; n++ {
		res = sumOfLeftLeaves(root)
	}
	result = res
}
```

```plaintext
=== RUN   Test_sumOfLeftLeaves
=== RUN   Test_sumOfLeftLeaves/示例_1
=== RUN   Test_sumOfLeftLeaves/示例_2
--- PASS: Test_sumOfLeftLeaves (0.00s)
    --- PASS: Test_sumOfLeftLeaves/示例_1 (0.00s)
    --- PASS: Test_sumOfLeftLeaves/示例_2 (0.00s)
PASS
coverage: 100.0% of statements
ok  	leetcode/404	0.002s	coverage: 100.0% of statements
```

```plaintext
goos: linux
goarch: amd64
pkg: leetcode/404
cpu: AMD Ryzen 7 5800H with Radeon Graphics
Benchmark_sumOfLeftLeaves
Benchmark_sumOfLeftLeaves-16    	125321624	         9.607 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 100.0% of statements
ok  	leetcode/404	2.176s
```
