package logger

import (
	"github.com/ThreeDotsLabs/watermill"
	"go.uber.org/zap"
)

type Watermill struct {
	logger *zap.Logger
}

var _ watermill.LoggerAdapter = (*Watermill)(nil)

func NewWatermill(logger *zap.Logger) *Watermill {
	return &Watermill{
		logger: logger,
	}
}

func (l *Watermill) Debug(msg string, fields watermill.LogFields) {
	l.logger.Debug(msg, l.convertFields(fields)...)
}

func (l *Watermill) Trace(msg string, fields watermill.LogFields) {
	convertFields := append(l.convertFields(fields), zap.Stack("stack"))
	l.logger.Debug(msg, convertFields...)
}

func (l *Watermill) Info(msg string, fields watermill.LogFields) {
	l.logger.Info(msg, l.convertFields(fields)...)
}

func (l *Watermill) Error(msg string, err error, fields watermill.LogFields) {
	l.logger.Error(msg, l.convertFields(fields)...)
}

func (l *Watermill) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return &Watermill{
		logger: l.logger.With(l.convertFields(fields)...),
	}
}

func (l *Watermill) convertFields(fields watermill.LogFields) []zap.Field {
	var zapFields []zap.Field
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return zapFields
}
