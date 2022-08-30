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
