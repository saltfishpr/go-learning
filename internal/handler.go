package internal

import (
	"context"
	"net/http"

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
	// ListTimeZones .
	ListTimeZones(ctx context.Context) ([]*TimeZone, error)
	// ListTimeZonesByTimestamp .
	ListTimeZonesByTimestamp(ctx context.Context, sec int64) ([]*TimeZone, error)
	// Count .
	Count(ctx context.Context) (int64, error)
	// ReplaceAllBySource .
	ReplaceAllBySource(ctx context.Context, tzs []*TimeZone, source string) error
}

type Handler struct {
	repo     Repo
	CacheDir string
}

func NewHandler(i *do.Injector) (*Handler, error) {
	config := do.MustInvoke[*Config](i)
	return &Handler{
		repo:     do.MustInvoke[Repo](i),
		CacheDir: config.CacheDir,
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
	loader := NewLoader(
		h.CacheDir,
		"https://timezonedb.com",
		"TimeZoneDB.csv.zip",
	)
	tzs, err := loader.Load(false)
	if err != nil {
		return echo.ErrInternalServerError.WithInternal(err)
	}
	for _, tz := range tzs {
		tz.ID = ulid.Make().String()
	}

	if err := h.repo.ReplaceAllBySource(c.Request().Context(), tzs, "timezonedb.com"); err != nil {
		return echo.ErrInternalServerError.WithInternal(err)
	}
	return c.NoContent(http.StatusOK)
}
