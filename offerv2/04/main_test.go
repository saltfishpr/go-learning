package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findNumberIn2DArray(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7, 9},
					{2, 4, 6, 8, 10},
					{11, 13, 15, 17, 19},
					{12, 14, 16, 18, 20},
					{21, 22, 23, 24, 25},
				},
				target: 13,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findNumberIn2DArray(tt.args.matrix, tt.args.target))
		})
	}
}
