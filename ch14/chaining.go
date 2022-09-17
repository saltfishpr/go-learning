// @file: chaining.go
// @description:
// @author: SaltFish
// @date: 2020/09/08

// Package ch14 is chapter 14
package ch14

import (
	"flag"
	"fmt"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines") // 使用flag包解析命令行参数

func f(left, right chan int) { left <- 1 + <-right }

func MyChain() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0      // bang!
	x := <-leftmost // wait for completion
	fmt.Println(x)  // 100000, ongeveer 1.5 s
}

/*
无缓存信道具有同步阻塞的特性
1.主线程的right <- 0，right不是最初循环的那个right，而是最终循环的right
2.for循环中最初的go f(left, right)因为没有发送者一直处于等待状态
3.当主线程的right <- 0执行时，类似于递归函数在最内层产生返回值一般
*/
