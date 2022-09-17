// @file: link_list.go
// @date: 2021/1/22

// Package linklist
package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 83. 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil {
		if slow.Val == fast.Val {
			slow.Next = fast.Next
		} else {
			slow = fast
		}
		fast = fast.Next
	}
	return head
}

// 82. 删除排序链表中的重复元素 II
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := &ListNode{Val: 0, Next: head}
	head = res
	var rmVal int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmVal = head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return res.Next
}

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	res := &ListNode{}
	for head != nil {
		res.Next, head, head.Next = head, head.Next, res.Next
	}
	return res.Next
}

// 92. 反转链表 II
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	res := &ListNode{Next: head}
	head = res
	var pre *ListNode
	var i = 0
	for ; i < m; i++ {
		pre = head
		head = head.Next
	}

	mid := head
	for ; i <= n && head != nil; i++ {
		pre.Next, head.Next, head = head, pre.Next, head.Next
	}
	mid.Next = head
	return res.Next
}

func reverseBetweenV2(head *ListNode, m int, n int) *ListNode {
	var successor *ListNode = nil

	var reserveN func(*ListNode, int) *ListNode
	reserveN = func(head *ListNode, n int) *ListNode {
		if n == 1 {
			successor = head.Next
			return head
		}
		last := reserveN(head.Next, n-1)
		head.Next.Next = head
		head.Next = successor
		return last
	}

	if m == 1 {
		return reserveN(head, n)
	}
	head.Next = reverseBetweenV2(head.Next, m-1, n-1)
	return head
}

// 21. 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	move := res
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			move.Next, l1, l1.Next = l1, l1.Next, move.Next
		} else {
			move.Next, l2, l2.Next = l2, l2.Next, move.Next
		}
		move = move.Next
	}
	if l1 != nil {
		move.Next = l1
	} else {
		move.Next = l2
	}
	return res.Next
}

// 25. K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 反转[a, b)的节点
	reverseAToB := func(a, b *ListNode) *ListNode {
		// pre  ->  cur  ->  nxt
		var pre, cur, nxt *ListNode = nil, a, a
		for cur != b {
			nxt = cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		return pre
	}

	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverseAToB(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

// 86. 分隔链表
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	smallDummy := &ListNode{}
	small := smallDummy
	bigDummy := &ListNode{}
	big := bigDummy
	for head != nil {
		if head.Val < x {
			small.Next, head = head, head.Next
			small = small.Next
		} else {
			big.Next, head = head, head.Next
			big = big.Next
		}
	}
	big.Next = nil
	small.Next = bigDummy.Next
	return smallDummy.Next
}

// 148. 排序链表
func sortList(head *ListNode) *ListNode {
	var mergeSort func(*ListNode) *ListNode
	mergeSort = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		mid := findMid(head)
		tail := mid.Next
		mid.Next = nil
		left := mergeSort(head)
		right := mergeSort(tail)
		result := mergeTwoLists(left, right)
		return result
	}

	return mergeSort(head)
}

// 143. 重排链表
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := findMid(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)

	mergeTwoLists := func(l1, l2 *ListNode) *ListNode {
		dummy := &ListNode{Val: 0}
		head := dummy
		toggle := true
		for l1 != nil && l2 != nil {
			if toggle {
				head.Next, l1 = l1, l1.Next
			} else {
				head.Next, l2 = l2, l2.Next
			}
			head = head.Next
			toggle = !toggle
		}
		if l1 != nil {
			head.Next = l1
		} else {
			head.Next = l2
		}
		return dummy.Next
	}

	head = mergeTwoLists(l1, l2)
}

// 141. 环形链表
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

// 142. 环形链表 II
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			fast = head
			slow = slow.Next
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}

// 234. 回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	mid := findMid(head)
	tail := reverseList(mid.Next)
	mid.Next = nil
	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head, tail = head.Next, tail.Next
	}
	// 这里不用判断链表是否走完，节点个数为奇数时 head.next!=nil
	return true
}

func isPalindromeV2(head *ListNode) bool {
	left := head

	var traverse func(*ListNode) bool
	traverse = func(right *ListNode) bool {
		if right == nil {
			return true
		}
		res := traverse(right.Next)
		res = res && right.Val == left.Val
		left = left.Next
		return res
	}

	return traverse(head)
}

// 138. 复制带随机指针的链表
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		clone := &Node{
			Val:    cur.Val,
			Next:   cur.Next,
			Random: nil,
		}
		cur.Next, cur = clone, cur.Next
	}
	// 处理random
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	cur = head
	cloneHead := cur.Next
	for cur != nil && cur.Next != nil {
		cur.Next, cur = cur.Next.Next, cur.Next
	}
	return cloneHead
}
