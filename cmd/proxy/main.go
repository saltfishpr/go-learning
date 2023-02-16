package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/samber/do"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	"github.com/saltfishpr/go-learning/internal/proxy"
	"github.com/saltfishpr/go-learning/internal/proxy/conf"
)

var (
	cfgFile string
	release bool
)

func init() {
	flag.StringVar(&cfgFile, "config", "configs/config.yaml", "config file path")
	flag.BoolVar(&release, "release", false, "release mode")
}

func main() {
	flag.Parse()

	injector, err := proxy.NewInjector(cfgFile, release)
	if err != nil {
		panic(err)
	}

	logger := do.MustInvoke[*zap.Logger](injector)
	config := do.MustInvoke[*conf.Config](injector)

	mux, err := initServer(injector)
	if err != nil {
		logger.Fatal("init server failed", zap.Error(err))
	}

	if err := http.ListenAndServe(config.Addr, mux); err != nil {
		logger.Fatal("listen and serve failed", zap.Error(err))
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func initServer(i *do.Injector) (*gwruntime.ServeMux, error) {
	mux := gwruntime.NewServeMux(
		gwruntime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
			md := metadata.Pairs()
			for k, v := range r.Header {
				md.Append(k, v...)
			}
			return md
		}),
		gwruntime.WithErrorHandler(func(ctx context.Context, sm *gwruntime.ServeMux, m gwruntime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
			newErr := &gwruntime.HTTPStatusError{
				HTTPStatus: 200,
				Err:        err,
			}
			gwruntime.DefaultHTTPErrorHandler(ctx, sm, m, w, r, newErr)
		}),
	)

	return mux, nil
}
