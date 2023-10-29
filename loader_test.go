package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getTimeZoneOffsetString(t *testing.T) {
	type args struct {
		offset int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				offset: -25200,
			},
			want: "-07:00",
		},
		{
			name: "2",
			args: args{
				offset: 28800,
			},
			want: "+08:00",
		},
		{
			name: "3",
			args: args{
				offset: 0,
			},
			want: "+00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getTimeZoneOffsetString(tt.args.offset))
		})
	}
}
