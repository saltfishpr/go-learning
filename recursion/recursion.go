// @file: recursion.go
// @date: 2021/1/16

// Package recursion
package recursion

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 344. 反转字符串
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

// 24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	var swap func(*ListNode) *ListNode
	swap = func(root *ListNode) *ListNode {
		if root == nil || root.Next == nil {
			return root
		}
		root, root.Next, root.Next.Next = root.Next, swap(root.Next.Next), root
		return root
	}
	return swap(head)
}

// 95. 不同的二叉搜索树 II
func generateTrees(n int) []*TreeNode {
	var generate func(int, int) []*TreeNode
	generate = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		res := make([]*TreeNode, 0)
		// 对根节点循环
		for i := start; i <= end; i++ {
			lefts, rights := generate(start, i-1), generate(i+1, end)
			for j := 0; j < len(lefts); j++ {
				for k := 0; k < len(rights); k++ {
					root := &TreeNode{Val: i, Left: lefts[j], Right: rights[k]}
					res = append(res, root)
				}
			}
		}
		return res
	}

	return generate(1, n)
}
