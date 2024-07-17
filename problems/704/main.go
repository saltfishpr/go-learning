package main

func main() {
	search([]int{-1, 0, 3, 5, 9, 12}, 13)
}

func search(nums []int, target int) int {
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
