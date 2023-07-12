package benchmark

import (
	"context"
	"io"
	"testing"

	"golang.org/x/exp/slog"
)

func BenchmarkGoSLogInfoSeq(b *testing.B) {
	logger := slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info(TestMessage)
	}
}

func BenchmarkGoSLogInfoWithTenAttributesWithCtxSeq(b *testing.B) {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.LogAttrs(ctx, slog.LevelInfo, TestMessage,
			slog.String("string", TestString),
			slog.Int("status", TestInt),
			slog.Duration("duration", TestDuration),
			slog.Time("time", TestTime),
			slog.Any("error", TestError),
			slog.String("string", TestString),
			slog.Int("status", TestInt),
			slog.Duration("duration", TestDuration),
			slog.Time("time", TestTime),
			slog.Any("error", TestError),
		)
	}
}

func BenchmarkGoSLogInfoWithTenAttributesWithoutCtxSeq(b *testing.B) {
	logger := slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.LogAttrs(nil, slog.LevelInfo, TestMessage,
			slog.String("string", TestString),
			slog.Int("status", TestInt),
			slog.Duration("duration", TestDuration),
			slog.Time("time", TestTime),
			slog.Any("error", TestError),
			slog.String("string", TestString),
			slog.Int("status", TestInt),
			slog.Duration("duration", TestDuration),
			slog.Time("time", TestTime),
			slog.Any("error", TestError),
		)
	}
}
