package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidSerialization(t *testing.T) {
	type args struct {
		preorder string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例 1",
			args: args{
				preorder: "9,3,4,#,#,1,#,#,2,#,6,#,#",
			},
			want: true,
		},
		{
			name: "示例 2",
			args: args{
				preorder: "1,#",
			},
			want: false,
		},
		{
			name: "示例 3",
			args: args{
				preorder: "9,#,#,1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isValidSerialization(tt.args.preorder))
		})
	}
}

var result bool

func Benchmark_isValidSerialization(b *testing.B) {
	var res bool
	for n := 0; n < b.N; n++ {
		res = isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#")
	}
	result = res
}
