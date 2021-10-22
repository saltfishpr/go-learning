// @file: 0199二叉树的右视图.go
// @description: 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
// @author: SaltFish
// @date: 2020/09/10

package tree

/*
示例:
输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---

*/

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nodeMap := make(map[int]int)
	var maxDepth int = -1
	var nodeStack []*TreeNode
	var depthStack []int

	push := func(node *TreeNode, depth int) {
		nodeStack = append(nodeStack, root)
		depthStack = append(depthStack, depth)
	}
	pop := func() (n *TreeNode, i int) {
		n = nodeStack[len(nodeStack)-1]
		i = depthStack[len(depthStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		depthStack = depthStack[:len(depthStack)-1]
		return
	}

	push(root, 0)
	for len(nodeStack) > 0 {
		node, depth := pop()
		if node != nil {
			maxDepth = max(maxDepth, depth)
			if _, ok := nodeMap[depth]; !ok {
				nodeMap[depth] = node.Val // 不存在则插入
			}
			push(node.Left, depth+1)
			push(node.Right, depth+1)
		}
	}
	var res []int
	for i := 0; i < maxDepth+1; i++ {
		res = append(res, nodeMap[i])
	}
	return res
}

func rightSideView2(root *TreeNode) []int {
	var dfs func(*TreeNode, int)
	var res []int
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth == len(res) { // 从右节点到达一个新的深度，添加这个节点
			res = append(res, root.Val)
		}
		depth++
		dfs(root.Right, depth)
		dfs(root.Left, depth)
	}
	dfs(root, 0)
	return res
}
