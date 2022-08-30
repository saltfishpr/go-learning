package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findCheapestPrice(t *testing.T) {
	type args struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例 1",
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
				src:     0,
				dst:     2,
				k:       1,
			},
			want: 200,
		},
		{
			name: "示例 2",
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
				src:     0,
				dst:     2,
				k:       0,
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findCheapestPrice(tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.k))
		})
	}
}

var result int

func Benchmark_findCheapestPrice(b *testing.B) {
	var res int
	for n := 0; n < b.N; n++ {
		res = findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1)
	}
	result = res
}
