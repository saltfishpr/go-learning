// @file: 0114二叉树展开为链表.go
// @description: 给定一个二叉树，原地将它展开为一个单链表。
// @author: SaltFish
// @date: 2020/09/04

package tree

/*
例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6

将其展开为：
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6

*/

func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}
