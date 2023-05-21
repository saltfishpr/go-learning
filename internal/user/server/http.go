package server

import (
	"context"
	"fmt"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/samber/do"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userv1 "github.com/saltfishpr/go-learning/gen/go/user/v1"
	"github.com/saltfishpr/go-learning/internal/user/conf"
)

func NewHTTP(i *do.Injector) *gwruntime.ServeMux {
	var opts []gwruntime.ServeMuxOption
	opts = append(opts, gwruntime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
		return gwruntime.DefaultHeaderMatcher(s)
	}))

	mux := gwruntime.NewServeMux(opts...)
	ctx := context.Background()
	endpoint := fmt.Sprintf("localhost:%d", do.MustInvoke[*conf.Config](i).Port)
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	_ = userv1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, endpoint, options)

	return mux
}
