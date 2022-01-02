// @description: 封装 zap 库，提供默认实现。
// @file: logger.go
// @date: 2021/11/16

// Package logger 提供日志记录。
package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	maxSize    = 10
	maxAge     = 30
	maxBackups = 5
)

// 初始化前使用默认配置。
var sugar = zap.NewExample(zap.Development()).Sugar()

// Init 初始化 sugar。
func Init(release bool) {
	sugar = New(release)
}

// New 使用配置创建 *zap.SugaredLogger。
func New(release bool) *zap.SugaredLogger {
	var (
		encoder zapcore.Encoder
		level   zapcore.Level
	)

	if release {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
		level = zapcore.WarnLevel
	} else {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
		level = zapcore.DebugLevel
	}

	writeSyncer := getLogWriter(fmt.Sprintf("logs/%s.log", time.Now().Format("D2006-01-02T15-04-05")))
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer, os.Stdout), level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger.Sugar()
}

// getLogWriter 将日志写入文件
func getLogWriter(filename string) zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		Compress:   false,
	}
	return zapcore.AddSync(hook)
}

func Debug(args ...interface{}) {
	sugar.Debug(args...)
}

func Info(args ...interface{}) {
	sugar.Info(args...)
}

func Warn(args ...interface{}) {
	sugar.Warn(args...)
}

func Error(args ...interface{}) {
	sugar.Error(args...)
}

func Panic(args ...interface{}) {
	sugar.Panic(args...)
}

func Fatal(args ...interface{}) {
	sugar.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	sugar.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	sugar.Fatalf(template, args...)
}

func Logger() *zap.SugaredLogger {
	return sugar
}

func Sync() {
	_ = sugar.Sync()
}
