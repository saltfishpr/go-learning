package main

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	heights := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		heights[i] = envelopes[i][1]
	}

	return lengthOfLISV2(heights)
}

func lengthOfLISV2(nums []int) int {
	var top []int // 存储牌堆顶部的牌

	for i := 0; i < len(nums); i++ {
		poker := nums[i]

		// 二分搜索左侧边界
		left, right := 0, len(top)-1
		for left <= right {
			mid := left + (right-left)/2
			if top[mid] > poker {
				right = mid - 1
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if left == len(top) {
			// left 是大于 poker 的最小索引
			top = append(top, poker)
		} else {
			top[left] = poker
		}
	}

	return len(top)
}
