## 501. 二叉搜索树中的众数

给你一个含重复值的二叉搜索树（BST）的根节点 `root`，找出并返回 `BST` 中的所有众数（即，出现频率最高的元素）。

如果树中有不止一个众数，可以按**任意顺序**返回。

假定 BST 满足如下定义：

- 结点左子树中所含节点的值**小于等于**当前节点的值
- 结点右子树中所含节点的值**大于等于**当前节点的值
- 左子树和右子树都是二叉搜索树

### Solution

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	m := make(map[int]int)
	var mode int

	var helper func(*TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		m[root.Val]++
		if m[root.Val] > mode {
			mode = m[root.Val]
		}
		helper(root.Right)
	}

	helper(root)

	var res []int
	for k, v := range m {
		if v == mode {
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

func Test_findMode(t *testing.T) {
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
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 2},
					},
				},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findMode(tt.args.root))
		})
	}
}

var result []int

func Benchmark_findMode(b *testing.B) {
	var res []int

	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 2},
		},
	}

	for n := 0; n < b.N; n++ {
		res = findMode(root)
	}

	result = res
}
```

```plaintext
=== RUN   Test_findMode
=== RUN   Test_findMode/示例_1
--- PASS: Test_findMode (0.00s)
    --- PASS: Test_findMode/示例_1 (0.00s)
PASS
coverage: 100.0% of statements
ok  	learning/501	0.002s	coverage: 100.0% of statements
```

```plaintext
goos: linux
goarch: amd64
pkg: learning/501
cpu: AMD Ryzen 7 5800H with Radeon Graphics
Benchmark_findMode
Benchmark_findMode-16    	 7676437	       159.0 ns/op	      56 B/op	       3 allocs/op
PASS
coverage: 100.0% of statements
ok  	learning/501	1.383s
```
