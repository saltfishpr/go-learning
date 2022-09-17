// @file: stack.go
// @date: 2021/2/7

// Package monotonicstack
package monotonicstack

func push(nums *[]int, x int) {
	*nums = append(*nums, x)
}

func pop(nums *[]int) {
	n := len(*nums)
	if n > 0 {
		*nums = (*nums)[:n-1]
	}
}

func top(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	return nums[n-1]
}

// 496. 下一个更大元素 I
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	bigger := make(map[int]int, 0)
	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) != 0 && top(stack) <= nums2[i] {
			pop(&stack)
		}
		if len(stack) == 0 {
			bigger[nums2[i]] = -1
		} else {
			bigger[nums2[i]] = top(stack)
		}
		push(&stack, nums2[i])
	}
	for i := range nums1 {
		nums1[i] = bigger[nums1[i]]
	}
	return nums1
}

// 503. 下一个更大元素 II
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	stack := make([]int, 0)
	// 循环数组，可以将数组长度翻倍 使用取模获得翻倍的效果
	for i := 2*n - 1; i >= 0; i-- {
		num := nums[i%n]
		for len(stack) != 0 && top(stack) <= num {
			pop(&stack)
		}
		if len(stack) == 0 {
			res[i%n] = -1
		} else {
			res[i%n] = top(stack)
		}
		push(&stack, num)
	}
	return res
}

// 739. 每日温度
func dailyTemperatures(T []int) []int {
	n := len(T)
	res := make([]int, n)
	stack := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		num := T[i]
		for len(stack) != 0 && T[top(stack)] <= num {
			pop(&stack)
		}
		if len(stack) == 0 {
			res[i] = 0
		} else {
			res[i] = top(stack) - i
		}
		push(&stack, i)
	}
	return res
}
