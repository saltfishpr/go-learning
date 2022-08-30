## 109. 有序链表转换二叉搜索树

给定一个单链表的头节点 `head`，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点的左右两个子树的高度差不超过 `1`。

### Think

添加辅助函数: 从链表的 `left` 和 `right` 构建二叉搜索树

先找到中间节点，中间节点作为二叉搜索树的根，再对左右链表递归调用辅助函数

### Solution

```go
func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func getMedian(left, right *ListNode) *ListNode {
	fast, slow := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// buildTree builds a bst from a sorted list.
// [left, right) is the range of the list.
func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMedian(left, right)
	root := &TreeNode{mid.Val, nil, nil}
	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)
	return root
}
```

### Tests

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedListToBST(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "示例 1",
			args: args{
				head: &ListNode{
					Val: -10,
					Next: &ListNode{
						Val: -3,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 9,
								},
							},
						},
					},
				},
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:  -3,
					Left: &TreeNode{Val: -10},
				},
				Right: &TreeNode{
					Val:  9,
					Left: &TreeNode{Val: 5},
				},
			},
		},
		{
			name: "示例 2",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sortedListToBST(tt.args.head))
		})
	}
}

var result *TreeNode

func Benchmark_sortedListToBST(b *testing.B) {
	head := &ListNode{
		Val: -10,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		},
	}

	var res *TreeNode
	for n := 0; n < b.N; n++ {
		res = sortedListToBST(head)
	}
	result = res
}
```
