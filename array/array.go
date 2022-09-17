// @file: array.go
// @date: 2021/2/8

// Package array
package array

import "strings"

func push(nums *[]byte, x byte) {
	*nums = append(*nums, x)
}

func pop(nums *[]byte) {
	n := len(*nums)
	if n > 0 {
		*nums = (*nums)[:n-1]
	}
}

func top(nums []byte) byte {
	n := len(nums)
	if n == 0 {
		return 0
	}
	return nums[n-1]
}

// 316. 去除重复字母
func removeDuplicateLetters(s string) string {
	count := [256]byte{}
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	inStack := [256]bool{}
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		c := s[i]
		count[c]--
		if inStack[c] {
			continue
		}
		for len(stack) > 0 && top(stack) > c && count[top(stack)] > 0 {
			inStack[top(stack)] = false
			pop(&stack)
		}
		push(&stack, c)
		inStack[c] = true
	}
	builder := new(strings.Builder)
	for i := 0; i < len(stack); i++ {
		builder.WriteByte(stack[i])
	}
	return builder.String()
}

// 26. 删除排序数组中的重复项
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	slow, fast := 0, 1
	for fast < n {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 83. 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil {
		if fast.Val != slow.Val {
			slow = slow.Next
			slow.Val = fast.Val
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}

// 27. 移除元素
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// 283. 移动零
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
