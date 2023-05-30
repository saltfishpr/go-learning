package logger

import (
	"fmt"

	"github.com/SkyAPM/go2sky/logger"
	"go.uber.org/zap"
)

type Go2Sky struct {
	logger *zap.Logger
}

var _ logger.Log = (*Go2Sky)(nil)

func NewGo2Sky(logger *zap.Logger) *Go2Sky {
	return &Go2Sky{
		logger: logger,
	}
}

func (l *Go2Sky) Info(args ...interface{}) {
	l.logger.Info(l.getMessage("", args))
}

func (l *Go2Sky) Warn(args ...interface{}) {
	l.logger.Warn(l.getMessage("", args))
}

func (l *Go2Sky) Error(args ...interface{}) {
	l.logger.Error(l.getMessage("", args))
}

func (l *Go2Sky) Infof(format string, args ...interface{}) {
	l.logger.Info(l.getMessage(format, args))
}

func (l *Go2Sky) Warnf(format string, args ...interface{}) {
	l.logger.Warn(l.getMessage(format, args))
}

func (l *Go2Sky) Errorf(format string, args ...interface{}) {
	l.logger.Error(l.getMessage(format, args))
}

// getMessage format with Sprint, Sprintf, or neither.
//
// Ref: zap.SugaredLogger
func (l *Go2Sky) getMessage(template string, fmtArgs []interface{}) string {
	if len(fmtArgs) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, fmtArgs...)
	}

	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}
