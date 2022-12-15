package main

import (
	"net/http"
	"time"

	xmiddleware "learning/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.StacktraceKey = ""
	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	logger, err := zapConfig.Build(zap.AddCaller())
	if err != nil {
		panic(err)
	}

	// Routes
	r := chi.NewRouter()
	r.Use(xmiddleware.RequestID)
	r.Use(
		middleware.RequestLogger(&xmiddleware.StructuredLogger{Logger: logger}),
	)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/wait", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		xmiddleware.LogEntrySetFields(r.Context(), zap.Bool("wait", true))
		w.Write([]byte("hi"))
	})
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic(errors.New("oops"))
	})

	http.ListenAndServe(":3333", r)
}
