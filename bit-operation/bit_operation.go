// @file: bit_operation.go
// @date: 2021/1/25

// Package bitoperation
package bitoperation

// 136. 只出现一次的数字
func singleNumber(nums []int) (res int) {
	for i := range nums {
		res ^= nums[i]
	}
	return
}

// 137. 只出现一次的数字 II
func singleNumber2(nums []int) int {
	var res int
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			// 统计1的个数
			sum += (nums[j] >> i) & 1
		}
		res |= (sum % 3) << i
	}
	return res
}

func singleNumber2Better(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = ones ^ num & ^twos
		twos = twos ^ num & ^ones
	}
	return ones
}

// 260. 只出现一次的数字 III
func singleNumber3(nums []int) []int {
	bitmask := 0
	for i := range nums {
		bitmask ^= nums[i]
	}
	// 获取bitmask为1的位，两个出现一次的数字在这一位上不同
	diff := (-bitmask) & bitmask
	x := 0
	for i := range nums {
		// 根据这个不同将nums分为两个部分
		if diff&nums[i] == 0 {
			x ^= nums[i]
		}
	}
	return []int{x, x ^ bitmask}
}

// 191. 位1的个数
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}

// 338. 比特位计数
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

// 190. 颠倒二进制位
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		// i 和 31-i 交换
		res |= (num & 1) << (31 - i)
		num >>= 1
	}
	return res
}

// 201. 数字范围按位与
func rangeBitwiseAnd(m int, n int) int {
	shift := 0
	for m != n {
		m >>= 1
		n >>= 1
		shift++
	}
	return m << shift
}
