package internal

import (
	"context"
	"log/slog"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/samber/do"
	"github.com/spf13/cast"
)

//go:generate mockgen -typed -destination=handler_mock_test.go -package=internal . Repo

type Repo interface {
	// CreateTimeZone .
	CreateTimeZone(ctx context.Context, tz *TimeZone) (*TimeZone, error)
	// UpdateTimeZone .
	UpdateTimeZone(ctx context.Context, tz *TimeZone, fields []string) (*TimeZone, error)
	// DeleteTimeZone .
	DeleteTimeZone(ctx context.Context, id string) error
	// GetTimeZone .
	GetTimeZone(ctx context.Context, id string) (*TimeZone, error)
	// ListTimeZones .
	ListTimeZones(ctx context.Context, offset int, limit int) ([]*TimeZone, error)
	// ListTimeZonesByTimestamp .
	ListTimeZonesByTimestamp(ctx context.Context, sec int64) ([]*TimeZone, error)
	// Count .
	Count(ctx context.Context) (int64, error)
	// ReplaceAllBySource .
	ReplaceAllBySource(ctx context.Context, tzs []*TimeZone, source string) error
}

type LoadTask struct {
	ID string `json:"id"`

	cancel func() `json:"-"`
}

type Handler struct {
	repo     Repo
	cacheDir string

	taskMu   sync.Mutex
	loadTask *LoadTask // LoadTask
}

func NewHandler(i *do.Injector) (*Handler, error) {
	config := do.MustInvoke[*Config](i)
	return &Handler{
		repo:     do.MustInvoke[Repo](i),
		cacheDir: config.CacheDir,
	}, nil
}

func (h *Handler) CreateTimeZone(c echo.Context) error {
	tz := new(TimeZone)
	if err := c.Bind(tz); err != nil {
		return echo.ErrBadRequest.WithInternal(err)
	}
	tz.ID = ulid.Make().String()
	if tz.EndTime == "" {
		tz.EndTime = MaxTimeString
	}

	tz, err := h.repo.CreateTimeZone(c.Request().Context(), tz)
	if err != nil {
		return echo.ErrInternalServerError.WithInternal(err)
	}
	return c.JSON(http.StatusOK, tz)
}

func (h *Handler) UpdateTimeZone(c echo.Context) error {
	type request struct {
		ID string `param:"id"`

		TimeZone *TimeZone `json:"timeZone"`
		Fields   []string  `json:"fields"`
	}
	req := new(request)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest.WithInternal(err)
	}
	req.TimeZone.ID = req.ID

	tz, err := h.repo.UpdateTimeZone(c.Request().Context(), req.TimeZone, req.Fields)
	if err != nil {
		return echo.ErrInternalServerError.WithInternal(err)
	}
	return c.JSON(http.StatusOK, tz)
}

func (h *Handler) DeleteTimeZone(c echo.Context) error {
	type request struct {
		ID string `param:"id"`
	}
	req := new(request)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest.WithInternal(err)
	}

	if err := h.repo.DeleteTimeZone(c.Request().Context(), req.ID); err != nil {
		return echo.ErrInternalServerError.WithInternal(err)
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) GetTimeZone(c echo.Context) error {
	type request struct {
		ID string `param:"id"`
	}
	req := new(request)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest.WithInternal(err)
	}

	tz, err := h.repo.GetTimeZone(c.Request().Context(), req.ID)
	if err != nil {
		return echo.ErrInternalServerError.WithInternal(err)
	}
	return c.JSON(http.StatusOK, tz)
}

func (h *Handler) ListTimeZones(c echo.Context) error {
	type request struct {
		Timestamp string `query:"t"`
	}
	req := new(request)
	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest.WithInternal(err)
	}

	count, err := h.repo.Count(c.Request().Context())
	if err != nil {
		return err
	}

	if req.Timestamp == "" {
		return echo.ErrBadRequest
	}

	sec, err := cast.ToInt64E(req.Timestamp)
	if err != nil {
		return echo.ErrInternalServerError.WithInternal(err)
	}
	timeZones, err := h.repo.ListTimeZonesByTimestamp(c.Request().Context(), sec)
	if err != nil {
		return echo.ErrInternalServerError.WithInternal(err)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data":      timeZones,
		"totalSize": count,
	})
}

func (h *Handler) LoadTimeZones(c echo.Context) error {
	if h.taskMu.TryLock() && h.loadTask != nil {
		defer h.taskMu.Unlock()
		return c.JSON(http.StatusOK, h.loadTask)
	}

	loader := NewLoader(
		h.cacheDir,
		"https://timezonedb.com",
		"TimeZoneDB.csv.zip",
	)

	ctx := context.Background()
	ctx = LoggerToContext(ctx, LoggerFromContext(c.Request().Context()))
	ctx, cancel := context.WithCancel(ctx)

	task := &LoadTask{
		ID:     ulid.Make().String(),
		cancel: cancel,
	}
	h.loadTask = task
	defer h.taskMu.Unlock()

	go func() {
		defer func() {
			h.taskMu.Lock()
			defer h.taskMu.Unlock()
			h.loadTask = nil
		}()

		tzs, err := loader.Load(ctx, false)
		if err != nil {
			LoggerFromContext(ctx).Error("load data error", slog.Any("error", err))
			return
		}
		for _, tz := range tzs {
			tz.ID = ulid.Make().String()
		}

		if err := h.repo.ReplaceAllBySource(ctx, tzs, "timezonedb.com"); err != nil {
			LoggerFromContext(ctx).Error("save to db error", slog.Any("error", err))
			return
		}

		LoggerFromContext(ctx).Info("successfully saved timezones")
	}()

	return c.JSON(http.StatusOK, task)
}

func (h *Handler) CancelLoadTask(c echo.Context) error {
	return nil
}
