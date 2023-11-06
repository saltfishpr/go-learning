package internal

import (
	"io"
	"strings"
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

const testTimeZoneFile = `Europe/Guernsey,GG,BST,16701210000,3600,1
Europe/Guernsey,GG,GMT,16719354000,0,0
Europe/Helsinki,FI,EEST,16701210000,10800,1
Europe/Helsinki,FI,EET,16719354000,7200,0`

func Test_parseTimeZones(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name      string
		args      args
		want      map[string]*loaderZone
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "1",
			args: args{
				reader: strings.NewReader(testTimeZoneFile),
			},
			want: map[string]*loaderZone{
				"Europe/Guernsey": {
					Name: "Europe/Guernsey",
					Trans: []*loaderZoneTran{
						{
							Abbreviation: "BST",
							StartTime:    16701210000,
							EndTime:      newInt64(16719354000),
							Offset:       3600,
							IsDST:        true,
						},
						{
							Abbreviation: "GMT",
							StartTime:    16719354000,
							EndTime:      nil,
							Offset:       0,
							IsDST:        false,
						},
					},
				},
				"Europe/Helsinki": {
					Name: "Europe/Helsinki",
					Trans: []*loaderZoneTran{
						{
							Abbreviation: "EEST",
							StartTime:    16701210000,
							EndTime:      newInt64(16719354000),
							Offset:       10800,
							IsDST:        true,
						},
						{
							Abbreviation: "EET",
							StartTime:    16719354000,
							EndTime:      nil,
							Offset:       7200,
							IsDST:        false,
						},
					},
				},
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTimeZones(tt.args.reader)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func newInt64(i int64) *int64 {
	return &i
}
