package benchmark

import (
	"errors"
	"io"
	"testing"
	"time"

	"github.com/rs/zerolog"
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
func BenchmarkZeroLogInfoPositive(t *testing.B) {
	logger := zerolog.New(io.Discard).Level(zerolog.InfoLevel)
	t.ResetTimer()
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg(TestMessage)
		}
	})
}

// Negative Test -> Level is Error , Log is Info
func BenchmarkZeroLogInfoNegative(t *testing.B) {
	logger := zerolog.New(io.Discard).Level(zerolog.ErrorLevel)
	t.ResetTimer()
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg(TestMessage)
		}
	})
}

func BenchmarkZeroLogTenAttributesWithCtx(t *testing.B) {
	logger := zerolog.New(io.Discard).With().Timestamp().Caller().Logger()
	t.ResetTimer()
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().
				Str("string", TestString).
				Int("status", TestInt).
				Dur("duration", TestDuration).
				Time("time", TestTime).
				Any("error", TestError).
				Str("string", TestString).
				Int("status", TestInt).
				Dur("duration", TestDuration).
				Time("time", TestTime).
				Any("error", TestError)
		}
	})
}

func BenchmarkZeroLogTenAttributesWithoutCtx(t *testing.B) {
	logger := zerolog.New(io.Discard)
	t.ResetTimer()
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().
				Str("string", TestString).
				Int("status", TestInt).
				Dur("duration", TestDuration).
				Time("time", TestTime).
				Any("error", TestError).
				Str("string", TestString).
				Int("status", TestInt).
				Dur("duration", TestDuration).
				Time("time", TestTime).
				Any("error", TestError)
		}
	})
}
