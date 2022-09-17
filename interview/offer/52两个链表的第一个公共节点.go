// @file: 52两个链表的第一个公共节点.go
// @date: 2021/2/23

// Package offer
package offer

/*
输入两个链表，找出它们的第一个公共节点。
*/

func getIntersectionNodeX52(headA, headB *ListNode) *ListNode {
	moveA, moveB := headA, headB
	for moveA != nil && moveB != nil {
		moveA, moveB = moveA.Next, moveB.Next
	}
	if moveA == nil {
		headA, headB = headB, headA
		moveA, moveB = moveB, moveA
	}
	A, B := headA, headB
	for moveA != nil {
		moveA, A = moveA.Next, A.Next
	}

	for A != nil && B != nil {
		if A == B {
			return A
		}
		A, B = A.Next, B.Next
	}
	return nil
}

func getIntersectionNodeX52V2(headA, headB *ListNode) *ListNode {
	moveA, moveB := headA, headB
	for moveA != moveB {
		if moveA != nil {
			moveA = moveA.Next
		} else {
			moveA = headA
		}

		if moveB != nil {
			moveB = moveB.Next
		} else {
			moveB = headB
		}
	}
	return moveA
}
