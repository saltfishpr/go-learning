// @file: 14-1剪绳子.go
// @date: 2021/2/13

// Package offer
package offer

/*
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

3*3 > 2*4
2*2 > 3*1
*/

func cuttingRopeX14(n int) int {
	pow := func(x, n int) int {
		if x == 0 {
			return 0
		}
		res := 1
		for n != 0 {
			// n是奇数
			if n&1 == 1 {
				res *= x
			}
			x *= x
			n >>= 1
		}
		return res
	}

	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	quotient := n / 3
	remainder := n % 3
	if remainder == 0 {
		return pow(3, quotient)
	} else if remainder == 1 {
		return pow(3, quotient-1) * 4
	} else {
		return pow(3, quotient) * 2
	}
}
