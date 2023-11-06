package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeZone_Contains(t *testing.T) {
	type fields struct {
		StartTime string
		EndTime   string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "1",
			fields: fields{
				StartTime: "2000-01-01 00:00:00",
				EndTime:   "2006-01-02 15:04:05",
			},
			args: args{
				s: "2000-01-01 10:00:00",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tz := TimeZone{
				StartTime: tt.fields.StartTime,
				EndTime:   tt.fields.EndTime,
			}
			assert.Equal(t, tt.want, tz.Contains(tt.args.s))
		})
	}
}
