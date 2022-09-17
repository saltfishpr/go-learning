// @file: 59-2队列的最大值.go
// @date: 2021/2/24

// Package offer59
package offer59

/*
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value需要返回 -1
*/

type MaxQueue struct {
	queue, max []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue: make([]int, 0),
		max:   make([]int, 0),
	}
}

func (mq *MaxQueue) Max_value() int {
	if len(mq.max) == 0 {
		return -1
	}
	return mq.max[0]
}

func (mq *MaxQueue) Push_back(value int) {
	for len(mq.max) != 0 && mq.max[len(mq.max)-1] < value {
		mq.max = mq.max[:len(mq.max)-1]
	}
	mq.queue = append(mq.queue, value)
	mq.max = append(mq.max, value)
}

func (mq *MaxQueue) Pop_front() int {
	if len(mq.queue) == 0 {
		return -1
	}
	value := mq.queue[0]
	if mq.max[0] == value {
		mq.max = mq.max[1:]
	}
	mq.queue = mq.queue[1:]
	return value
}
