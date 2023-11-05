package internal

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// echo logger middleware
func Logger(logger *slog.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogError:   true,
		LogLatency: true,
		BeforeNextFunc: func(c echo.Context) {
			ctx := c.Request().Context()
			ctx = LoggerToContext(ctx, logger)
			c.SetRequest(c.Request().WithContext(ctx))
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			ctx := c.Request().Context()
			requestGroup := slog.Group("request", slog.String("uri", v.URI), slog.Int("status", v.Status), slog.Duration("status", v.Latency))
			if v.Error == nil {
				LoggerFromContext(ctx).InfoContext(ctx, "success", requestGroup)
			} else {
				LoggerFromContext(ctx).ErrorContext(ctx, "failure", requestGroup, slog.Any("error", v.Error))
			}
			return nil
		},
	})
}
