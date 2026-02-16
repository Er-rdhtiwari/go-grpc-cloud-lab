package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// Context is the “request lifetime” concept in Go.
	// Even in main(), we use it to coordinate shutdown.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logger.Info("service starting", "service", "todo", "pid", os.Getpid())

	// Mini “work” loop (just to prove graceful stop wiring).
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logger.Info("shutdown signal received", "err", ctx.Err())
			logger.Info("service stopped", "service", "todo")
			return
		case <-ticker.C:
			logger.Info("heartbeat", "service", "todo", "ts", time.Now().UTC().Format(time.RFC3339))
		}
	}
}
