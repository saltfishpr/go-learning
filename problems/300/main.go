package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	var res int
	for _, v := range dp {
		if v > res {
			res = v
		}
	}
	return res
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
