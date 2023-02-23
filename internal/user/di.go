package user

import (
	"io"
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/do"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/saltfishpr/go-learning/internal/user/biz"
	"github.com/saltfishpr/go-learning/internal/user/conf"
	"github.com/saltfishpr/go-learning/internal/user/service"
)

func init() {
	// disable kratos log
	log.SetLogger(log.NewStdLogger(io.Discard))
}

func NewInjector(cfgFile string) (*do.Injector, error) {
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
	if os.Getenv("LOG_LEVEL") == "debug" {
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	} else {
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
	logger, err := zapConfig.Build(zap.AddCaller())
	if err != nil {
		return nil, err
	}
	do.ProvideValue(injector, logger)

	do.Provide(injector, biz.NewUserUseCase)

	do.Provide(injector, service.NewUserService)

	return injector, nil
}
