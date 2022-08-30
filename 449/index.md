## 449. 序列化和反序列化二叉搜索树

序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。

设计一个算法来序列化和反序列化二叉搜索树。对序列化/反序列化算法的工作方式没有限制。您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。

编码的字符串应尽可能紧凑。

### Think

二叉搜索树的中序遍历是有序的；根据一个(前序、中序)或者(中序，后序)遍历结果，可以还原一棵二叉树

因此，我们有二叉搜索树的前序或后序遍历结果，即可还原这颗二叉树

前序遍历的第一个元素为根结点的值，小于根结点值的在左子树，大于根结点值的在右子树

### Solution

```go
package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (*Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var preOrder []string

	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		preOrder = append(preOrder, strconv.Itoa(root.Val))
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)

	return strings.Join(preOrder, ",")
}

// Deserializes your encoded data to tree.
func (*Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	ss := strings.Split(data, ",")
	preOrder := make([]int, len(ss))
	for i, s := range ss {
		val, _ := strconv.Atoi(s)
		preOrder[i] = val
	}

	var helper func(left, right int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		val := preOrder[left]
		mid := right + 1
		root := &TreeNode{val, nil, nil}
		for i := left + 1; i <= right; i++ {
			if preOrder[i] > val {
				mid = i
				break
			}
		}
		root.Left = helper(left+1, mid-1)
		root.Right = helper(mid, right)
		return root
	}

	return helper(0, len(preOrder)-1)
}
```

### Tests

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var c = Constructor()

func TestCodec(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
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
		},
		{
			name: "示例 2",
			args: args{
				root: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := c.serialize(tt.args.root)
			root := c.deserialize(data)
			assert.Equal(t, tt.args.root, root)
		})
	}
}

var (
	result1 string
	result2 *TreeNode

	root = &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
)

func BenchmarkCodec_serialize(b *testing.B) {
	var res string
	for n := 0; n < b.N; n++ {
		res = c.serialize(root)
	}
	result1 = res
}

func BenchmarkCodec_deserialize(b *testing.B) {
	var res *TreeNode
	for n := 0; n < b.N; n++ {
		res = c.deserialize("2,1,3")
	}
	result2 = res
}
```

```plaintext
=== RUN   TestCodec
=== RUN   TestCodec/示例_1
=== RUN   TestCodec/示例_2
--- PASS: TestCodec (0.00s)
    --- PASS: TestCodec/示例_1 (0.00s)
    --- PASS: TestCodec/示例_2 (0.00s)
PASS
coverage: 100.0% of statements
ok  	leetcode/449	0.552s	coverage: 100.0% of statements
```

```plaintext
goos: darwin
goarch: arm64
pkg: leetcode/449
BenchmarkCodec_serialize
BenchmarkCodec_serialize-10      	 8954449	       128.4 ns/op	     117 B/op	       4 allocs/op
BenchmarkCodec_deserialize
BenchmarkCodec_deserialize-10    	 8484712	       139.5 ns/op	     144 B/op	       5 allocs/op
PASS
coverage: 94.3% of statements
ok  	leetcode/449	2.745s
```
