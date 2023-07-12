package benchmark

import (
	"context"
	"io"
	"testing"

	"golang.org/x/exp/slog"
)

var (
	logger = NewLogger()
)

func NewLogger() *slog.Logger {
	globalHandler := slog.NewTextHandler(io.Discard, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo})
	return slog.New(globalHandler)
}

func BenchmarkGoSLog(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("benchmark message")
	}
}

func BenchmarkGoSLogWithFields(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.LogAttrs(context.Background(), slog.LevelInfo, "benchmark message",
			slog.String("benchmark", "message"),
			slog.Int("benchmark", 1))
	}
}
