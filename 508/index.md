## 508. 出现次数最多的子树元素和

给你一个二叉树的根结点 `root`，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。

一个结点的 **「子树元素和」** 定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。

### Solution

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	m := make(map[int]int)

	var helper func(*TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		val := helper(root.Left) + helper(root.Right) + root.Val
		m[val]++
		return val
	}

	helper(root)

	var maxCount int
	for _, v := range m {
		if v > maxCount {
			maxCount = v
		}
	}

	var res []int
	for k, v := range m {
		if v == maxCount {
			res = append(res, k)
		}
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

func Test_findFrequentTreeSum(t *testing.T) {
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
					Val:   5,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: -3},
				},
			},
			want: []int{2, -3, 4},
		},
		{
			name: "示例 2",
			args: args{
				root: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: -5},
				},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, findFrequentTreeSum(tt.args.root))
		})
	}
}

var result []int

func Benchmark_findFrequentTreeSum(b *testing.B) {
	var res []int

	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: -3},
	}

	for n := 0; n < b.N; n++ {
		res = findFrequentTreeSum(root)
	}

	result = res
}
```

```plaintext
=== RUN   Test_findFrequentTreeSum
=== RUN   Test_findFrequentTreeSum/示例_1
=== RUN   Test_findFrequentTreeSum/示例_2
--- PASS: Test_findFrequentTreeSum (0.00s)
    --- PASS: Test_findFrequentTreeSum/示例_1 (0.00s)
    --- PASS: Test_findFrequentTreeSum/示例_2 (0.00s)
PASS
coverage: 100.0% of statements
ok  	learning/508	0.002s	coverage: 100.0% of statements
```

```plaintext
goos: linux
goarch: amd64
pkg: learning/508
cpu: AMD Ryzen 7 5800H with Radeon Graphics
Benchmark_findFrequentTreeSum
Benchmark_findFrequentTreeSum-16    	 4956513	       246.3 ns/op	     104 B/op	       5 allocs/op
PASS
coverage: 100.0% of statements
ok  	learning/508	1.471s
```
