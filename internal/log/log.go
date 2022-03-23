// @file: log.go
// @date: 2022/3/23

// Package log 提供日志记录功能，内核由 zap 提供，适配 fiber, gorm 框架.
package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	sugar *zap.SugaredLogger
}

func New() *Logger {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &Logger{
		sugar: logger.Sugar(),
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.sugar.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.sugar.Debugf(template, args...)
}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.sugar.Debugw(msg, keysAndValues...)
}

func (l *Logger) Info(args ...interface{}) {
	l.sugar.Info(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.sugar.Infof(format, args...)
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.sugar.Infow(msg, keysAndValues...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.sugar.Warn(args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.sugar.Warnf(format, args...)
}

func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.sugar.Warnw(msg, keysAndValues...)
}

func (l *Logger) Error(args ...interface{}) {
	l.sugar.Error(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.sugar.Errorf(format, args...)
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.sugar.Errorw(msg, keysAndValues...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.sugar.Fatal(args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.sugar.Fatalf(format, args...)
}

func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.sugar.Fatalw(msg, keysAndValues...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	std := zap.NewStdLog(l.sugar.Desugar())
	std.Printf(format, args...)
}

func (l *Logger) Sync() error {
	return l.sugar.Sync()
}
