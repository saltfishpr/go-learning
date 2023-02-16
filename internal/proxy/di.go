package proxy

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/samber/do"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/saltfishpr/go-learning/internal/proxy/conf"
)

func NewInjector(cfgFile string, release bool) (*do.Injector, error) {
	var injector *do.Injector

	c := config.New(
		config.WithSource(
			file.NewSource(cfgFile),
		),
	)
	if err := c.Load(); err != nil {
		return nil, err
	}
	do.ProvideValue(injector, c)

	cfg := new(conf.Config)
	if err := c.Scan(cfg); err != nil {
		return nil, err
	}
	do.ProvideValue(injector, cfg)

	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.StacktraceKey = ""
	if release {
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	} else {
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}
	logger, err := zapConfig.Build(zap.AddCaller())
	if err != nil {
		return nil, err
	}
	do.ProvideValue(injector, logger)

	return injector, nil
}
