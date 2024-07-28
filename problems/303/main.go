package main

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	data := make([]int, len(nums))
	data[0] = nums[0]
	for i := 1; i < len(data); i++ {
		data[i] = data[i-1] + nums[i]
	}
	return NumArray{
		data: data,
	}
}

func (na *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return na.data[right]
	}
	return na.data[right] - na.data[left-1]
}
