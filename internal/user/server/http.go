package server

import (
	"context"
	"fmt"
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/samber/do"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	userv1 "github.com/saltfishpr/go-learning/gen/go/user/v1"
	"github.com/saltfishpr/go-learning/internal/user/conf"
)

func NewHTTP(i *do.Injector) *gwruntime.ServeMux {
	mux := gwruntime.NewServeMux(
		gwruntime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
			md := metadata.Pairs()
			for k, v := range r.Header {
				md.Append(k, v...)
			}
			return md
		}),
	)

	ctx := context.Background()
	config := do.MustInvoke[*conf.Config](i)
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	userv1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", config.Port), options)

	return mux
}
