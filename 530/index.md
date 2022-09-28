## 530. 二叉搜索树的最小绝对差

给你一个二叉搜索树的根节点 `root`，返回 树中任意两不同节点值之间的最小差值。

差值是一个正数，其数值等于两值之差的绝对值。

### Think

中序遍历，比较当前值和前一个值的差

### Solution

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	pre := 10000
	min := 10000

	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		now := abs(root.Val - pre)
		if now < min {
			min = now
		}
		pre = root.Val
		helper(root.Right)
	}

	helper(root)

	return min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

### Tests

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMinimumDifference(t *testing.T) {
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
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 6},
				},
			},
			want: 1,
		},
		{
			name: "示例 2",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   48,
						Left:  &TreeNode{Val: 12},
						Right: &TreeNode{Val: 49},
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getMinimumDifference(tt.args.root))
		})
	}
}

var result int

func Benchmark_getMinimumDifference(b *testing.B) {
	var res int

	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 0},
		Right: &TreeNode{
			Val:   48,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 49},
		},
	}

	for n := 0; n < b.N; n++ {
		res = getMinimumDifference(root)
	}

	result = res
}
```

```plaintext
=== RUN   Test_getMinimumDifference
=== RUN   Test_getMinimumDifference/示例_1
=== RUN   Test_getMinimumDifference/示例_2
--- PASS: Test_getMinimumDifference (0.00s)
    --- PASS: Test_getMinimumDifference/示例_1 (0.00s)
    --- PASS: Test_getMinimumDifference/示例_2 (0.00s)
PASS
coverage: 100.0% of statements
ok  	learning/530	0.020s	coverage: 100.0% of statements
```

```plaintext
goos: linux
goarch: amd64
pkg: learning/530
cpu: AMD Ryzen 7 5800H with Radeon Graphics
Benchmark_getMinimumDifference
Benchmark_getMinimumDifference-16    	  349310	      3352 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 100.0% of statements
ok  	learning/530	1.236s
```
