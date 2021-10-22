// @file: 57-2和为s的连续正数序列.go
// @date: 2021/2/23

// Package offer
package offer

/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
*/

func findContinuousSequenceX57(target int) [][]int {
	res := make([][]int, 0)
	sum := 0
	i, j := 1, 1 // [i, j)
	for i <= target/2 {
		if sum < target {
			sum += j
			j++
		} else if sum > target {
			sum -= i
			i++
		} else {
			tmp := make([]int, 0)
			for k := i; k < j; k++ {
				tmp = append(tmp, k)
			}
			res = append(res, tmp)
			sum -= i
			i++
		}
	}
	return res
}
