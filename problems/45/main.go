package main

func jump(nums []int) int {
	position := len(nums) - 1
	steps := 0
	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] >= position {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}

func jump_V2(nums []int) int {
	var (
		maxPos int // 这一步能跳到的最远的位置
		end    int // 上一步能跳到的位置，到达 end 就要多跳一步
		step   int // 跳了几步
	)

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > maxPos {
			maxPos = i + nums[i]
		}
		if i == end {
			step++
			end = maxPos
		}
	}

	return step
}
