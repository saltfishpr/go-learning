package user

import (
	"context"
	"fmt"
	"time"

	configloader "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/samber/do"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"

	"github.com/saltfishpr/go-learning/internal/user/biz"
	"github.com/saltfishpr/go-learning/internal/user/conf"
	"github.com/saltfishpr/go-learning/internal/user/service"
)

func NewInjector(ctx context.Context, cfgFile string) (*do.Injector, error) {
	injector := do.New()

	cl := configloader.New(
		configloader.WithSource(
			file.NewSource(cfgFile),
		),
	)
	if err := cl.Load(); err != nil {
		return nil, err
	}
	do.ProvideValue(injector, cl)
	config := conf.New()
	if err := cl.Scan(config); err != nil {
		return nil, err
	}
	do.ProvideValue(injector, config)

	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.StacktraceKey = ""
	if config.LogLevel == "debug" {
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	} else {
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
	logger, err := zapConfig.Build(zap.AddCaller())
	if err != nil {
		return nil, err
	}
	do.ProvideValue(injector, logger)

	res, err := sdkresource.New(ctx,
		sdkresource.WithAttributes(
			semconv.ServiceName(config.OTEL.ServiceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	metricExp, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(config.OTEL.CollectorAddr))
	if err != nil {
		return nil, fmt.Errorf("failed to create the collector metric exporter: %w", err)
	}
	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				metricExp,
				sdkmetric.WithInterval(2*time.Second),
			),
		),
	)
	otel.SetMeterProvider(meterProvider)

	traceExp, err := otlptrace.New(ctx, otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(config.OTEL.CollectorAddr),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
	))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}
	// 使用批处理跨度处理器在导出前聚合跨度
	bsp := sdktrace.NewBatchSpanProcessor(traceExp)
	// 使用 TracerProvider 注册跟踪导出器
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	do.Provide(injector, biz.NewUserUseCase)

	do.Provide(injector, service.NewUserService)

	return injector, nil
}
