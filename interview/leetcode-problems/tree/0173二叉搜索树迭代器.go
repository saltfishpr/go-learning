// @file: 0173二叉搜索树迭代器.go
// @description:
// @author: SaltFish
// @date: 2020/09/10

package tree

/*
实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
调用 next() 将返回二叉搜索树中的下一个最小的数。

	7
   / \
  3  15
    /  \
   9   20

BSTIterator iterator = new BSTIterator(root);
iterator.next();    // 返回 3
iterator.next();    // 返回 7
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 9
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 15
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 20
iterator.hasNext(); // 返回 false

提示：
next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。

*/

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	bt := new(BSTIterator)
	bt.leftmostInorder(root)
	return *bt
}

func (bst *BSTIterator) leftmostInorder(root *TreeNode) {
	for root != nil {
		bst.stack = append(bst.stack, root)
		root = root.Left
	}
}

/*Next return the next smallest number */
func (bst *BSTIterator) Next() int {
	topmostNode := bst.stack[len(bst.stack)-1]
	bst.stack = bst.stack[:len(bst.stack)-1]
	if topmostNode.Right != nil {
		bst.leftmostInorder(topmostNode.Right)
	}
	return topmostNode.Val
}

/*HasNext return whether we have a next smallest number */
func (bst *BSTIterator) HasNext() bool {
	return len(bst.stack) > 0
}
