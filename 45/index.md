## 45. 跳跃游戏 II

给你一个非负整数数组 `nums`，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。

### Think

题目给出**总是**可以到达最后一个位置，可以从最后一个位置逆推。

### Solution

```go
func jump(nums []int) int {
	position := len(nums) - 1 // 要达到的位置
	steps := 0                // 要用的步数
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
		end    int // 上一步能跳到的位置, 到达 end 就要多跳一步
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
```

### Tests

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例 1",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: 2,
		},
		{
			name: "示例 2",
			args: args{
				nums: []int{2, 3, 0, 1, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, jump(tt.args.nums))
		})
	}
}

var result int

func Benchmark_jump(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = jump([]int{2, 3, 1, 2, 4, 2, 3, 1, 1, 1})
	}
	result = res
}

func Benchmark_jump_V2(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = jump_V2([]int{2, 3, 1, 2, 4, 2, 3, 1, 1, 1})
	}
	result = res
}
```
