// @description:
// @file: log_test.go
// @date: 2022/3/24

package log

import (
	"testing"
)

func TestLogger_Debug(t *testing.T) {
	l := New()
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "debug",
			args: args{
				args: []interface{}{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Debug(tt.args.args...)
		})
	}
}

func TestLogger_Debugf(t *testing.T) {
	l := New()
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "debugf",
			args: args{
				template: "%s",
				args:     []interface{}{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Debugf(tt.args.template, tt.args.args...)
		})
	}
}

func TestLogger_Debugw(t *testing.T) {
	l := New()
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "debugw",
			args: args{
				msg:           "test",
				keysAndValues: []interface{}{"key", "value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Debugw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestLogger_Info(t *testing.T) {
	l := New()
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "info",
			args: args{
				args: []interface{}{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Info(tt.args.args...)
		})
	}
}

func TestLogger_Infof(t *testing.T) {
	l := New()
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "infof",
			args: args{
				template: "%s",
				args:     []interface{}{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Infof(tt.args.template, tt.args.args...)
		})
	}
}

func TestLogger_Infow(t *testing.T) {
	l := New()
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "infow",
			args: args{
				msg:           "test",
				keysAndValues: []interface{}{"key", "value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Infow(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestLogger_Warn(t *testing.T) {
	l := New()
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "warn",
			args: args{
				args: []interface{}{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Warn(tt.args.args...)
		})
	}
}

func TestLogger_Warnf(t *testing.T) {
	l := New()
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "warnf",
			args: args{
				template: "%s",
				args:     []interface{}{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Warnf(tt.args.template, tt.args.args...)
		})
	}
}

func TestLogger_Warnw(t *testing.T) {
	l := New()
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "warnw",
			args: args{
				msg:           "test",
				keysAndValues: []interface{}{"key", "value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Warnw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestLogger_Error(t *testing.T) {
	l := New()
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "error",
			args: args{
				args: []interface{}{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Error(tt.args.args...)
		})
	}
}

func TestLogger_Errorf(t *testing.T) {
	l := New()
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "errorf",
			args: args{
				template: "%s",
				args:     []interface{}{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Errorf(tt.args.template, tt.args.args...)
		})
	}
}

func TestLogger_Errorw(t *testing.T) {
	l := New()
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "errorw",
			args: args{
				msg:           "test",
				keysAndValues: []interface{}{"key", "value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Errorw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestLogger_Fatal(t *testing.T) {
	l := New()
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "fatal",
			args: args{
				args: []interface{}{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Fatal(tt.args.args...)
		})
	}
}

func TestLogger_Fatalf(t *testing.T) {
	l := New()
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "fatalf",
			args: args{
				template: "%s",
				args:     []interface{}{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Fatalf(tt.args.template, tt.args.args...)
		})
	}
}

func TestLogger_Fatalw(t *testing.T) {
	l := New()
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "fatalw",
			args: args{
				msg:           "test",
				keysAndValues: []interface{}{"key", "value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Fatalw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestLogger_Printf(t *testing.T) {
	l := New()
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "printf",
			args: args{
				format: "%s",
				args:   []interface{}{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Printf(tt.args.format, tt.args.args...)
		})
	}
}
