// @file: 06从尾到头打印链表.go
// @date: 2021/2/12

// Package offer
package offer

/*
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
*/

func reversePrintX06(head *ListNode) []int {
	reverse := func(nums []int, i, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	reverse(res, 0, len(res)-1)
	return res
}
