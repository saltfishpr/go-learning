package ctxutil

import (
	"context"

	jwtv4 "github.com/golang-jwt/jwt/v4"
)

var ctxKeyClaims int

func ClaimsToContext(ctx context.Context, claims jwtv4.MapClaims) context.Context {
	return context.WithValue(ctx, &ctxKeyClaims, claims)
}

func ClaimsFromContext(ctx context.Context) jwtv4.MapClaims {
	claims, ok := ctx.Value(&ctxKeyClaims).(jwtv4.MapClaims)
	if !ok {
		return nil
	}
	return claims
}
