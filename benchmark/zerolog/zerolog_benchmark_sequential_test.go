package benchmark

import (
	"io"
	"testing"

	"github.com/rs/zerolog"
)

var (
	zeroLogger = NewZeroLogger()
)

func NewZeroLogger() *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: io.Discard}
	logger := zerolog.New(output).With().Caller().Logger()
	return &logger
}

func BenchmarkZeroLog(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zeroLogger.Info().Msg("benchmark message")
	}
}

func BenchmarkZeroLogWithFields(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zeroLogger.Info().
			Str("benchmark", "message").
			Int("benchmark", 1).
			Msg("benchmark message")
	}
}
