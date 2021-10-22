// @file: init.go
// @description:
// @author: SaltFish
// @date: 2020/07/31

// Package trans is a submod of user_init
package trans

import "math"

// Pi is 圆周率
var Pi float64

func init() {
	Pi = 4 * math.Atan(1) // init() function computes Pi
}
