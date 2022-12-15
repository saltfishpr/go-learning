package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const RequestIDHeader = "X-Request-Id"

// RequestID is a middleware that injects a request ID into the context of each
// request. If a request ID is already present in the request header, it will be
// used. Otherwise, a uuid v4 string will be generated and added to the request
// header and response header.
func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			requestID = uuid.NewString()
			r.Header.Set(RequestIDHeader, requestID)
		}
		w.Header().Set(RequestIDHeader, requestID)
		r = r.WithContext(WithRequestID(r.Context(), requestID))
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// ctxKeyRequestID is the key for the request ID in a context.
var ctxKeyRequestID int

// WithRequestID returns a new context with the given request ID.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, &ctxKeyRequestID, requestID)
}

// GetRequestID returns a request ID from the given context if one is present.
// Returns the empty string if a request ID cannot be found.
func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if requestID, ok := ctx.Value(&ctxKeyRequestID).(string); ok {
		return requestID
	}
	return ""
}
