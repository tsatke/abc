package abc

import (
	"testing"
)

func BenchmarkColoredLogger_SimpleLogger_Printf(b *testing.B) {
	logger := &ColoredLogger{
		wrapped: &SimpleLogger{
			clock: &mockClock{},
			lvl:   LevelVerbose,
			out:   &MockWriter{},
		},
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Printf(LevelVerbose, "formatted: %v", "some input")
		}
	})
}
