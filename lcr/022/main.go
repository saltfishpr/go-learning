package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表长度为 a + b，成环的长度为 b
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	// fast 一次走两步，slow 一次走一步，直到第一次相遇
	// 此时 slow 走 nb 步，fast 走 2nb 步
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			break
		}
	}

	// 需要让 slow 再走 a 步，就可以找到环的“入口”
	// 此时让 fast 重新指向 head
	// fast 和 slow 一起走，相遇节点则为环的“入口”
	fast = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
