package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumOfDistancesInTree(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例 1",
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}},
			},
			want: []int{8, 12, 6, 10, 10, 10},
		},
		{
			name: "示例 2",
			args: args{
				n:     1,
				edges: [][]int{},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sumOfDistancesInTree(tt.args.n, tt.args.edges))
		})
	}
}

var result []int

func Benchmark_sumOfDistancesInTree(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = sumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}})
	}
	result = res
}
