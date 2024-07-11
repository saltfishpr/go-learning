## 437. 路径总和 III

给定一个二叉树的根节点 `root`，和一个整数 `targetSum`，求该二叉树里节点值之和等于 `targetSum` 的**路径**的数目。

**路径**不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

### Think

dfs 过程中记录前缀和到 map 中

### Solution

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) (ans int) {
	preSum := map[int]int{0: 1}

	var helper func(*TreeNode, int)
	// cur 当前的前缀和
	helper = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		cur += node.Val
		// 是否存在前缀和刚好等于 cur−targetSum
		ans += preSum[cur-targetSum]
		preSum[cur]++
		helper(node.Left, cur)
		helper(node.Right, cur)
		// 回溯
		preSum[cur]--
	}
	helper(root, 0)

	return
}
```

### Tests

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var root1 = &TreeNode{
	Val: 10,
	Left: &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: -2},
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 1},
		},
	},
	Right: &TreeNode{
		Val:   -3,
		Right: &TreeNode{Val: 11},
	},
}

var root2 = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   11,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 2},
		},
	},
	Right: &TreeNode{
		Val:  8,
		Left: &TreeNode{Val: 13},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 1},
		},
	},
}

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例 1",
			args: args{
				root:      root1,
				targetSum: 8,
			},
			want: 3,
		},
		{
			name: "示例 2",
			args: args{
				root:      root2,
				targetSum: 22,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, pathSum(tt.args.root, tt.args.targetSum))
		})
	}
}

var result int

func Benchmark_pathSum(b *testing.B) {
	var res int
	for n := 0; n < b.N; n++ {
		res = pathSum(root2, 22)
	}
	result = res
}
```

```plaintext
=== RUN   Test_pathSum
=== RUN   Test_pathSum/示例_1
=== RUN   Test_pathSum/示例_2
--- PASS: Test_pathSum (0.00s)
    --- PASS: Test_pathSum/示例_1 (0.00s)
    --- PASS: Test_pathSum/示例_2 (0.00s)
PASS
coverage: 100.0% of statements
ok  	leetcode/437	0.005s	coverage: 100.0% of statements
```

```plaintext
goos: linux
goarch: amd64
pkg: leetcode/437
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
Benchmark_pathSum
Benchmark_pathSum-4   	 1287446	       876.5 ns/op	     292 B/op	       1 allocs/op
PASS
coverage: 100.0% of statements
ok  	leetcode/437	2.082s
```
