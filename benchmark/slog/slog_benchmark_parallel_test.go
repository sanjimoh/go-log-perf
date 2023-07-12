package benchmark

import (
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"golang.org/x/exp/slog"
)

const TestMessage = "Test logging, but use a somewhat realistic message length."

var (
	TestTime     = time.Date(2022, time.May, 1, 0, 0, 0, 0, time.UTC)
	TestString   = "7e3b3b2aaeff56a7108fe11e154200dd/7819479873059528190"
	TestInt      = 32768
	TestDuration = 23 * time.Second
	TestError    = errors.New("fail")
)

// Positive Test -> Level is Info , Log is Info
func BenchmarkSlogInfoPositive(t *testing.B) {
	logger := slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}))
	t.ResetTimer()
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info(TestMessage)
		}
	})
}

// Negative Test -> Level is Error , Log is Info
func BenchmarkSlogInfoNegative(t *testing.B) {
	logger := slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{AddSource: true, Level: slog.LevelError}))
	t.ResetTimer()
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info(TestMessage)
		}
	})
}
func BenchmarkSlogTenAttributesWithCtx(t *testing.B) {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}))
	t.ResetTimer()
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
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
	})
}

func BenchmarkSlogTenAttributesWithoutCtx(t *testing.B) {
	logger := slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}))
	t.ResetTimer()
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
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
	})
}
