package benchmark

import (
	"io"
	"testing"

	"github.com/rs/zerolog"
)

func BenchmarkZeroLogInfoSeq(b *testing.B) {
	zeroLogger := zerolog.New(io.Discard).Level(zerolog.InfoLevel)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zeroLogger.Info().Msg("benchmark message")
	}
}

func BenchmarkZeroLogInfoWithTenAttributesWithCtxSeq(b *testing.B) {
	zeroLogger := zerolog.New(io.Discard).With().Timestamp().Caller().Logger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zeroLogger.Info().
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
}

func BenchmarkZeroLogInfoWithTenAttributesWithoutCtxSeq(b *testing.B) {
	zeroLogger := zerolog.New(io.Discard).Level(zerolog.InfoLevel)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zeroLogger.Info().
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
}
