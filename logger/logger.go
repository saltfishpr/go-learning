// @file: logger.go
// @date: 2021/10/22

// Package logger
package logger

import (
	"fmt"
	"os"
	"time"

	"learning/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugar = New(config.Release)

func New(release string) *zap.SugaredLogger {
	var (
		encoder zapcore.Encoder
		level   zapcore.Level
	)

	if release == "true" {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
		level = zapcore.WarnLevel
	} else {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
		level = zapcore.DebugLevel
	}

	writeSyncer := getLogWriter(fmt.Sprintf("logs/%s.log", time.Now().Format("D2006-01-02T15-04-05")))
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer, os.Stdout), level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger.Sugar()
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 5,
		Compress:   false,
	}
	return zapcore.AddSync(hook)
}

func Debug(args ...interface{}) {
	sugar.Debug(args)
}

func Info(args ...interface{}) {
	sugar.Info(args)
}

func Warn(args ...interface{}) {
	sugar.Warn(args)
}

func Error(args ...interface{}) {
	sugar.Error(args)
}

func Panic(args ...interface{}) {
	sugar.Panic(args)
}

func Fatal(args ...interface{}) {
	sugar.Fatal(args)
}

func Debugf(template string, args ...interface{}) {
	sugar.Debugf(template, args)
}

func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args)
}

func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args)
}

func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args)
}

func Panicf(template string, args ...interface{}) {
	sugar.Panicf(template, args)
}

func Fatalf(template string, args ...interface{}) {
	sugar.Fatalf(template, args)
}

func Sync() error {
	return sugar.Sync()
}
