package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_construct(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "示例 1",
			args: args{grid: [][]int{
				{0, 1},
				{1, 0},
			}},
			want: &Node{
				Val:         true,
				IsLeaf:      false,
				TopLeft:     &Node{Val: false, IsLeaf: true},
				TopRight:    &Node{Val: true, IsLeaf: true},
				BottomLeft:  &Node{Val: true, IsLeaf: true},
				BottomRight: &Node{Val: false, IsLeaf: true},
			},
		},
		{
			name: "",
			args: args{grid: [][]int{
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
			}},
			want: &Node{
				Val: true, IsLeaf: false,
				TopLeft: &Node{Val: true, IsLeaf: true},
				TopRight: &Node{
					Val: true, IsLeaf: false,
					TopLeft:     &Node{Val: false, IsLeaf: true},
					TopRight:    &Node{Val: false, IsLeaf: true},
					BottomLeft:  &Node{Val: true, IsLeaf: true},
					BottomRight: &Node{Val: true, IsLeaf: true},
				},
				BottomLeft:  &Node{Val: true, IsLeaf: true},
				BottomRight: &Node{Val: false, IsLeaf: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, construct(tt.args.grid))
		})
	}
}

var result *Node

func Benchmark_construct(b *testing.B) {
	grid := [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	var res *Node
	for n := 0; n < b.N; n++ {
		res = construct(grid)
	}
	result = res
}
