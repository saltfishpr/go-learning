package internal

import (
	"context"
	"os"
	"os/signal"
	"time"
)

func TimeString2Timestamp(s string, offset int) (int64, error) {
	t, err := time.ParseInLocation(time.DateTime, s, time.UTC)
	if err != nil {
		return 0, err
	}
	return t.Add(time.Duration(-offset) * time.Minute).Unix(), nil
}

func TimestampToTimeString(sec int64, offset int) string {
	return time.Unix(sec, 0).UTC().Add(time.Duration(offset) * time.Minute).Format(time.DateTime)
}

// ContextForSignal returns a context object which is cancelled when a signal
// is received.
// It returns context.Background() if no signal parameter is provided.
func ContextForSignal(signals ...os.Signal) context.Context {
	if len(signals) == 0 {
		return context.Background()
	}

	ch := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())

	// Send message on channel when signal received
	signal.Notify(ch, signals...)

	// When any signal received, call cancel
	go func() {
		<-ch
		cancel()
	}()

	// Return success
	return ctx
}
