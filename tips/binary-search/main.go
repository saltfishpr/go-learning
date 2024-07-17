package main

// binarySearch 在单调数组 nums 中搜索值为 target 索引。
func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return -1
}

// leftBound 在单调数组 nums 中搜索值为 target 的最小索引。
func leftBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] == target {
			r = mid - 1 //
		}
	}
	if l < 0 || l >= len(nums) {
		return -1
	}
	if nums[l] == target {
		return l
	}
	return -1
}

// rightBound 在单调数组 nums 中搜索值为 target 的最大索引。
func rightBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] == target {
			l = mid + 1
		}
	}
	if r < 0 || r >= len(nums) {
		return -1
	}
	if nums[r] == target {
		return r
	}
	return -1
}
