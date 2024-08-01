package main

func countTarget(scores []int, target int) int {
	l := leftBound(scores, target)
	r := rightBound(scores, target)
	if l == -1 {
		return 0
	}
	return r - l + 1
}

func leftBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			r = mid - 1
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

func rightBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
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
